package ks

import (
	"build-tools/exec"
	"build-tools/glb"
	"build-tools/step"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
)

func init() {
	step.RegisterStepmap(&step.Factory{
		F:    steprks,
		Name: "ks",
	})
	step.RegisterStepFlag(InitKSFlag())
}

var KS = cli.Command{
	Name:   "ks",
	Usage:  "k8s 推送",
	Action: steprks,
	Flags:  InitKSFlag(),
}

func InitKSFlag() []cli.Flag {
	flag := []cli.Flag{
		&cli.StringFlag{
			Name:    "ks-namespace",
			Usage:   "namespace 名称",
			EnvVars: []string{"KS_NAMESPACE"},
		},
	}
	return flag
}

func steprks(c *cli.Context) error {
	fmt.Println("********************************************")
	fmt.Println("*************    ks 阶段开始    *************")
	fmt.Println("********************************************")
	namespace := c.String("ks-namespace")
	if namespace == "" {
		namespace = findNamespace()
	}
	if namespace == "" {
		return errors.New("没有找到 空间")
	}
	ks := &kubectl{
		Namespace: namespace,
	}
	val, e := exec.ExecCommand("kubectl", []string{
		"get", "Deployment", "-n" + ks.Namespace, "-ojson",
	})
	if val == "" {
		return errors.New("没有找到 下发对象")
	}
	if e != nil {
		return e
	}
	dep := finddockimage(val)
	if dep == nil {
		return errors.New("没有找到 可以跟新对象")
	}
	ks.Deployment = dep
	return ks.SetImage(glb.Con.Docker.Tags[0])
}

type kubectl struct {
	Namespace  string
	Deployment *Deployment
}

func (d *kubectl) SetImage(tag string) error {
	_, e := exec.ExecCommand("kubectl", []string{
		"set", "image", "Deployment", d.Deployment.Name, d.Deployment.ImName + "=" + d.Deployment.Image + ":" + tag, "-n" + d.Namespace,
	})
	if e == nil {
		fmt.Println(fmt.Sprintf("更新下发成功  镜像 %s", tag))
	}
	return e
}
