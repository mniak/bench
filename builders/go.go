package builders

import (
	"bytes"
	"os/exec"
	"path/filepath"
)

func NewGoBuilder() Loader {
	return &GoBuilder{}
}

type GoBuilder struct{}

func (py *GoBuilder) Name() string {
	return "Go"
}

func (py *GoBuilder) Load() (Builder, error) {
	var buffer bytes.Buffer
	cmd := exec.Command("go", "version")
	cmd.Stdout = &buffer

	err := cmd.Run()
	if err != nil {
		return nil, ErrBuilderNotFound
	}

	return py, nil
}

func (py *GoBuilder) CanRun(filename string) bool {
	extension := filepath.Ext(filename)
	switch extension {
	case ".py":
		return true
	default:
		return false
	}
}

func (py *GoBuilder) Start(cmd Cmd) (StartedCmd, error) {
	args := append([]string{cmd.Path}, cmd.Args...)
	c := exec.Command("go", args...)
	c.Stdin = cmd.Stdin
	c.Stdout = cmd.Stdout
	c.Stderr = cmd.Stderr
	err := c.Start()
	if err != nil {
		return nil, err
	}
	return newStartedBuilderCmd(c), nil
}
