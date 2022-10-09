package ks

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestExecKubectl(t *testing.T) {
	eks := &ExecKubectl{
		Search: &Kubectl{
			Namespace: "ynrhjc",
			Deployment: &Deployment{
				Imtag: "rhjcht-new",
			},
		},
	}
	b, _ := json.Marshal(eks)
	fmt.Println("ks json marshal :", string(b))
	err := eks.SearchDeployment()
	if err != nil {
		return
	}

}
