package impl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/bench/internal/mocks"
	"github.com/mniak/bench/internal/utils"
	"github.com/mniak/bench/toolchain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_WhenFilenameHasRunnableExtension_AndSourceExists_ShouldReturnToolchain(t *testing.T) {
	BaseName := gofakeit.Word()
	ExtRun := "." + gofakeit.FileExtension()
	ExtSource := "." + gofakeit.FileExtension()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tloader := mocks.NewMockToolchainLoader(ctrl)

	tloader.EXPECT().
		InputExtensions().
		Return([]string{ExtSource})

	tloader.EXPECT().
		OutputExtension().
		Return(ExtRun)

	tempdir, err := os.MkdirTemp("", "test_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempdir)

	// when filename has runnable extension
	runnable := filepath.Join(tempdir, BaseName+ExtRun)

	// and source exists
	source, err := utils.TempFile(tempdir, BaseName+ExtSource)
	require.NoError(t, err)
	defer source.CloseAndRemove()

	// should return toolchain
	sut := NewToolchainFinderFromToolchainLoaders(tloader)
	result, err := sut.Find(runnable)
	require.NoError(t, err)
	assert.Equal(t, tloader, result)
}

func Test_WhenFilenameHasRunnableExtension_AndSourceDoesNotExist_ShouldErrorNotFound(t *testing.T) {
	BASENAME := gofakeit.Word()
	EXT_RUN := "." + gofakeit.FileExtension()
	EXT_SOURCE := "." + gofakeit.FileExtension()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tloader := mocks.NewMockToolchainLoader(ctrl)

	tloader.EXPECT().
		InputExtensions().
		Return([]string{EXT_SOURCE})

	tloader.EXPECT().
		OutputExtension().
		Return(EXT_RUN)

	tempdir, err := os.MkdirTemp("", "test_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempdir)

	// when filename has runnable extension
	runnable := filepath.Join(tempdir, BASENAME+EXT_RUN)

	// should return error: Not Found
	sut := NewToolchainFinderFromToolchainLoaders(tloader)
	_, err = sut.Find(runnable)
	assert.Same(t, toolchain.ErrToolchainNotFound, err)
}

func Test_WhenFilenameHasSourceExtension_AndItExists_ShouldReturnToolchain(t *testing.T) {
	BASENAME := gofakeit.Word()
	EXT_RUN := "." + gofakeit.FileExtension()
	EXT_SOURCE := "." + gofakeit.FileExtension()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tloader := mocks.NewMockToolchainLoader(ctrl)

	tloader.EXPECT().
		InputExtensions().
		Return([]string{EXT_SOURCE})

	tloader.EXPECT().
		OutputExtension().
		Return(EXT_RUN)

	tempdir, err := os.MkdirTemp("", "test_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempdir)

	// when filename has source extension
	filename := BASENAME + EXT_SOURCE

	// and it exists
	file, err := utils.TempFile(tempdir, filename)
	require.NoError(t, err)
	defer file.CloseAndRemove()

	// should return toolchain
	sut := NewToolchainFinderFromToolchainLoaders(tloader)
	result, err := sut.Find(file.Name())
	require.NoError(t, err)
	assert.Equal(t, tloader, result)
}

func Test_WhenFilenameHasSourceExtension_AndItDoesNotExist_ShouldReturnToolchain(t *testing.T) {
	BASENAME := gofakeit.Word()
	EXT_RUN := "." + gofakeit.FileExtension()
	EXT_SOURCE := "." + gofakeit.FileExtension()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tloader := mocks.NewMockToolchainLoader(ctrl)

	tloader.EXPECT().
		InputExtensions().
		Return([]string{EXT_SOURCE})

	tloader.EXPECT().
		OutputExtension().
		Return(EXT_RUN)

	tempdir, err := os.MkdirTemp("", "test_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempdir)

	// when filename has source extension
	// and it does not exist
	filename := filepath.Join(tempdir, BASENAME+EXT_SOURCE)

	// should return toolchain
	sut := NewToolchainFinderFromToolchainLoaders(tloader)
	_, err = sut.Find(filename)
	assert.Same(t, toolchain.ErrToolchainNotFound, err)
}

func Test_WhenFilenameHasUnknownExtension_AndItExists_ShouldReturnToolchain(t *testing.T) {
	BASENAME := gofakeit.Word()
	EXT_RUN := "." + gofakeit.FileExtension()
	EXT_SOURCE := "." + gofakeit.FileExtension()
	EXT_UNKNOWN := "." + gofakeit.FileExtension()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tloader := mocks.NewMockToolchainLoader(ctrl)

	tloader.EXPECT().
		InputExtensions().
		Return([]string{EXT_SOURCE})

	tloader.EXPECT().
		OutputExtension().
		Return(EXT_RUN)

	tempdir, err := os.MkdirTemp("", "test_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempdir)

	// when filename has unknown extension
	filename := BASENAME + EXT_UNKNOWN

	// and it exists
	file, err := utils.TempFile(tempdir, filename)
	require.NoError(t, err)
	defer file.CloseAndRemove()

	// should return toolchain
	sut := NewToolchainFinderFromToolchainLoaders(tloader)
	_, err = sut.Find(filename)
	assert.Same(t, toolchain.ErrToolchainNotFound, err)
}

func Test_WhenFilenameHasUnknownExtension_AndItDoesNotExist_ShouldReturnToolchain(t *testing.T) {
	BASENAME := gofakeit.Word()
	EXT_RUN := "." + gofakeit.FileExtension()
	EXT_SOURCE := "." + gofakeit.FileExtension()
	EXT_UNKNOWN := "." + gofakeit.FileExtension()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tloader := mocks.NewMockToolchainLoader(ctrl)

	tloader.EXPECT().
		InputExtensions().
		Return([]string{EXT_SOURCE})

	tloader.EXPECT().
		OutputExtension().
		Return(EXT_RUN)

	tempdir, err := os.MkdirTemp("", "test_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempdir)

	// when filename has unknown extension
	// and it does not exist
	filename := filepath.Join(tempdir, BASENAME+EXT_UNKNOWN)

	// should return toolchain
	sut := NewToolchainFinderFromToolchainLoaders(tloader)
	_, err = sut.Find(filename)
	assert.Same(t, toolchain.ErrToolchainNotFound, err)
}
