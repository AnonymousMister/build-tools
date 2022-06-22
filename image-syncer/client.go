package image_syncer

import (
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/AliyunContainerService/image-syncer/pkg/sync"
	"github.com/AliyunContainerService/image-syncer/pkg/tools"
	"github.com/sirupsen/logrus"
	"sort"
	"strings"
	sync2 "sync"
	"time"
)

// Client describes a synchronization client
type Client struct {
	// a sync.Task list
	taskList *list.List

	// a URLPair list
	urlPairList *list.List

	// failed list
	failedTaskList         *list.List
	failedTaskGenerateList *list.List

	config *Config

	routineNum int
	retries    int
	increment  bool
	logger     *logrus.Logger

	// mutex
	taskListChan               chan int
	urlPairListChan            chan int
	failedTaskListChan         chan int
	failedTaskGenerateListChan chan int
	tasksChan                  chan *sync.Task
}

// URLPair is a pair of source and destination url
type URLPair struct {
	source      string
	destination string
}

// NewSyncClient creates a synchronization client
func NewSyncClient(authFile, imageFile, logFile string,
	routineNum, retries int, increment bool, defaultDestRegistry, defaultDestNamespace string,
	osFilterList, archFilterList []string) (*Client, error) {

	logger := NewFileLogger(logFile)

	config, err := NewSyncConfig(authFile, imageFile,
		defaultDestRegistry, defaultDestNamespace, osFilterList, archFilterList)
	if err != nil {
		return nil, fmt.Errorf("generate config error: %v", err)
	}

	return &Client{
		taskList:                   list.New(),
		urlPairList:                list.New(),
		failedTaskList:             list.New(),
		failedTaskGenerateList:     list.New(),
		config:                     config,
		routineNum:                 routineNum,
		retries:                    retries,
		increment:                  increment,
		logger:                     logger,
		taskListChan:               make(chan int, 1),
		urlPairListChan:            make(chan int, 1),
		failedTaskListChan:         make(chan int, 1),
		failedTaskGenerateListChan: make(chan int, 1),
		tasksChan:                  make(chan *sync.Task, 500),
	}, nil
}

// Run is main function of a synchronization client
func (c *Client) Run() {
	fmt.Println("开始执行同步..............")

	//var finishChan = make(chan struct{}, c.routineNum)

	// open num of goroutines and wait c for close
	openRoutinesGenTaskAndWaitForFinish := func() {
		wg := sync2.WaitGroup{}
		for i := 0; i < c.routineNum; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					urlPair, empty := c.GetAURLPair()
					// no more task to generate
					if empty {
						break
					}
					moreURLPairs, err := c.GenerateSyncTask(urlPair.source, urlPair.destination)
					if err != nil {
						c.logger.Errorf("Generate sync task %s to %s error: %v", urlPair.source, urlPair.destination, err)
						// put to failedTaskGenerateList
						c.PutAFailedURLPair(urlPair)
					}
					if moreURLPairs != nil {
						c.PutURLPairs(moreURLPairs)
					}
				}
			}()
		}
		wg.Wait()
	}

	//开始初始化线程 并且监听 队列数据
	openRoutinesHandleTaskAndWaitForFinish := func() {
		for i := 0; i < c.routineNum; i++ {
			go func(i int) {
				for {
					task, ok := <-c.tasksChan
					// no more tasks need to handle
					if !ok || task == nil {
						break
					}

					v, _ := json.Marshal(task)
					c.logger.Infof("开始执行 线程【%d】 推送任务 %#v", i, string(v))
					if err := task.Run(); err != nil {
						c.logger.Infof("执行出现错误 线程【%d】", i)
						// put to failedTaskList
						c.PutAFailedTask(task)
					}
					c.logger.Infof("执行结束 线程【%d】", i)
				}
			}(i)
		}
	}

	for source, dest := range c.config.GetImageList() {
		c.urlPairList.PushBack(&URLPair{
			source:      source,
			destination: dest,
		})
	}

	// generate goroutines to handle sync tasks
	openRoutinesHandleTaskAndWaitForFinish()
	// 监听状态
	go func() {
		time.Sleep(time.Second * 3)
		for {
			time.Sleep(time.Second * 3)
			cd := len(c.tasksChan)
			if cd == 0 {
				break
			}
			fmt.Println("监听同步任务处理完成...还剩下【%v】", cd)
		}
	}()
	// generate sync tasks
	openRoutinesGenTaskAndWaitForFinish()
	fmt.Println("Start to handle sync tasks, please wait ...")
	for {
		cd := len(c.tasksChan)
		if cd == 0 {
			time.Sleep(time.Second * 3)
			close(c.tasksChan)
			break
		} else {
			fmt.Println("等待同步任务处理完成...还剩下【%v】", cd)
		}
	}
	fmt.Printf("Finished, %v sync tasks failed, %v tasks generate failed\n", c.failedTaskList.Len(), c.failedTaskGenerateList.Len())
	c.logger.Infof("Finished, %v sync tasks failed, %v tasks generate failed", c.failedTaskList.Len(), c.failedTaskGenerateList.Len())
}

func (c *Client) ImageSource(sourceURL *tools.RepoURL) (imageSource *sync.ImageSource, err error) {
	if auth, exist := c.config.GetAuth(sourceURL.GetRegistry(), sourceURL.GetNamespace()); exist {
		c.logger.Infof("Find auth information for %v, username: %v", sourceURL.GetURL(), auth.Username)
		imageSource, err = sync.NewImageSource(sourceURL.GetRegistry(), sourceURL.GetRepoWithNamespace(), sourceURL.GetTag(),
			auth.Username, auth.Password, auth.Insecure)
		if err != nil {
			return nil, fmt.Errorf("generate %s image source error: %v", sourceURL.GetURL(), err)
		}
	} else {
		c.logger.Infof("Cannot find auth information for %v, pull actions will be anonymous", sourceURL.GetURL())
		imageSource, err = sync.NewImageSource(sourceURL.GetRegistry(), sourceURL.GetRepoWithNamespace(), sourceURL.GetTag(),
			"", "", false)
		if err != nil {
			return nil, fmt.Errorf("generate %s image source error: %v", sourceURL.GetURL(), err)
		}
	}
	return imageSource, nil
}

func (c *Client) screen(sourceTags []string, destinationTags []string) []string {
	if destinationTags == nil || len(destinationTags) == 0 {
		return sourceTags
	}
	sourceTagsLen := len(sourceTags)
	destinationTagsLen := len(destinationTags)
	if sourceTagsLen == destinationTagsLen {
		return nil
	}
	var tags = make([]string, 0)
	sort.Strings(destinationTags)
	for _, v := range sourceTags {
		if strings.HasSuffix(v, "latest") ||
			strings.HasPrefix(v, "rc-") {
			tags = append(tags, v)
			continue
		}
		index := sort.SearchStrings(destinationTags, v)
		if index == destinationTagsLen || destinationTags[index] != v {
			tags = append(tags, v)
		}
	}
	return tags
}

// GenerateSyncTask creates synchronization tasks from source and destination url, return URLPair array if there are more than one tags
func (c *Client) GenerateSyncTask(source string, destination string) ([]*URLPair, error) {
	if source == "" {
		return nil, fmt.Errorf("source url should not be empty")
	}

	sourceURL, err := tools.NewRepoURL(source)
	if err != nil {
		return nil, fmt.Errorf("url %s format error: %v", source, err)
	}

	// if dest is not specific, use default registry and namespace
	if destination == "" {
		return nil, fmt.Errorf("the default registry and namespace should not be nil if you want to use them")
	}

	destURL, err := tools.NewRepoURL(destination)
	if err != nil {
		return nil, fmt.Errorf("url %s format error: %v", destination, err)
	}

	tags := sourceURL.GetTag()

	// multi-tags config
	if moreTag := strings.Split(tags, ","); len(moreTag) > 1 {
		if destURL.GetTag() != "" && destURL.GetTag() != sourceURL.GetTag() {
			return nil, fmt.Errorf("multi-tags source should not correspond to a destination with tag: %s:%s",
				sourceURL.GetURL(), destURL.GetURL())
		}

		// contains more than one tag
		var urlPairs []*URLPair
		for _, t := range moreTag {
			urlPairs = append(urlPairs, &URLPair{
				source:      sourceURL.GetURLWithoutTag() + ":" + t,
				destination: destURL.GetURLWithoutTag() + ":" + t,
			})
		}

		return urlPairs, nil
	}

	var imageSource *sync.ImageSource
	var imageDestination *sync.ImageDestination

	if imageSource, err = c.ImageSource(sourceURL); err != nil {
		return nil, err
	}
	// if tag is not specific, return tags
	if sourceURL.GetTag() == "" {

		if destURL.GetTag() != "" {
			return nil, fmt.Errorf("tag should be included both side of the config: %s:%s", sourceURL.GetURL(), destURL.GetURL())
		}

		// get all tags of this source repo
		tags, err := imageSource.GetSourceRepoTags()
		if err != nil {
			return nil, fmt.Errorf("get tags failed from %s error: %v", sourceURL.GetURL(), err)
		}
		if !c.increment {
			var imageDest *sync.ImageSource
			if imageDest, err = c.ImageSource(destURL); err != nil {
				return nil, err
			}
			dtags, err := imageDest.GetSourceRepoTags()
			if err == nil {
				c.logger.Infof("************** Get tags of %s successfully ", destURL.GetURL())
				tags = c.screen(tags, dtags)
			}
		}

		if tags == nil || len(tags) == 0 {
			return nil, fmt.Errorf("不需要同步的tag")
		}
		c.logger.Infof("Get tags of %s successfully: %v", sourceURL.GetURL(), tags)
		// generate url pairs for tags
		var urlPairs = []*URLPair{}
		for _, tag := range tags {
			urlPairs = append(urlPairs, &URLPair{
				source:      sourceURL.GetURL() + ":" + tag,
				destination: destURL.GetURL() + ":" + tag,
			})
		}
		return urlPairs, nil
	}

	// if source tag is set but without destination tag, use the same tag as source
	destTag := destURL.GetTag()
	if destTag == "" {
		destTag = sourceURL.GetTag()
	}

	if auth, exist := c.config.GetAuth(destURL.GetRegistry(), destURL.GetNamespace()); exist {
		c.logger.Infof("Find auth information for %v, username: %v", destURL.GetURL(), auth.Username)
		imageDestination, err = sync.NewImageDestination(destURL.GetRegistry(), destURL.GetRepoWithNamespace(),
			destTag, auth.Username, auth.Password, auth.Insecure)
		if err != nil {
			return nil, fmt.Errorf("generate %s image destination error: %v", sourceURL.GetURL(), err)
		}
	} else {
		c.logger.Infof("Cannot find auth information for %v, push actions will be anonymous", destURL.GetURL())
		imageDestination, err = sync.NewImageDestination(destURL.GetRegistry(), destURL.GetRepoWithNamespace(),
			destTag, "", "", false)
		if err != nil {
			return nil, fmt.Errorf("generate %s image destination error: %v", destURL.GetURL(), err)
		}
	}

	c.PutATask(sync.NewTask(imageSource, imageDestination, c.config.osFilterList, c.config.archFilterList, c.logger))
	c.logger.Infof("Generate a task for %s to %s", sourceURL.GetURL(), destURL.GetURL())
	return nil, nil
}

// PutATask puts a sync.Task struct to task list
func (c *Client) PutATask(task *sync.Task) {
	c.tasksChan <- task
	if c.taskList != nil {
		c.taskList.PushBack(task)
	}
}

// GetAURLPair gets a URLPair from urlPairList
func (c *Client) GetAURLPair() (*URLPair, bool) {
	c.urlPairListChan <- 1
	defer func() {
		<-c.urlPairListChan
	}()

	urlPair := c.urlPairList.Front()
	if urlPair == nil {
		return nil, true
	}
	c.urlPairList.Remove(urlPair)

	return urlPair.Value.(*URLPair), false
}

// PutURLPairs puts a URLPair array to urlPairList
func (c *Client) PutURLPairs(urlPairs []*URLPair) {
	c.urlPairListChan <- 1
	defer func() {
		<-c.urlPairListChan
	}()

	if c.urlPairList != nil {
		for _, urlPair := range urlPairs {
			c.urlPairList.PushBack(urlPair)
		}
	}
}

// PutAFailedTask puts a failed task to failedTaskList
func (c *Client) PutAFailedTask(failedTask *sync.Task) {
	c.failedTaskListChan <- 1
	defer func() {
		<-c.failedTaskListChan
	}()

	if c.failedTaskList != nil {
		c.failedTaskList.PushBack(failedTask)
	}
}

// GetAFailedURLPair get a URLPair from failedTaskGenerateList
func (c *Client) GetAFailedURLPair() (*URLPair, bool) {
	c.failedTaskGenerateListChan <- 1
	defer func() {
		<-c.failedTaskGenerateListChan
	}()

	failedURLPair := c.failedTaskGenerateList.Front()
	if failedURLPair == nil {
		return nil, true
	}
	c.failedTaskGenerateList.Remove(failedURLPair)

	return failedURLPair.Value.(*URLPair), false
}

// PutAFailedURLPair puts a URLPair to failedTaskGenerateList
func (c *Client) PutAFailedURLPair(failedURLPair *URLPair) {
	c.failedTaskGenerateListChan <- 1
	defer func() {
		<-c.failedTaskGenerateListChan
	}()

	if c.failedTaskGenerateList != nil {
		c.failedTaskGenerateList.PushBack(failedURLPair)
	}
}
