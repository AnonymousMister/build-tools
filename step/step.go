package step

import (
	"github.com/AnonymousMister/build-tools/java"
	"github.com/urfave/cli"
)

var (
	stepmap   = make(map[string]cli.ActionFunc)
)

type Factory struct {
	F    cli.ActionFunc
	Name string
}

func RegisterStepmap(factory *Factory) {
	stepmap[factory.Name] = factory.F
}

func InitStepFlag() []cli.Flag {
	flag := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "step",
			Usage:   "step 步骤 可选值：java,docker,artifacts,node",
		},
	}
	flag = append(flag, java.InitMavenFlag()...)
	return flag
}
func Step(c *cli.Context) error {
	steps:=c.StringSlice("step")
	for _, step := range steps {
		if f, ok := stepmap[step]; ok {
			err := f(c)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
