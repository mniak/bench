package bench

import (
	"os"
	"path"
	"path/filepath"
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

func TestProgramFinder_WhenFolder_ShouldFindFilenameWithExtension(t *testing.T) {
	tempDir := os.TempDir()
	tempSubFolder, err := os.MkdirTemp(tempDir, "test_*")
	require.NoError(t, err, "create temp dir")
	defer os.RemoveAll(tempSubFolder)

	tempName := gofakeit.Word()
	tempExtension := "." + gofakeit.Word()
	tempBaseName := tempName + tempExtension

	file, err := os.Create(path.Join(tempSubFolder, tempBaseName))
	require.NoError(t, err, "create temp file")
	defer file.Close()

	finder := _ProgramFinder{
		filenames:  []string{tempName},
		extensions: []string{tempExtension},
	}
	result, err := finder.Find(tempSubFolder)
	assert.NoError(t, err)

	fullPath := path.Join(tempSubFolder, tempBaseName)
	assert.Equal(t, fullPath, result)
}

func TestProgramFinder_WhenFolder_ShouldFindFolderNameWithExtension(t *testing.T) {
	tempDir := os.TempDir()
	tempSubFolder, err := os.MkdirTemp(tempDir, "test_*")
	require.NoError(t, err, "create temp dir")
	defer os.RemoveAll(tempSubFolder)

	folderBaseName := filepath.Base(tempSubFolder)
	tempExtension := "." + gofakeit.Word()
	tempBaseName := folderBaseName + tempExtension

	file, err := os.Create(path.Join(tempSubFolder, tempBaseName))
	require.NoError(t, err, "create temp file")
	defer file.Close()

	finder := _ProgramFinder{
		filenames:  []string{gofakeit.Word()},
		extensions: []string{tempExtension},
	}
	result, err := finder.Find(tempSubFolder)
	assert.NoError(t, err)

	fullPath := path.Join(tempSubFolder, tempBaseName)
	assert.Equal(t, fullPath, result)
}
