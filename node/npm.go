package node

import (
	"build-tools/exec"
	"build-tools/glb"
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
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
	options []string
	profile string
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
	return exec.ExecCommand("npm", option)
}

func npm(c *cli.Context) error {
	fmt.Println("********************************************")
	fmt.Println("***********     npm 阶段开始    ***********")
	fmt.Println("********************************************")
	var options []string
	option := c.String("npm-option")
	taobao := c.Bool("npm-taobao")
	if taobao {
		index := strings.Index(option, "--registry")
		if index == -1 {
			options = append(options, "--registry=https://registry.npmmirror.com")
		}
		index = strings.Index(option, "--disturl")
		if index == -1 {
			options = append(options, "--disturl=https://npmmirror.com/mirrors/node")
		}
	}
	if option != "" {
		options = append(options, option)
	}
	npm := Npm{
		options: options,
		profile: c.String("npm-profile"),
	}
	return npm.RunProfile()
}
