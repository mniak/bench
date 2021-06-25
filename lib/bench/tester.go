package bench

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/andreyvit/diff"
	"github.com/pkg/errors"
)

type _Tester struct{}

func NewTester() Tester {
	return new(_Tester)
}

type (
	Test struct {
		Program        string
		Input          string
		ExpectedOutput string
	}
	StartedTest struct {
		stdin  *bytes.Buffer
		stdout *bytes.Buffer
		stderr *bytes.Buffer
		cmd    *exec.Cmd

		expectedOutput string
	}
	TestResult struct {
		Output      string
		ErrorOutput string
	}
)

func (t *_Tester) Start(test Test) (started StartedTest, err error) {
	started.stdin = new(bytes.Buffer)
	started.stdout = new(bytes.Buffer)
	started.stderr = new(bytes.Buffer)
	started.expectedOutput = test.ExpectedOutput
	_, err = started.stdin.WriteString(test.Input)
	if err != nil {
		return
	}
	started.cmd = exec.Command(test.Program)
	started.cmd.Stdin = started.stdin
	started.cmd.Stdout = started.stdout
	started.cmd.Stderr = started.stderr

	err = started.cmd.Start()
	if err != nil {
		err = errors.Wrap(err, "program start failed")
	}
	return
}

func (t *_Tester) Wait(started StartedTest) (result TestResult, err error) {
	err = started.cmd.Wait()
	if err != nil {
		err = errors.Wrap(err, "program wait failed")
		return
	}

	result.Output = started.stdout.String()
	result.ErrorOutput = started.stderr.String()

	if strings.Compare(started.expectedOutput, result.Output) != 0 {
		err = fmt.Errorf("output expectation not satisfied\n%s",
			diff.LineDiff(started.expectedOutput, result.Output),
		)
		return
	}
	return
}
