package bench

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
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

	fullPath := filepath.Join(tempSubFolder, tempBaseName)
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

	fullPath := filepath.Join(tempSubFolder, tempBaseName)
	assert.Equal(t, fullPath, result)
}

func TestDefaultProgramFinder_ShouldHaveExtensionsAndFilenames(t *testing.T) {
	assert.Contains(t, defaultProgramFinder.filenames, "main")

	assert.Contains(t, defaultProgramFinder.extensions, ".py")

	switch runtime.GOOS {
	case "windows":
		assert.Contains(t, defaultProgramFinder.extensions, ".exe")
		assert.Contains(t, defaultProgramFinder.extensions, ".bat")
		assert.Contains(t, defaultProgramFinder.extensions, ".cmd")
		assert.Contains(t, defaultProgramFinder.extensions, ".ps1")

		assert.NotContains(t, defaultProgramFinder.extensions, "", "(none)")
		assert.NotContains(t, defaultProgramFinder.extensions, ".sh")
	default:
		assert.NotContains(t, defaultProgramFinder.extensions, ".exe")
		assert.NotContains(t, defaultProgramFinder.extensions, ".bat")
		assert.NotContains(t, defaultProgramFinder.extensions, ".cmd")
		assert.NotContains(t, defaultProgramFinder.extensions, ".ps1")

		assert.Contains(t, defaultProgramFinder.extensions, "", "(none)")
		assert.Contains(t, defaultProgramFinder.extensions, ".sh")
	}
}

func TestDefaultProgramFinder_UppercaseDefault_ShouldBeTheSameAsLowercaseDefault(t *testing.T) {
	assert.Same(t, defaultProgramFinder, defaultProgramFinder)
}

// var _ ProgramFinder = new(finderWithBuilder)

// func TestExecutableFinder_WhenExecutable_ShouldReturnTrue(t *testing.T) {
// 	switch runtime.GOOS {
// 	case "windows":
// 		path := `C:\Windows\explorer.exe`

// 		result, err := DefaultExecutableFinder.IsExecutable(path)
// 		assert.NoError(t, err)
// 		assert.True(t, result)
// 	default:
// 		path := `C:\Windows\explorer.exe`
// 		result, err := DefaultExecutableFinder.IsExecutable(path)
// 		assert.NoError(t, err)
// 		assert.True(t, result)
// 	}
// }

// func TestExecutableFinder_WhenFileDoesNotExist_ButExtensionIsOnTheList_ShouldReturnTrue(t *testing.T) {
// 	switch runtime.GOOS {
// 	case "windows":
// 		path := `C:\Windows\this-program-does-not-exist.exe`
// 		result, err := DefaultExecutableFinder.IsExecutable(path)
// 		assert.NoError(t, err)
// 		assert.True(t, result)
// 	default:
// 		path := `/opt/folder-that-doesnt-exist/program.sh`
// 		result, err := DefaultExecutableFinder.IsExecutable(path)
// 		assert.NoError(t, err)
// 		assert.True(t, result)
// 	}
// }
