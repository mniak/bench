package bench

import (
	"bytes"
	"fmt"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/andreyvit/diff"
	"github.com/pkg/errors"
)

type Test struct {
	Program    string
	Args       []string
	WorkingDir string

	Input          string
	ExpectedOutput string

	stdin  *bytes.Buffer
	stdout *bytes.Buffer
	stderr *bytes.Buffer
	cmd    *exec.Cmd
}

type TestResult struct {
	Output      string
	ErrorOutput string
}

func (t *Test) Start() error {
	t.stdin = new(bytes.Buffer)
	t.stdout = new(bytes.Buffer)
	t.stderr = new(bytes.Buffer)

	program := t.Program
	if !path.IsAbs(t.Program) {
		var err error
		program, err = filepath.Abs(path.Join(t.WorkingDir, t.Program))
		if err != nil {
			return err
		}
	}

	t.cmd = exec.Command(program, t.Args...)
	t.cmd.Dir = t.WorkingDir
	t.cmd.Stdin = t.stdin
	t.cmd.Stdout = t.stdout
	t.cmd.Stderr = t.stderr

	err := t.cmd.Start()
	if err != nil {
		return errors.Wrap(err, "program start failed")
	}
	return nil
}

func (t *Test) Wait() (TestResult, error) {
	var result TestResult

	t.stdin.WriteString(t.Input)
	err := t.cmd.Wait()
	if err != nil {
		return result, errors.Wrap(err, "program wait failed")
	}

	result.Output = t.stdout.String()
	result.ErrorOutput = t.stderr.String()

	if strings.Compare(result.Output, t.ExpectedOutput) != 0 {
		return result, fmt.Errorf("output expectation not satisfied\n%s",
			diff.LineDiff(t.ExpectedOutput, t.stdout.String()),
		)
	}
	return result, nil
}
