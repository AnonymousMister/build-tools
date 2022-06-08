package docker

import "build-tools/exec"

func (d *Docker) BuildX() error {
	var options []string
	if d.Option != nil && len(d.Option) > 0 {
		options = append(options, d.Option...)
	}
	options = append(options, "buildx", "build")

	/* 拼接 版本号*/
	for _, tag := range d.Tags {
		options = append(options, "-t", d.DockerRegistry+":"+tag)
	}
	options = append(options, "--push", ".")
	_, err := exec.ExecCommand(d.commandName, options)
	return err
}
