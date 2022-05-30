package ks

import (
	"build-tools/glb"
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
	"strings"
)

type Deployment struct {
	Name   string
	ImName string
	Image  string
}

func findNamespace() string {
	if glb.Con.Docker == nil {
		return ""
	}
	if glb.Con.Docker.DockerRegistry == "" {
		return ""
	}
	spt := strings.Split(glb.Con.Docker.DockerRegistry, "/")
	return spt[len(spt)-1]

}

func finddockimage(json string) *Deployment {
	if glb.Con.Docker == nil {
		return nil
	}
	dep := &Deployment{}
	quer := func(q *gojsonq.JSONQ) *gojsonq.JSONQ {
		return q.WhereStartsWith("image", glb.Con.Docker.DockerRegistry)
	}
	if len(glb.Con.Docker.Tags) > 1 {
		tag := strings.Split(glb.Con.Docker.Tags[0], "-")
		if len(tag) > 1 {
			a := tag[len(tag)-1]
			quer = func(q *gojsonq.JSONQ) *gojsonq.JSONQ {
				return q.WhereEndsWith("image", a)
			}
		}
	}
	t := gojsonq.New().FromString(json).From("items")
	it := t.Count()
	for i := 0; i < it; i++ {
		coun := fmt.Sprintf("[%v]", i)
		a := gojsonq.New().FromString(json).From("items."+coun+".spec.template.spec.containers").Select("name", "image")
		o := quer(a).First()
		if o != nil {
			dep.ImName = o.(map[string]interface{})["name"].(string)
			name := gojsonq.New().FromString(json).Find("items." + coun + ".metadata.name")
			dep.Name = name.(string)
			break
		}
	}
	dep.Image = glb.Con.Docker.DockerRegistry
	return dep
}
