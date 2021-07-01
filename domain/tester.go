package domain

import (
	"bytes"
	"os/exec"
)

type (
	Test struct {
		Program        string
		Input          string
		ExpectedOutput string
	}
	StartedTest interface {
		Stdin() *bytes.Buffer
		Stdout() *bytes.Buffer
		Stderr() *bytes.Buffer
		Cmd() *exec.Cmd

		ExpectedOutput() string
	}
	TestResult struct {
		Output      string
		ErrorOutput string
	}
)

type Tester interface {
	Start(t Test) (started StartedTest, err error)
	Wait(started StartedTest) (result TestResult, err error)
}
