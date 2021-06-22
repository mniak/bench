package bench

import (
	"os"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProgramFinder_WhenCommand(t *testing.T) {
	sentence := gofakeit.Sentence(5)

	finder := _ProgramFinder{}
	program, err := finder.Find(sentence)
	assert.NoError(t, err)
	assert.Equal(t, sentence, program)
}

func TestProgramFinder_WhenFolder_ShouldFindMain(t *testing.T) {
	tempDir := os.TempDir()
	tempSubFolder, err := os.MkdirTemp(tempDir, "test_*")
	require.NoError(t, err, "create temp dir")
	defer os.RemoveAll(tempSubFolder)

	file, err := os.Create(path.Join(tempSubFolder, "main.exe"))
	require.NoError(t, err, "create temp file")
	defer file.Close()

	finder := _ProgramFinder{}
	result, err := finder.Find(tempSubFolder)
	assert.NoError(t, err)

	fullPath := path.Join(tempSubFolder, "main.exe")
	assert.Equal(t, fullPath, result)
}
