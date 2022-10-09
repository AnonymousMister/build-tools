package ks

import (
	"build-tools/exec"
	"fmt"
)

type ExecKubectl struct {
	Json   string
	tag    string
	Search *Kubectl
	Update []*Kubectl
}

func (d *ExecKubectl) SearchDeployment() (e error) {
	d.Json, e = exec.ExecNoCommand("kubectl", []string{
		"get", "Deployment", "-n" + d.Search.Namespace, "-ojson",
	})
	if e != nil {
		return e
	}
	d.Update, e = d.Search.SearchDeployment(d.Json)
	return e
}

func (d *ExecKubectl) SetImage(tag string) error {
	isErroe := false
	if len(d.Update) > 0 {
		for _, v := range d.Update {
			image := v.Deployment.Image + ":" + tag
			_, e := exec.ExecCommand("kubectl", []string{
				"set", "image", "Deployment", v.Deployment.Name, v.Deployment.ImName + "=" + image, "-n" + v.Namespace,
			})
			if e == nil {
				fmt.Println(fmt.Sprintf("更新下发成功  镜像 %s  空间 %s", image, v.Namespace))
			} else {
				isErroe = true
				fmt.Println(fmt.Sprintf("更新下发失败  镜像 %s 空间 %s", image, v.Namespace))
			}
		}

	}
	if isErroe {
		return fmt.Errorf("有下发失败的 镜像 请查看 日志")
	}
	return nil
}
