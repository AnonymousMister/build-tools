package node

import "build-tools/exec"

type node struct {
	commandName string
	config      map[string]string
}

func (d *node) SetConfigs() error {
	for k, v := range d.config {
		if e := d.SetConfig(k, v); e != nil {
			return e
		}
	}
	return nil
}

func (d *node) SetConfig(key, value string) error {
	_, e := exec.ExecCommand(d.commandName, []string{
		"config", "set", key, value,
	})
	return e
}

func (d *node) Install(Options ...string) error {
	_, e := exec.ExecCommand(d.commandName, append([]string{"install"}, Options...))
	return e
}

func (d *node) Run(script string) error {
	_, e := exec.ExecCommand(d.commandName, []string{
		"run", script,
	})
	return e
}
