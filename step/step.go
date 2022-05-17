package step

import (
	"errors"
	"github.com/urfave/cli/v2"
	"strings"
)

var (
	stepmap = make(map[string]cli.ActionFunc)
	flag    = []cli.Flag{
		&cli.StringFlag{
			Name:    "step",
			Usage:   "step 步骤 可选值：java,docker,artifacts,node",
			EnvVars: []string{"BUILD_TOOLS_STEP"},
		},
	}
)

type Factory struct {
	F    cli.ActionFunc
	Name string
}

func RegisterStepmap(factory *Factory) {
	stepmap[factory.Name] = factory.F
}

func RegisterStepFlag(flags []cli.Flag) {
	flag = append(flag, flags...)
}

func InitStepFlag() []cli.Flag {
	return flag
}
func Step(c *cli.Context) error {
	step := c.String("step")
	if step == "" {
		return errors.New("step 不能是空的")
	}
	steps := strings.Split(step, ",")
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
