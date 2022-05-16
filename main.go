package main

import (
	"fmt"
	"github.com/AnonymousMister/build-tools/java"
	"github.com/AnonymousMister/build-tools/step"
	"github.com/urfave/cli/v2"
	"os"
)


func main() {
	var app = &cli.App{
		Usage: "编译辅助工具",
		Action: step.Step,
		Flags: step.InitStepFlag(),
		Commands: []*cli.Command{
			&java.MavenCommand,
		},
		Version: "0.1",
	}
	if err := app.Run(os.Args); nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
}

