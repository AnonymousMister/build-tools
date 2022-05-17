package main

import (
	"build-tools/artifacts"
	"build-tools/docker"
	"build-tools/java"
	"build-tools/step"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	var app = &cli.App{
		Usage:  "编译辅助工具",
		Action: step.Step,
		Flags:  step.InitStepFlag(),
		Commands: []*cli.Command{
			&java.MavenCommand,
			&artifacts.ArtifactsCommand,
			&docker.DockerCommand,
		},
		Version: "0.1",
	}
	if err := app.Run(os.Args); nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
}
