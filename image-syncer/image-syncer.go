package image_syncer

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var ImageSyncerCommand = cli.Command{
	Name:   "image-syncer",
	Usage:  "docker image 同步 工具",
	Action: steprImageSyncer,
	Flags:  ImageSyncerFlag(),
}

func ImageSyncerFlag() []cli.Flag {
	flag := []cli.Flag{
		&cli.StringFlag{
			Name:  "auth",
			Usage: "auth 配置文件 路径",
		},
		&cli.StringFlag{
			Name:  "images",
			Usage: "images 配置文件 路径",
		},
		&cli.StringFlag{
			Name:  "log",
			Usage: "log  文件路径",
		},
		&cli.BoolFlag{
			Name:  "increment",
			Usage: "是否覆盖镜像",
			Value: false,
		},
		&cli.IntFlag{
			Name:  "proc",
			Usage: "并发数",
			Value: 5,
		},
	}
	return flag
}
func steprImageSyncer(c *cli.Context) error {
	increment := c.Bool("increment")
	auth := c.String("auth")
	images := c.String("images")
	log := c.String("log")
	proc := c.Int("proc")

	client, err := NewSyncClient(auth, images, log, proc, 2, increment, "", "",
		[]string{},
		[]string{},
	)
	if err != nil {
		return fmt.Errorf("init sync client error: %v", err)
	}
	client.Run()
	return nil
}
