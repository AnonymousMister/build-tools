package ks

import (
	"errors"
	"github.com/thedevsaddam/gojsonq/v2"
	"strings"
)

type Kubectl struct {
	Namespace  string
	Deployment *Deployment
}

type Deployment struct {
	Name   string
	ImName string
	Image  string
	Imtag  string
}

type serch = func(q *gojsonq.JSONQ) *gojsonq.JSONQ

func (d *Kubectl) SearchDeployment(Json string) ([]*Kubectl, error) {
	kubectls := make([]*Kubectl, 0)
	// 获取 metadata 过滤选项
	metadatas := d.searchDeploymentMetadata()
	deploymentJson := gojsonq.New().FromString(Json).From("items")
	// 执行 过滤 和 过滤 replicas 有效果的 数据
	aa := metadatas(deploymentJson).WhereNotEqual("spec.replicas", 0).Get()
	if aa == nil {
		return nil, errors.New("没有找到 下发对象")
	}
	deploymentObjects := aa.([]interface{})
	if it := len(deploymentObjects); it > 0 {
		containers := d.searchDeploymentContainers()
		for i := 0; i < it; i++ {
			deploymentObject := deploymentObjects[i]
			containersJson := gojsonq.New().FromInterface(deploymentObject).
				From("spec.template.spec.containers").Select("name", "image")
			o := containers(containersJson).First()
			if o != nil {
				mapo := o.(map[string]interface{})
				deployment := &Deployment{}
				deployment.ImName = mapo["name"].(string)
				images := strings.Split(mapo["image"].(string), ":")
				deployment.Image = images[0]
				if len(images) >= 2 {
					deployment.Imtag = images[1]
				}
				name := containersJson.Reset().Find("metadata.name")
				deployment.Name = name.(string)
				kubectls = append(kubectls, &Kubectl{
					Namespace:  d.Namespace,
					Deployment: deployment,
				})
			}
		}
	}
	if len(kubectls) == 0 {
		return nil, errors.New("没有找到 下发对象")
	}
	return kubectls, nil
}

func (d *Kubectl) searchDeploymentMetadata() serch {
	return func(q *gojsonq.JSONQ) *gojsonq.JSONQ {
		if d.Deployment.Name != "" {
			q.WhereStartsWith("metadata.name", d.Deployment.Name)
		}
		return q
	}
}

func (d *Kubectl) searchDeploymentContainers() serch {
	var quer serch
	if d.Deployment.Imtag != "" {
		quer = func(q *gojsonq.JSONQ) *gojsonq.JSONQ {
			return q.WhereEndsWith("image", d.Deployment.Imtag)
		}
	} else if d.Deployment.Image != "" {
		quer = func(q *gojsonq.JSONQ) *gojsonq.JSONQ {
			return q.WhereStartsWith("image", d.Deployment.Image)
		}
	} else {
		quer = func(q *gojsonq.JSONQ) *gojsonq.JSONQ {
			return q
		}
	}
	return quer
}
