package ks

import (
	"testing"
)

func TestExecKubectl(t *testing.T) {
	eks := &ExecKubectl{
		Search: &Kubectl{
			Namespace: "scygkj",
			Deployment: &Deployment{
				Imtag: "dslyht",
			},
		},
	}
	err := eks.SearchDeployment()
	if err != nil {
		return
	}

}
