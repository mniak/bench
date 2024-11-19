package newcore

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
)

func FindProgram(dirname string) (*Program, error) {
	entries, err := os.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	knownCompilerExtensions := make(map[string]Compiler)
	knownRunnerExtensions := make(map[string]Runner)

	for _, toolchain := range Toolchains() {
		if compiler, ok := toolchain.(Compiler); ok {
			exts := compiler.CompilerInputExtensions()
			for _, ext := range exts {
				knownCompilerExtensions[ext] = compiler
			}
		}
		if runner, ok := toolchain.(Runner); ok {
			exts := runner.RunnerInputExtensions()
			for _, ext := range exts {
				knownRunnerExtensions[ext] = runner
			}
		}
	}

	for _, file := range entries {
		if file.IsDir() {
			continue
		}
		ext := filepath.Ext(file.Name())
		compiler, hasCompileExtension := knownCompilerExtensions[ext]
		runner, hasRunnerExtension := knownRunnerExtensions[ext]
		if hasRunnerExtension || hasCompileExtension {
			fullpath := filepath.Join(dirname, file.Name())
			prog := Program{
				Program: fullpath,
			}
			if hasRunnerExtension && runner.CanRun(fullpath) {
				prog.Runner = runner
			}
			if hasCompileExtension && compiler.CanCompile(fullpath) {
				prog.Compiler = compiler
			}
			return &prog, nil
		}
	}

	return nil, nil

	// for _, dir := range entries {
	// 	if dir.IsDir() {
	// 		continue
	// 	}
	// }
}

type Program struct {
	Program  string
	Compiler Compiler
	Runner   Runner
}

func (p Program) Run(a RunArgs) (waiter Waiter, err error) {
	if p.Runner != nil {
		return p.Runner.Start(p.Program, a)
	}
	if p.Compiler != nil {
		temp, err := os.CreateTemp("", "test_*.exe")
		if err != nil {
			return nil, err
		}
		temp.Close()
		defer func() {
			if err != nil {
				temp.Close()
			}
		}()

		var compileOut bytes.Buffer
		var compileErr bytes.Buffer

		err = p.Compiler.Compile(CompilationInput{
			Stdout:         &compileOut,
			Stderr:         &compileErr,
			Filename:       p.Program,
			OutputFilename: temp.Name(),
		})
		if err != nil {
			io.Copy(os.Stdout, &compileOut)
			io.Copy(os.Stderr, &compileErr)
			return nil, err
		}

		binr := BinaryRunner()
		started, err := binr.Start(temp.Name(), a)
		if err != nil {
			return nil, err
		}

		return callbackWaiter{
			waiter: started,
			callback: func() {
				defer os.Remove(temp.Name())
			},
		}, nil

	}
	return nil, errors.New("invalid program found. must have a compiler or a runner attached.")
}

type callbackWaiter struct {
	waiter   Waiter
	callback func()
}

func (w callbackWaiter) Wait() error {
	defer func() {
		if w.callback != nil {
			w.callback()
		}
	}()
	return w.waiter.Wait()
}
