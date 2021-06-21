package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSplitProgramDir_WhenFile(t *testing.T) {
	tempDir := os.TempDir()
	file, err := os.CreateTemp(tempDir, "test_*")
	require.NoError(t, err, "create temp file")
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	dir, program, err := SplitDirAndProgram(file.Name())
	require.NoError(t, err, "split file and dir")

	assert.Equal(t, tempDir, dir)
	assert.Equal(t, filepath.Base(file.Name()), program)
}

func TestSplitProgramDir_WhenFolder(t *testing.T) {
	tempDir := os.TempDir()
	folder, err := os.MkdirTemp(tempDir, "test_*")
	require.NoError(t, err, "create temp file")
	defer func() {
		os.Remove(folder)
	}()

	dir, program, err := SplitDirAndProgram(folder)
	require.NoError(t, err, "split file and dir")

	assert.Equal(t, folder, dir)
	assert.Equal(t, "", program)
}
