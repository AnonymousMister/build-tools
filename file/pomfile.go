package file

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	ostype = os.Getenv("GOOS") // 获取系统类型
)

var listfile []string //获取文件列表

func Listfunc(path string, f os.FileInfo, err error) error {
	var strRet string
	strRet, _ = os.Getwd()
	//ostype := os.Getenv("GOOS") // windows, linux

	if ostype == "windows" {
		strRet += "\\"
	} else if ostype == "linux" {
		strRet += "/"
	}

	if f == nil {
		return err
	}
	if f.IsDir() {
		return nil
	}

	strRet += path //+ " "

	//用strings.HasSuffix(src, suffix)//判断src中是否包含 suffix结尾
	ok := strings.HasSuffix(strRet, ".go")
	if ok {

		listfile = append(listfile, strRet) //将目录push到listfile []string中
	}
	//fmt.Println(ostype) // print ostype
	fmt.Println(strRet) //list the file

	return nil
}

func getFileList(path string) ([]string, error) {
	var listfile []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		//var strRet string
		//if ostype == "windows" {
		//	strRet += "\\"
		//} else if ostype == "linux" {
		//	strRet += "/"
		//}
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		ok := strings.HasSuffix(path, "pom.xml")
		if ok {
			listfile = append(listfile, path) //将目录push到listfile []string中
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return listfile, err
}

func ListFileFunc(p []string) {
	for index, value := range p {
		fmt.Println("Index = ", index, "Value = ", value)
	}
}

var PomfileCommand = cli.Command{
	Name:   "tihuan",
	Usage:  "tihuan命令",
	Action: steprPomfile,
	Flags:  PomfileFlag(),
}

func PomfileFlag() []cli.Flag {
	flag := []cli.Flag{
		&cli.StringFlag{
			Name:    "path",
			Usage:   "路径",
			Aliases: []string{"p"},
		},
	}
	return flag
}

func steprPomfile(c *cli.Context) error {
	fmt.Println("********************************************")
	path := c.String("path")
	list, err := getFileList(path)
	if err != nil {
		return err
	}
	var data string
	for _, s := range list {
		err, data = readPomfile(s)
		if err != nil {
			continue
		}
		updatePomfile(s, data)

	}
	return nil
}
func updatePomfile(file string, data string) error {
	fw, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	if err != nil {
		return err
	}
	defer fw.Close()
	w := bufio.NewWriter(fw)
	_, err = w.WriteString(data)
	if err != nil {
		return err
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}

func readPomfile(file string) (error, string) {
	f, err := os.Open(file)
	if err != nil {
		return err, ""
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, c := buf.ReadLine()

		if c == io.EOF {
			break
		}
		line := string(a)
		if strings.Contains(line, "<groupId>com.xuxueli</groupId>") {
			result += strings.Replace(line, "<groupId>com.xuxueli</groupId>", "<groupId>com.github.AnonymousMister</groupId>", -1) + "\n"
		} else {
			result += line + "\n"
		}
	}
	return nil, result
}
