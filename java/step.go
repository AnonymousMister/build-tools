package java

import (
	"errors"
	"github.com/AnonymousMister/build-tools/exec"
	"github.com/AnonymousMister/build-tools/step"
	"github.com/urfave/cli"
)
func init() {
	step.RegisterStepmap(&step.Factory{
		F:    stepJava,
		Name: "java",
	})
}
func stepJava(c *cli.Context) error{
	if exec.CheckFileIsExist("pom.xml") {
		return maven(c)
	}
	return errors.New("没有找到 pom.xml 文件")
}
