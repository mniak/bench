package impl

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/andreyvit/diff"
	"github.com/mniak/bench/domain"
	"github.com/pkg/errors"
)

type (
	_BaseTester  struct{}
	_StartedTest struct {
		stdin  *bytes.Buffer
		stdout *bytes.Buffer
		stderr *bytes.Buffer
		cmd    *exec.Cmd

		expectedOutput string
	}
)

func (s *_StartedTest) Stdin() *bytes.Buffer {
	return s.stdin
}

func (s *_StartedTest) Stdout() *bytes.Buffer {
	return s.stdout
}

func (s *_StartedTest) Stderr() *bytes.Buffer {
	return s.stderr
}

func (s *_StartedTest) Cmd() *exec.Cmd {
	return s.cmd
}

func (s *_StartedTest) ExpectedOutput() string {
	return s.expectedOutput
}

func NewTester() domain.Tester {
	return new(_BaseTester)
}

func (t *_BaseTester) Start(test domain.Test) (domain.StartedTest, error) {
	var started _StartedTest
	started.stdin = new(bytes.Buffer)
	started.stdout = new(bytes.Buffer)
	started.stderr = new(bytes.Buffer)
	started.expectedOutput = test.ExpectedOutput

	_, err := started.stdin.WriteString(test.Input)
	if err != nil {
		return &started, err
	}

	started.cmd = exec.Command(test.Program)
	started.cmd.Stdin = started.stdin
	started.cmd.Stdout = started.stdout
	started.cmd.Stderr = started.stderr

	err = started.cmd.Start()
	if err != nil {
		err = errors.Wrap(err, "program start failed")
	}
	return &started, err
}

func (t *_BaseTester) Wait(started domain.StartedTest) (result domain.TestResult, err error) {
	err = started.Cmd().Wait()
	if err != nil {
		err = errors.Wrap(err, "program wait failed")
		return
	}

	result.Output = started.Stdout().String()
	result.ErrorOutput = started.Stderr().String()

	if strings.Compare(started.ExpectedOutput(), result.Output) != 0 {
		err = fmt.Errorf("output expectation not satisfied\n%s",
			diff.LineDiff(started.ExpectedOutput(), result.Output),
		)
		return
	}
	return
}
