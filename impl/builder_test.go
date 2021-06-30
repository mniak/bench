package impl

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/bench/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_WhenFindToolchain_ShouldBuild(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakepath := gofakeit.Sentence(5)
	fakebuilt := gofakeit.Sentence(5)

	tchain := mocks.NewMockToolchain(ctrl)
	toolchainFinder := mocks.NewMockToolchainFinder(ctrl)

	// when find toolchain
	toolchainFinder.EXPECT().
		Find(fakepath).
		Return(tchain, nil)

	tchain.EXPECT().
		Build(fakepath).
		Return(fakebuilt, nil)

	// should build
	sut := NewBuilder(toolchainFinder)
	builtPath, err := sut.Build(fakepath)
	require.NoError(t, err)
	assert.Equal(t, fakebuilt, builtPath)
}

func Test_WhenFindToolchain_ButToolchainFailsToBuild_ShouldFailBuildWithTheSameError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakepath := gofakeit.Sentence(5)
	failure := errors.New(gofakeit.Sentence(5))

	tchain := mocks.NewMockToolchain(ctrl)

	toolchainFinder := mocks.NewMockToolchainFinder(ctrl)

	// when find toolchain
	toolchainFinder.EXPECT().
		Find(fakepath).
		Return(tchain, nil)

	// but toolchain fails to build
	tchain.EXPECT().
		Build(fakepath).
		Return("", failure)

	// should fail build with the same error
	sut := NewBuilder(toolchainFinder)
	_, err := sut.Build(fakepath)
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
