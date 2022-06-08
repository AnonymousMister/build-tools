package ks

import (
	"build-tools/exec"
	"fmt"
)

type ExecKubectl struct {
	Json   string
	tag    string
	Search *Kubectl
	Update *Kubectl
}

func (d *ExecKubectl) SearchDeployment() (e error) {
	d.Json, e = exec.ExecCommand("kubectl", []string{
		"get", "Deployment", "-n" + d.Search.Namespace, "-ojson",
	})
	if e != nil {
		return e
	}
	d.Update, e = d.Search.SearchDeployment(d.Json)
	return e
}

func (d *ExecKubectl) SetImage(tag string) error {
	_, e := exec.ExecCommand("kubectl", []string{
		"set", "image", "Deployment", d.Update.Deployment.Name, d.Update.Deployment.ImName + "=" + d.Update.Deployment.Image + ":" + tag, "-n" + d.Update.Namespace,
	})
	if e == nil {
		fmt.Println(fmt.Sprintf("更新下发成功  镜像 %s", tag))
	}
	return e
}
