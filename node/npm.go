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
		&cli.BoolFlag{
			Name:    "pnpm-shamefully-hoist",
			Usage:   `是否 扁平化安装 （兼容老框架）`,
			EnvVars: []string{"PNPM_SHAMEFULLY_HOIST"},
			Value:   false,
		},
	}
	return flag
}

type Npm struct {
	node
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
		config["cache"] = "/home/cache/.npm/.cache/cnpm"
	}
	npm := node{
		commandName: "npm",
		config:      config,
	}
	if glb.IsDebug {
		exec.ExecCommand("nvm", []string{
			"use",
			"v14.19.1",
		})
	}
	e := npm.SetConfigs()
	if e != nil {
		return e
	}
	e = npm.Install()
	if e != nil {
		return e
	}
	return npm.Run(c.String("npm-profile"))
}
