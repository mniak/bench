package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/mniak/bench/old_stuff/lib/bench"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Build_WithFolderName_AndRunToCheckIfBuiltCorrectly(t *testing.T) {
	expectedExitCode := rand.Intn(200) + 1
	sourceCode := fmt.Sprintf(`
#include <string>
#include <iostream>

using namespace std;

int main() {
    cout << "Hello, world (%d)!";
	return %d;
}`, expectedExitCode, expectedExitCode)

	tempdir, err := os.MkdirTemp("", "build_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempdir)

	inputPath := filepath.Join(tempdir, "main.cpp")
	err = os.WriteFile(inputPath, []byte(sourceCode), 0o644)
	require.NoError(t, err)

	_, err = bench.Build(tempdir)
	require.NoError(t, err)

	outputPath := filepath.Join(tempdir, "main"+OSBinaryExtension)
	assert.FileExists(t, outputPath)

	var stdout bytes.Buffer
	cmd := exec.Command(outputPath)
	cmd.Stdout = &stdout

	err = cmd.Run()
	require.IsType(t, &exec.ExitError{}, err)
	_, ok := err.(*exec.ExitError)
	require.True(t, ok)
	require.Equal(t, expectedExitCode, cmd.ProcessState.ExitCode())

	require.Equal(t, fmt.Sprintf("Hello, world (%d)!", expectedExitCode), stdout.String())
}
