package tests

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/bench/internal/mocks"
	"github.com/mniak/bench/lib/bench"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuilder_Build(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakepath := gofakeit.Sentence(5)
	fakebuilt := gofakeit.Sentence(5)

	tchain := mocks.NewMockToolchain(ctrl)
	tchain.EXPECT().
		Build(fakepath).
		Return(fakebuilt, nil)

	toolchainProducer := mocks.NewMockToolchainProducer(ctrl)
	toolchainProducer.EXPECT().
		Produce(fakepath).
		Return(tchain, nil)

	sut := bench.NewBuilder(toolchainProducer)

	builtPath, err := sut.Build(fakepath)
	require.NoError(t, err, "build")
	assert.Equal(t, fakebuilt, builtPath, "built path")
}

func TestBuild(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakepath := gofakeit.Sentence(5)
	fakebuilt := gofakeit.Sentence(5)

	builder := mocks.NewMockBuilder(ctrl)
	builder.EXPECT().
		Build(fakepath).
		Return(fakebuilt, nil)

	bench.DefaultBuilder = builder
	result, err := bench.Build(fakepath)

	assert.NoError(t, err)
	assert.Equal(t, fakebuilt, result)
}

// func TestBuilderWithProgramFinder_Build(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	fakepath := gofakeit.Sentence(5)
// 	fakefullpath := gofakeit.Sentence(5)
// 	fakebuilt := gofakeit.Sentence(5)

// 	programFinder := mocks.NewMockFileFinder(ctrl)
// 	programFinder.EXPECT().
// 		Find(fakepath).
// 		Return(fakefullpath, nil)

// 	innerBuilder := mocks.NewMockBuilder(ctrl)
// 	innerBuilder.EXPECT().
// 		Build(fakefullpath).
// 		Return(fakebuilt, nil)

// 	sut := WrapBuilderWithProgramFinder(
// 		innerBuilder,
// 		programFinder,
// 	)

// 	builtPath, err := sut.Build(fakepath)
// 	require.NoError(t, err, "build")
// 	assert.Equal(t, fakebuilt, builtPath, "built path")
// }
