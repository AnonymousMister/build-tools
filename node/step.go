package node

import (
	"build-tools/exec"
	"build-tools/step"
	"errors"
	"github.com/urfave/cli/v2"
)

func init() {
	step.RegisterStepmap(&step.Factory{
		F:    stepNode,
		Name: "node",
	})
	step.RegisterStepFlag(InitNpmFlag())
}
func stepNode(c *cli.Context) error {
	if exec.CheckFileIsExist("package-lock.json") {
		return npm(c)
	} else if exec.CheckFileIsExist("package.json") {
		return npm(c)
	}
	return errors.New("没有找到 package.json 文件")
}
