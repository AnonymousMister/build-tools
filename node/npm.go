package node

import (
	"build-tools/exec"
	"build-tools/glb"
	"fmt"
	"github.com/urfave/cli/v2"
)

var NpmCommand = cli.Command{
	Name:   "npm",
	Usage:  "npm 编译",
	Action: npm,
	Flags:  InitNpmFlag(),
}

func InitNpmFlag() []cli.Flag {
	flag := []cli.Flag{
		&cli.StringFlag{
			Name:    "npm-option",
			Usage:   "npm-option 选项",
			EnvVars: []string{"NPM_OPTION"},
		},
		&cli.StringFlag{
			Name:    "npm-profile",
			Usage:   "npm-profile profiles",
			EnvVars: []string{"NPM_PROFILE"},
		},
		&cli.BoolFlag{
			Name: "npm-taobao",
			Usage: `是否起用 淘宝 加速 
		启用添加以下参数 
		--registry=https://registry.npmmirror.com 
		--disturl=https://npmmirror.com/mirrors/node
		此项 会被 npm-option 给覆盖`,
			EnvVars: []string{"TAOBAO_AGENT"},
			Value:   true,
		},
	}
	return flag
}

type Npm struct {
	node
}

func (n *Npm) RunProfile() error {
	option := append(n.options, "run")
	if n.profile != "" {
		option = append(option, n.profile)
	}
	if glb.IsDebug {
		exec.ExecCommand("nvm", []string{
			"use",
			"v14.19.1",
		})
	}

	_, e := exec.ExecCommand("npm", option)

	return e
}

func npm(c *cli.Context) error {
	fmt.Println("********************************************")
	fmt.Println("***********     npm 阶段开始    ***********")
	fmt.Println("********************************************")
	config := make(map[string]string)

	taobao := c.Bool("npm-taobao")
	if taobao {
		config["registry"] = "https://registry.npmmirror.com"
		config["disturl"] = "https://npmmirror.com/mirrors/node"
	}

	npm := node{
		commandName: "npm",
		config:      config,
	}
	return npm.Run(c.String("npm-profile"))
}
