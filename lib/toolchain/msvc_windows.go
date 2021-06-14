package toolchain

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type MSVCToolchain struct {
	clpath  string
	envvars []string
}

func findVc2017() (string, error) {
	programfiles := os.Getenv("ProgramFiles(x86)")
	if programfiles == "" {
		programfiles = os.Getenv("ProgramFiles")
	}
	if programfiles == "" {
		return "", ErrToolchainNotFound
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

	path := path.Join(strings.TrimSpace(pathBuffer.String()), "VC", "Auxiliary", "Build")
	info, err := os.Stat(path)
	if os.IsNotExist(err) || !info.IsDir() {
		return "", ErrToolchainNotFound
	}
	return path, nil
}

func findVcVarsBat() (string, error) {
	vc, err := findVc2017()
	if err != nil && err != ErrToolchainNotFound {
		return vc, err
	}

	// vc, err = findVc2015()
	// if err == nil || err != ToolchainNotFound {
	// 	return vc, err
	// }

	if err != nil {
		return "", err
	}

	location := path.Join(vc, "vcvars64.bat")
	info, err := os.Stat(location)
	if os.IsNotExist(err) || info.IsDir() {
		return "", ErrToolchainNotFound
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

func NewMSVC() (*MSVCToolchain, error) {
	var result MSVCToolchain
	var err error

	result.envvars, err = getVcVars()
	if err != nil {
		return &result, err
	}
	var paths []string
	for _, v := range result.envvars {
		if len(v) < len(pathEnvVarPrefix) {
			continue
		}
		prefix := v[:len(pathEnvVarPrefix)]
		if strings.EqualFold(prefix, pathEnvVarPrefix) {
			withoutPrefix := v[len(pathEnvVarPrefix):]
			paths = strings.Split(withoutPrefix, string(os.PathListSeparator))
			break
		}

	}
	if paths == nil {
		return nil, ErrToolchainNotFound
	}

	result.clpath, err = findExe("cl.exe", paths)
	if err != nil {
		return &result, err
	}

	return &result, nil
}

func findExe(exe string, paths []string) (string, error) {
	if paths == nil {
		paths = strings.Split(os.Getenv("path"), string(os.PathListSeparator))
	}
	for _, p := range paths {
		abs, err := filepath.Abs(p)
		if err != nil {
			continue
		}
		filename := path.Join(abs, exe)

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
	return "", ErrToolchainNotFound
}

func (tc *MSVCToolchain) Build(mainFullPath string) (string, error) {
	workingDir := path.Dir(mainFullPath)
	main := path.Base(mainFullPath)
	binary := strings.TrimSuffix(main, filepath.Ext(main)) + ".exe"

	cmd := exec.Command(tc.clpath, main, "/link", "/out:"+binary)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Dir = workingDir
	cmd.Env = append(cmd.Env, tc.envvars...)
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return path.Join(workingDir, binary), nil
}
