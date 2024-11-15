package newcore

import (
	"errors"
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

func (p Program) Run(a RunArgs) (StartedProgram, error) {
	if p.Runner != nil {
		return p.Runner.Start(p.Program, a)
	}
	if p.Compiler != nil {
		temp, err := os.CreateTemp("", "test_*.exe")
		if err != nil {
			return nil, err
		}
		// defer os.Remove(temp.Name())
		defer temp.Close()

		err = p.Compiler.Compile(CompilationInput{
			Stdin:          a.Stdin,
			Stdout:         a.Stdout,
			Stderr:         a.Stderr,
			Filename:       p.Program,
			OutputFilename: temp.Name(),
		})
		if err != nil {
			return nil, err
		}

		go run()

		binr := BinaryRunner()
		return binr.Start(temp.Name(), a)

	}
	return nil, errors.New("invalid program found. must have a compiler or a runner attached.")
}
