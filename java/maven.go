package java

import (
	"github.com/urfave/cli/v2"

	"github.com/AnonymousMister/build-tools/exec"
)

var MavenCommand = cli.Command{
	Name:   "java-maven",
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
	option string
	profiles string
}

func maven(c *cli.Context) error {
	maven := Maven{
		option: c.String("maven-option"),
		profiles:  c.String("maven-profiles"),
	}
	var option []string
	if maven.option != ""	{
		option = append(option, maven.option)
	}
	option = append(option, "install")
	if maven.profiles != ""	{
		option = append(option,"-P", maven.profiles)
	}
	exec.ExecCommand("mvn", option)
	return nil
}
