package ks

import (
	"build-tools/glb"
	"strings"
)

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

func findSearchDeployment() *Deployment {
	if glb.Con.Docker == nil {
		return nil
	}
	dep := &Deployment{}
	dep.Image = glb.Con.Docker.DockerRegistry
	if len(glb.Con.Docker.Tags) > 1 {
		tag := strings.Split(glb.Con.Docker.Tags[0], "-")
		if len(tag) > 1 {
			a := tag[len(tag)-1]
			dep.Image = ""
			dep.Imtag = a
		}
	}
	return dep
}
