package ks

import (
	"testing"
)

func TestExecKubectl(t *testing.T) {
	eks := &ExecKubectl{
		Search: &Kubectl{
			Namespace: "ynrhjc",
			Deployment: &Deployment{
				Imtag: "rhjcjava",
			},
		},
	}
	err := eks.SearchDeployment()
	if err != nil {
		return
	}

}
