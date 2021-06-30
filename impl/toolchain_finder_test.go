package impl

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/bench/domain"
	"github.com/mniak/bench/internal/mocks"
	"github.com/mniak/bench/internal/utils"
	"github.com/mniak/bench/toolchain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_WhenFilenameHasRunnableExtension_AndSourceExists_ShouldReturnToolchain(t *testing.T) {
	BASENAME := gofakeit.Word()
	EXT_RUN := "." + gofakeit.FileExtension()
	EXT_SOURCE := "." + gofakeit.FileExtension()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tchain := mocks.NewMockToolchain(ctrl)

	tchain.EXPECT().
		InputExtensions().
		Return([]string{EXT_SOURCE})

	tchain.EXPECT().
		OutputExtension().
		Return(EXT_RUN)

	tempdir, err := ioutil.TempDir("", "test_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempdir)

	// when filename has runnable extension
	runnable := filepath.Join(tempdir, BASENAME+EXT_RUN)

	// and source exists
	source, err := utils.TempFile(tempdir, BASENAME+EXT_SOURCE)
	require.NoError(t, err)
	defer source.CloseAndRemove()

	// should return toolchain
	sut := NewToolchainFinderFromToolchains([]domain.Toolchain{
		tchain,
	})

	result, err := sut.Produce(runnable)
	require.NoError(t, err)
	assert.Equal(t, tchain, result)
}

func Test_WhenFilenameHasRunnableExtension_AndSourceDoesNotExist_ShouldErrorNotFound(t *testing.T) {
	BASENAME := gofakeit.Word()
	EXT_RUN := "." + gofakeit.FileExtension()
	EXT_SOURCE := "." + gofakeit.FileExtension()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tchain := mocks.NewMockToolchain(ctrl)

	tchain.EXPECT().
		InputExtensions().
		Return([]string{EXT_SOURCE})

	tchain.EXPECT().
		OutputExtension().
		Return(EXT_RUN)

	tempdir, err := ioutil.TempDir("", "test_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempdir)

	// when filename has runnable extension
	runnable := filepath.Join(tempdir, BASENAME+EXT_RUN)

	// should return error: Not Found
	sut := NewToolchainFinderFromToolchains([]domain.Toolchain{
		tchain,
	})

	_, err = sut.Produce(runnable)
	assert.Same(t, toolchain.ErrToolchainNotFound, err)
}
