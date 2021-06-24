package bench

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/andreyvit/diff"
	"github.com/mniak/bench/lib/toolchain"
	"github.com/pkg/errors"
)

type Tester interface {
	Start(t Test) (started StartedTest, err error)
	Wait(started StartedTest) (result TestResult, err error)
}

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

type _TesterWithProgramFinder struct {
	Tester
	programFinder FileFinder
}

func (t *_TesterWithProgramFinder) Start(test Test) (started StartedTest, err error) {
	test.Program, err = t.programFinder.Find(test.Program)
	if err != nil {
		return
	}
	return t.Tester.Start(test)
}

func WrapTesterWithProgramFinder(tester Tester, programFinder FileFinder) Tester {
	return &_TesterWithProgramFinder{
		Tester:        tester,
		programFinder: programFinder,
	}
}

type _TesterWithBuilder struct {
	Tester
	builder    Builder
	toolchains []toolchain.Toolchain
}

func (t *_TesterWithBuilder) Start(test Test) (started StartedTest, err error) {
	test.Program, err = buildIfHasSource(t.builder, t.toolchains, test.Program)
	if err != nil {
		return
	}
	return t.Start(test)
}

func WrapTesterWithBuilder(tester Tester, builder Builder) Tester {
	return &_TesterWithBuilder{
		Tester:  tester,
		builder: builder,
	}
}

func findPlausibleSourceExtensions(toolchains []toolchain.Toolchain, filename string) []string {
	fileExt := filepath.Ext(filename)
	for _, tc := range toolchains {
		if fileExt == tc.OutputExtension() {
			return tc.InputExtensions()
		}
	}
	return []string{}
}

func changeExtension(filename, newExtension string) string {
	ext := filepath.Ext(filename)
	filenameWithoutExtension := strings.TrimSuffix(filename, ext)
	return filenameWithoutExtension + newExtension
}

func findPlausibleSource(toolchains []toolchain.Toolchain, filename string) (string, error) {
	extensions := findPlausibleSourceExtensions(toolchains, filename)
	for _, ext := range extensions {
		filenameWithNewExtension := changeExtension(filename, ext)
		_, err := os.Stat(filenameWithNewExtension)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return "", err
		}
	}
	return "", nil
}

func buildIfHasSource(builder Builder, toolchains []toolchain.Toolchain, filename string) (string, error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return filename, nil
	}
	source, err := findPlausibleSource(toolchains, filename)
	if err != nil {
		return "", err
	}

	if source == "" {
		return "", err
	}

	return builder.Build(filename)
}
