package dingding

import (
	"build-tools/step"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
)

func init() {
	step.RegisterStepmap(&step.Factory{
		F:    steprDingding,
		Name: "dingding",
	})
	step.RegisterStepFlag(InitDingdingFlag())
}

var DingdingCommand = cli.Command{
	Name:   "dingding",
	Usage:  "dingding 通知",
	Action: steprDingding,
	Flags:  InitDingdingFlag(),
}

func InitDingdingFlag() []cli.Flag {
	flag := []cli.Flag{
		&cli.StringFlag{
			Name:    "dingding-config-file",
			Usage:   "dingding 配置文件 路径",
			EnvVars: []string{"DINGDING_CONFIG_FILE"},
		},
		&cli.StringFlag{
			Name:    "docker-url",
			Usage:   "docker 仓库域名",
			EnvVars: []string{"DOCKER_REGISTRY_URL"},
		},
		&cli.StringFlag{
			Name:    "docker-namespace",
			Usage:   "docker 命名空间",
			EnvVars: []string{"DOCKER_REGISTRY_NAMESPACE"},
		},
		&cli.StringFlag{
			Name:    "docker-tag",
			Usage:   "docker 打包标签",
			EnvVars: []string{"DOCKER_REGISTRY_TAG"},
		},
	}
	return flag
}

func steprDingding(c *cli.Context) error {
	fmt.Println("********************************************")
	fmt.Println("**********    dingding 阶段开始    ***********")
	fmt.Println("********************************************")
	if c.String("dingding-config-file") == "" {
		return errors.New("dingding 配置文件不能为空")
	}

	return nil
}
