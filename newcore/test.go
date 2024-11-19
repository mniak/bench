package newcore

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/andreyvit/diff"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
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
		Wait() error

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

type _StartedTest struct {
	stdin  *bytes.Buffer
	stdout *bytes.Buffer
	stderr *bytes.Buffer
	cmd    Waiter

	expectedOutput string
}

func StartTest(t Test) (StartedTest, error) {
	prog, err := FindProgram(t.Program)
	cobra.CheckErr(err)
	if prog == nil {
		cobra.CheckErr("could not find the specified test")
	}

	var stdin bytes.Buffer
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	runArgs := RunArgs{
		Stdin:  &stdin,
		Stdout: &stdout,
		Stderr: &stderr,
	}
	var run Waiter
	run, err = prog.Run(runArgs)
	if err != nil {
		return nil, err
	}
	stdin.WriteString(t.Input)
	started := _StartedTest{
		cmd:    run,
		stdin:  &stdin,
		stdout: &stdout,
		stderr: &stderr,
	}
	return &started, nil
}

func (s *_StartedTest) Stdin() *bytes.Buffer {
	return s.stdin
}

func (s *_StartedTest) Stdout() *bytes.Buffer {
	return s.stdout
}

func (s *_StartedTest) Stderr() *bytes.Buffer {
	return s.stderr
}

func (s *_StartedTest) Wait() error {
	return s.cmd.Wait()
}

func (s *_StartedTest) ExpectedOutput() string {
	return s.expectedOutput
}

func WaitTest(t StartedTest) (result TestResult, err error) {
	err = t.Wait()
	result.Output = t.Stdout().String()
	result.ErrorOutput = t.Stderr().String()
	if err != nil {
		exitErr := new(exec.ExitError)
		if errors.As(err, &exitErr) {
			err = fmt.Errorf("program failed with status code %d", exitErr.ExitCode())
		} else {
			err = errors.Wrap(err, "program wait failed")
		}
		return
	}

	if strings.Compare(t.ExpectedOutput(), result.Output) != 0 {
		err = fmt.Errorf("output expectation not satisfied\n%s",
			diff.LineDiff(t.ExpectedOutput(), result.Output),
		)
		return
	}
	return
}
