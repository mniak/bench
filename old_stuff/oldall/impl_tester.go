package oldall

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/andreyvit/diff"
	"github.com/mniak/bench/newcore"
	"github.com/pkg/errors"
)

type (
	_BaseTester  struct{}
	_StartedTest struct {
		stdin  *bytes.Buffer
		stdout *bytes.Buffer
		stderr *bytes.Buffer
		cmd    newcore.StartedCmd

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

func (s *_StartedTest) Wait() error {
	return s.cmd.Wait()
}

func (s *_StartedTest) ExpectedOutput() string {
	return s.expectedOutput
}

func NewTester() Tester {
	return new(_BaseTester)
}

func (t *_BaseTester) Start(test Test) (StartedTest, error) {
	var started _StartedTest
	started.stdin = new(bytes.Buffer)
	started.stdout = new(bytes.Buffer)
	started.stderr = new(bytes.Buffer)
	started.expectedOutput = test.ExpectedOutput

	_, err := started.stdin.WriteString(test.Input)
	if err != nil {
		return &started, err
	}

	runner, err := newcore.RunnerFor(test.Program)
	if err != nil {
		return nil, err
	}

	cmd := newcore.Cmd{
		Path:   test.Program,
		Stdin:  started.stdin,
		Stdout: started.stdout,
		Stderr: started.stderr,
	}

	started.cmd, err = runner.Start(cmd)
	if err != nil {
		err = errors.Wrap(err, "program start failed")
	}
	return &started, err
}

func (t *_BaseTester) Wait(started StartedTest) (result TestResult, err error) {
	err = started.Wait()
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