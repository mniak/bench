package bench

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/andreyvit/diff"
	"github.com/pkg/errors"
)

type Test struct {
	programFinder ProgramFinder

	program string

	Input          string
	ExpectedOutput string
	Args           []string

	stdin  *bytes.Buffer
	stdout *bytes.Buffer
	stderr *bytes.Buffer
	cmd    *exec.Cmd
}

func NewTest(program, input, expectedOutput string) Test {
	return Test{
		program:        program,
		Input:          input,
		ExpectedOutput: expectedOutput,
		Args:           make([]string, 0),

		programFinder: DefaultProgramFinder,
	}
}

func (t *Test) WithProgramFinder(programFinder ProgramFinder) {
	t.programFinder = programFinder
}

func (t *Test) Start() error {
	t.stdin = new(bytes.Buffer)
	t.stdout = new(bytes.Buffer)
	t.stderr = new(bytes.Buffer)

	program, err := t.programFinder.Find(t.program)
	if err != nil {
		return err
	}

	_, err = t.stdin.WriteString(t.Input)
	if err != nil {
		return err
	}
	t.cmd = exec.Command(program)
	t.cmd.Stdin = t.stdin
	t.cmd.Stdout = t.stdout
	t.cmd.Stderr = t.stderr

	err = t.cmd.Start()
	if err != nil {
		return errors.Wrap(err, "program start failed")
	}
	return nil
}

func (t *Test) Wait() (TestResult, error) {
	var result TestResult

	err := t.cmd.Wait()
	if err != nil {
		return result, errors.Wrap(err, "program wait failed")
	}

	result.Output = t.stdout.String()
	result.ErrorOutput = t.stderr.String()

	if strings.Compare(t.ExpectedOutput, result.Output) != 0 {
		return result, fmt.Errorf("output expectation not satisfied\n%s",
			diff.LineDiff(t.ExpectedOutput, result.Output),
		)
	}
	return result, nil
}

func (t *Test) Program() string {
	return t.program
}

type TestResult struct {
	Output      string
	ErrorOutput string
}
