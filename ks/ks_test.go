package ks

import (
	"build-tools/glb"
	"encoding/json"
	"fmt"
	"testing"
)

func TestKs(t *testing.T) {
	//初始化
	glb.Con.Docker = &glb.DockerContext{
		Tags:           []string{"e1b7fdb9-rhjcht-new", "latest-rhjcht-new"},
		DockerRegistry: "registry.cn-hangzhou.aliyuncs.com/winjoin/scygkj",
	}
	dep := findSearchDeployment()
	b, _ := json.Marshal(dep)
	fmt.Println("ks json marshal :", string(b))
}
