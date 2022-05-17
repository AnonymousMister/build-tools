package java

import (
	"build-tools/exec"
	"build-tools/step"
	"errors"
	"github.com/urfave/cli/v2"
)

func init() {
	step.RegisterStepmap(&step.Factory{
		F:    stepJava,
		Name: "java",
	})
	step.RegisterStepFlag(InitMavenFlag())
}
func stepJava(c *cli.Context) error {
	if exec.CheckFileIsExist("pom.xml") {
		return maven(c)
	}
	return errors.New("没有找到 pom.xml 文件")
}
