package impl

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/bench/internal/mocks"
	extra "github.com/oxyno-zeta/gomock-extra-matcher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_WhenFindToolchain_ShouldBuild(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeinputpath := gofakeit.Sentence(5)
	fakeoutputpath := getBinaryPath(fakeinputpath)

	tchain := mocks.NewMockToolchain(ctrl)
	tchainLoader := mocks.NewMockToolchainLoader(ctrl)
	toolchainFinder := mocks.NewMockToolchainFinder(ctrl)

	// when find toolchain
	toolchainFinder.EXPECT().
		Find(fakeinputpath).
		Return(tchainLoader, nil)

	tchainLoader.EXPECT().
		Load().
		Return(tchain, nil)

	tchain.EXPECT().
		Build(extra.StructMatcher().
			Field("Input", fakeinputpath).
			Field("Output", fakeoutputpath)).
		Return(nil)

	// should build
	sut := NewBuilder(toolchainFinder)
	builtPath, err := sut.Build(fakeinputpath)
	require.NoError(t, err)
	assert.Equal(t, fakeoutputpath, builtPath)
}

func Test_WhenFindToolchain_ButToolchainFailsToBuild_ShouldFailBuildWithTheSameError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeinputpath := gofakeit.Sentence(5)
	fakeoutputpath := getBinaryPath(fakeinputpath)

	failure := errors.New(gofakeit.Sentence(5))

	tchain := mocks.NewMockToolchain(ctrl)
	tchainLoader := mocks.NewMockToolchainLoader(ctrl)
	toolchainFinder := mocks.NewMockToolchainFinder(ctrl)

	// when find toolchain
	toolchainFinder.EXPECT().
		Find(fakeinputpath).
		Return(tchainLoader, nil)

	tchainLoader.EXPECT().
		Load().
		Return(tchain, nil)

	// but toolchain fails to build
	tchain.EXPECT().
		Build(extra.StructMatcher().
			Field("Input", fakeinputpath).
			Field("Output", fakeoutputpath)).
		Return(failure)

	// should fail build with the same error
	sut := NewBuilder(toolchainFinder)
	_, err := sut.Build(fakeinputpath)
	assert.Same(t, failure, err)
}

func Test_WhenDoesNotFindToolchain_ShouldFailBuildWithTheSameError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakepath := gofakeit.Sentence(5)
	failure := errors.New(gofakeit.Sentence(5))

	toolchainFinder := mocks.NewMockToolchainFinder(ctrl)

	// when find toolchain
	toolchainFinder.EXPECT().
		Find(fakepath).
		Return(nil, failure)

	// should fail build with the same error
	sut := NewBuilder(toolchainFinder)
	_, err := sut.Build(fakepath)
	assert.Same(t, failure, err)
}
