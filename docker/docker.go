package docker

import (
	"build-tools/exec"
	"build-tools/glb"
	"build-tools/step"
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
)

func init() {
	step.RegisterStepmap(&step.Factory{
		F:    steprDocker,
		Name: "docker",
	})
	step.RegisterStepFlag(InitDockerFlag())
}

var DockerCommand = cli.Command{
	Name:   "docker",
	Usage:  "docker 阶段 命令",
	Action: steprDocker,
	Flags:  InitDockerFlag(),
}

func InitDockerFlag() []cli.Flag {
	flag := []cli.Flag{
		&cli.StringFlag{
			Name:    "docker-config-file",
			Usage:   "docker 配置文件 路径",
			EnvVars: []string{"DOCKER_CONFIG_FILE"},
		},
		&cli.StringFlag{
			Name:    "docker-url",
			Usage:   "docker 仓库域名",
			EnvVars: []string{"DOCKER_REGISTRY_URL"},
		},
		&cli.StringFlag{
			Name:    "docker-namespace",
			Usage:   "docker 命名空间",
			EnvVars: []string{"DOCKER_REGISTRY_NAMESPACE"},
		},
		&cli.StringFlag{
			Name:    "docker-tag",
			Usage:   "docker 打包标签",
			EnvVars: []string{"DOCKER_REGISTRY_TAG"},
		},
	}
	return flag
}

func steprDocker(c *cli.Context) error {
	fmt.Println("********************************************")
	fmt.Println("***********    docker 阶段开始    ***********")
	fmt.Println("********************************************")
	tags := strings.Split(c.String("docker-tag"), ",")
	dockerRegistry := c.String("docker-url")
	dockerNamespace := c.String("docker-namespace")
	if dockerNamespace != "" {
		dockerRegistry = dockerRegistry + "/" + dockerNamespace
	}
	var option []string
	configFile := c.String("docker-config-file")
	if configFile != "" {
		option = append(option, "--config="+configFile)
	}
	commandName := "docker"
	if glb.IsDebug {
		commandName = "podman"
	}

	dockerContext := &glb.DockerContext{
		DockerRegistry: dockerRegistry,
		Tags:           tags,
	}
	glb.Con.Docker = dockerContext
	docker := &Docker{
		DockerRegistry: dockerRegistry,
		Tags:           tags,
		Option:         option,
		commandName:    commandName,
	}
	err := docker.BuildX()
	if err != nil {
		return err
	}
	err = docker.Rmi()
	if err != nil {
		return err
	}
	return nil
}

type Docker struct {
	DockerRegistry string
	Tags           []string
	Option         []string
	commandName    string
}

func (d *Docker) Build() error {
	var options []string
	if d.Option != nil && len(d.Option) > 0 {
		options = append(options, d.Option...)
	}
	options = append(options, "build", "-t", d.DockerRegistry+":"+d.Tags[0], ".")
	_, err := exec.ExecCommand(d.commandName, options)
	return err
}

func (d *Docker) Tag() error {
	var options []string
	if d.Option != nil && len(d.Option) > 0 {
		options = append(options, d.Option...)
	}
	for i, tag := range d.Tags {
		if i == 0 {
			continue
		}
		params := append(options, "tag", d.DockerRegistry+":"+d.Tags[0], d.DockerRegistry+":"+tag)
		_, err := exec.ExecCommand(d.commandName, params)
		if err != nil {
			return err
		}
	}
	return nil
}
func (d *Docker) Push() error {
	var options []string
	if d.Option != nil && len(d.Option) > 0 {
		options = append(options, d.Option...)
	}
	for _, tag := range d.Tags {
		params := append(options, "push", d.DockerRegistry+":"+tag)
		_, err := exec.ExecCommand(d.commandName, params)
		if err != nil {
			return err
		}
	}
	return nil
}
func (d *Docker) Rmi() error {
	var options []string
	if d.Option != nil && len(d.Option) > 0 {
		options = append(options, d.Option...)
	}
	options = append(options, "rmi")
	for _, tag := range d.Tags {
		options = append(options, d.DockerRegistry+":"+tag)
	}
	_, err := exec.ExecCommand(d.commandName, options)
	if err != nil {
		return err
	}
	return nil
}
