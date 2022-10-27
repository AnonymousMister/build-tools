package ks

import (
	"build-tools/glb"
	"build-tools/step"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

func init() {
	step.RegisterStepmap(&step.Factory{
		F:    steprks,
		Name: "ks",
	})
	step.RegisterStepFlag(InitKSFlag())
}

var KS = cli.Command{
	Name:   "ks",
	Usage:  "k8s 推送",
	Action: steprks,
	Flags:  InitKSFlag(),
}

func InitKSFlag() []cli.Flag {
	flag := []cli.Flag{
		&cli.StringFlag{
			Name:    "ks-namespace",
			Usage:   "namespace 名称",
			EnvVars: []string{"KS_NAMESPACE"},
		},
		&cli.StringFlag{
			Name:    "ks-deployment",
			Usage:   "deployment 名称",
			EnvVars: []string{"KS_DEPLOYMENT"},
		},
	}
	return flag
}

func steprks(c *cli.Context) error {
	fmt.Println("********************************************")
	fmt.Println("*************    ks 阶段开始    *************")
	fmt.Println("********************************************")
	namespace := c.String("ks-namespace")
	if namespace == "" {
		namespace = findNamespace()
	}
	if namespace == "" {
		return errors.New("没有找到 空间")
	}
	dep := findSearchDeployment()
	deploymentName := c.String("ks-deployment")
	if deploymentName != "" {
		dep.Name = deploymentName
	}
	namespaces := namespaceGrouping(namespace)

	errorList := []error{}
	for _, namespace := range namespaces {
		eks := &ExecKubectl{
			Search: &Kubectl{
				Namespace:  namespace,
				Deployment: dep,
			},
		}
		b, a := json.Marshal(eks)
		if a != nil {
			fmt.Println("ks json marshal error:", a)
		}
		fmt.Println("ks json marshal :", string(b))
		e := eks.SearchDeployment()
		if e != nil {
			return e
		}
		error := eks.SetImage(glb.Con.Docker.Tags[0])
		if error != nil {
			errorList = append(errorList, error)
		}

	}
	if len(errorList) > 0 {
		errorString := ""
		for _, e := range errorList {
			fmt.Println(e)
			errorString += e.Error() + "\n"
		}
		return errors.New(errorString)
	}
	return nil
}

func namespaceGrouping(namespace string) []string {

	return strings.Split(namespace, ",")
}
