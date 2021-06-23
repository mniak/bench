package bench

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/andreyvit/diff"
	"github.com/pkg/errors"
)

type Tester interface {
	Start(t Test) (started StartedTest, err error)
	Wait(started StartedTest) (result TestResult, err error)
}

type _Tester struct {
	programFinder ProgramFinder
}

func NewTester(programFinder ProgramFinder) Tester {
	return &_Tester{
		programFinder: programFinder,
	}
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

func (tester *_Tester) Start(t Test) (started StartedTest, err error) {
	started.stdin = new(bytes.Buffer)
	started.stdout = new(bytes.Buffer)
	started.stderr = new(bytes.Buffer)
	started.expectedOutput = t.ExpectedOutput

	program, err := tester.programFinder.Find(t.Program)
	if err != nil {
		return
	}

	_, err = started.stdin.WriteString(t.Input)
	if err != nil {
		return
	}
	started.cmd = exec.Command(program)
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

var DefaultTester Tester = NewTester(defaultProgramFinder)

func StartTest(t Test) (started StartedTest, err error) {
	return DefaultTester.Start(t)
}

func WaitTest(started StartedTest) (result TestResult, err error) {
	return DefaultTester.Wait(started)
}
