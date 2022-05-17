package java

import (
	"build-tools/exec"
	"github.com/urfave/cli/v2"
)

var MavenCommand = cli.Command{
	Name:   "maven",
	Usage:  "java maven 编译",
	Action: maven,
	Flags:  InitMavenFlag(),
}

func InitMavenFlag() []cli.Flag {
	flag := []cli.Flag{
		&cli.StringFlag{
			Name:    "maven-option",
			Usage:   "java-maven 选项",
			EnvVars: []string{"MAVEN_OPTION"},
		},
		&cli.StringFlag{
			Name:    "maven-profiles",
			Usage:   "java-maven profiles",
			EnvVars: []string{"MAVEN_PROFILES"},
		},
	}
	return flag
}

type Maven struct {
	option   string
	profiles string
}

func (m *Maven) Install() error {
	var option []string
	if m.option != "" {
		option = append(option, m.option)
	}
	option = append(option, "install")
	if m.profiles != "" {
		option = append(option, "-P", m.profiles)
	}
	return exec.ExecCommand("mvn", option)
}

func maven(c *cli.Context) error {
	maven := Maven{
		option:   c.String("maven-option"),
		profiles: c.String("maven-profiles"),
	}
	return maven.Install()
}
