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
	}
	npm := Pnpm{
		node{
			commandName: "pnpm",
			config:      config,
		},
	}
	return npm.Run(c.String("npm-profile"))
}
