package node

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

type Pnpm struct {
	node
}

func pnpm(c *cli.Context) error {
	fmt.Println("********************************************")
	fmt.Println("***********     pnpm 阶段开始    ***********")
	fmt.Println("********************************************")
	config := make(map[string]string)
	taobao := c.Bool("npm-taobao")
	config["registry"] = "/home/cache/.pnpm-store"
	if taobao {
		config["registry"] = "https://registry.npmmirror.com"
		config["disturl"] = "https://npmmirror.com/mirrors/node"
		config["cache"] = "/home/cache/.npm/.cache/cnpm"

	}
	npm := Pnpm{
		node{
			commandName: "pnpm",
			config:      config,
		},
	}
	e := npm.SetConfigs()
	if e != nil {
		return e
	}
	shamefully := c.Bool("pnpm-shamefully-hoist")
	if shamefully {
		e = npm.Install("--shamefully-hoist")
	} else {
		e = npm.Install()
	}
	if e != nil {
		return e
	}
	return npm.Run(c.String("npm-profile"))
}
