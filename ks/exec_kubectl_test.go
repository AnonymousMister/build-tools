package ks

import (
	"testing"
)

func TestExecKubectl(t *testing.T) {

	eks := &ExecKubectl{
		Search: &Kubectl{
			Namespace: "wjrb",
			Deployment: &Deployment{
				Name: "wjrb-test",
			},
		},
	}
	err := eks.SearchDeployment()
	if err != nil {
		return
	}

}
