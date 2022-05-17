package artifacts

import (
	"bufio"
	"build-tools/step"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	step.RegisterStepmap(&step.Factory{
		F:    steprAtifacts,
		Name: "artifacts",
	})
	step.RegisterStepFlag(InitArtifactsFlag())
}

var ArtifactsCommand = cli.Command{
	Name:   "artifacts",
	Usage:  "artifacts 文件迁移",
	Action: steprAtifacts,
	Flags:  InitArtifactsFlag(),
}

func InitArtifactsFlag() []cli.Flag {
	flag := []cli.Flag{
		&cli.StringFlag{
			Name:    "artifacts-source",
			Usage:   "artifacts 源文件",
			EnvVars: []string{"ARTIFACTS_SOURCE"},
		},
		&cli.StringFlag{
			Name:    "artifacts-target",
			Usage:   "artifacts 目标文件",
			EnvVars: []string{"ARTIFACTS_TARGET"},
		},
	}
	return flag
}

func steprAtifacts(c *cli.Context) error {
	fmt.Println("********************************************")
	fmt.Println("***********   artifacts 阶段开始  ***********")
	fmt.Println("********************************************")
	str, _ := os.Getwd()
	folderPath := filepath.Join(str, "artifacts")
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.Mkdir(folderPath, 0777)
		if err != nil {
			return err
		}
	}
	artifacts, err := New(c.String("artifacts-source"), c.String("artifacts-target"))
	if err != nil {
		return err
	}
	_, err = artifacts.CopyFile()
	return err
}

type Artifacts struct {
	source  string
	workDir string
	target  string
}

func New(source string, target string) (*Artifacts, error) {
	str, _ := os.Getwd()
	artifacts := &Artifacts{
		workDir: str,
	}
	err := artifacts.SetSource(source)
	if err != nil {
		return nil, err
	}
	err = artifacts.SetTarget(target)
	if err != nil {
		return nil, err
	}
	return artifacts, nil
}

func (a *Artifacts) SetSource(source string) error {
	if source == "" {
		return errors.New("source 不能为空")
	}
	if strings.HasPrefix(source, "/") {
		a.source = source
	} else if strings.HasPrefix(source, "./") {
		a.source = filepath.Join(a.workDir, source[2:])
	} else {
		a.source = filepath.Join(a.workDir, source)
	}
	return nil
}

func (a *Artifacts) SetTarget(target string) error {
	if target == "" {
		return errors.New("source 不能为空")
	}
	if strings.HasPrefix(target, "/") {
		a.target = target
	} else if strings.HasPrefix(target, "./") {
		a.target = filepath.Join(a.workDir, "artifacts", target[2:])
	} else {
		a.target = filepath.Join(a.workDir, "artifacts", target)
	}
	return nil
}

func (a *Artifacts) CopyFile() (written int64, err error) {
	srcFile, err := os.Open(a.source)
	if err != nil {
		fmt.Println("open file err:", err)
	}
	//关闭流
	defer srcFile.Close()
	//获取到reader
	reader := bufio.NewReader(srcFile)
	//打开dstFileName
	dstFile, err := os.OpenFile(a.target, os.O_WRONLY|os.O_CREATE, 0666) //0666 在windos下无效
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	writer := bufio.NewWriter(dstFile)
	//关闭流
	defer dstFile.Close()
	//调用copy函数
	return io.Copy(writer, reader)
}
