package newcore

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

type MSVCLoader struct{}

func (l *MSVCLoader) Load() (Toolchain, error) {
	var result _MSVCToolchain
	var err error

	result.Env, err = getVcVars()
	if err != nil {
		return &result, err
	}
	var paths []string
	for _, v := range result.Env {
		if strings.HasPrefix(strings.ToUpper(v), pathEnvVarPrefix) {
			withoutPrefix := v[len(pathEnvVarPrefix):]
			paths = strings.Split(withoutPrefix, string(os.PathListSeparator))
			break
		}
	}
	if paths == nil {
		return nil, errors.New("toolchain not loaded: MSVC not found")
	}

	result.CLPath, err = findBin("cl", paths...)
	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (l *MSVCLoader) ToolchainType() reflect.Type {
	return reflect.TypeOf(_MSVCToolchain{})
}

type _MSVCToolchain struct {
	Env    []string
	CLPath string
}

var ErrMSVCNotFound = errors.New("MSVC not found")

func (tc *_MSVCToolchain) CanCompile(filename string) bool {
	extension := filepath.Ext(filename)
	switch strings.ToLower(extension) {
	case ".c", ".cc":
		return true
	case ".cpp", ".c++", ".cxx", ".cp":
		return true
	default:
		return false
	}
}

func (tc *_MSVCToolchain) CompilerInputExtensions() []string {
	return []string{
		".c", ".cc",
		".cpp", ".c++", ".cxx", ".cp",
	}
}

func (tc *_MSVCToolchain) Compile(input CompilationInput) error {
	basename := filepath.Base(input.Filename)
	workingDir := filepath.Dir(input.Filename)
	outputpath, err := filepath.Abs(input.OutputFilename)
	if err != nil {
		return err
	}

	cmd := exec.Command(tc.CLPath, basename, "/link", "/out:"+outputpath)
	cmd.Stderr = input.Stderr
	cmd.Stdout = input.Stdout

	cmd.Dir = workingDir
	cmd.Env = append(cmd.Env, tc.Env...)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func findVc2017() (string, error) {
	programfiles := os.Getenv("ProgramFiles(x86)")
	if programfiles == "" {
		programfiles = os.Getenv("ProgramFiles")
	}
	if programfiles == "" {
		return "", ErrMSVCNotFound
	}

	var pathBuffer strings.Builder
	cmd := exec.Command(path.Join(programfiles, "Microsoft Visual Studio", "Installer", "vswhere.exe"),
		"-latest",
		"-prerelease",
		"-requires", "Microsoft.VisualStudio.Component.VC.Tools.x86.x64",
		"-property", "installationPath",
		"-products", "*",
	)
	cmd.Stdout = &pathBuffer
	if err := cmd.Run(); err != nil {
		return "", err
	}

	path := filepath.Join(strings.TrimSpace(pathBuffer.String()), "VC", "Auxiliary", "Build")
	info, err := os.Stat(path)
	if os.IsNotExist(err) || !info.IsDir() {
		return "", ErrMSVCNotFound
	}
	return path, nil
}

func findVcVarsBat() (string, error) {
	vc, err := findVc2017()
	if err != nil && err != ErrMSVCNotFound {
		return vc, err
	}

	// vc, err = findVc2015()
	// if err == nil || err != ToolchainNotFound {
	// 	return vc, err
	// }

	if err != nil {
		return "", err
	}

	location := filepath.Join(vc, "vcvars64.bat")
	info, err := os.Stat(location)
	if os.IsNotExist(err) || info.IsDir() {
		return "", ErrMSVCNotFound
	}

	return location, nil
}

func getVcVars() ([]string, error) {
	vcvarsbat, err := findVcVarsBat()
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(vcvarsbat, "&", "set")
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(&buffer)
	scanner.Split(bufio.ScanLines)
	vars := make([]string, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(string(scanner.Bytes()))
		if strings.ContainsRune(line, '=') {
			vars = append(vars, line)
		}
	}

	return vars, nil
}

const pathEnvVarPrefix = "PATH="

func findBin(exe string, paths ...string) (string, error) {
	var OSBinaryExtension string
	if runtime.GOOS == "windows" {
		OSBinaryExtension = ".exe"
	}

	if len(paths) == 0 {
		paths = strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))
	}
	for _, p := range paths {
		abs, err := filepath.Abs(p)
		if err != nil {
			continue
		}
		filename := filepath.Join(abs, exe+OSBinaryExtension)

		info, err := os.Stat(filename)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return "", err
		}
		if !info.IsDir() {
			return filename, nil
		}

	}
	return "", ErrMSVCNotFound
}
