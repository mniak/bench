package impl

import (
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/bench/old_stuff/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_BuilderWithSkipWhenNotExist_WhenProgramExists_ShouldCallBuild(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	inputfile, err := os.CreateTemp("", "test_*.src")
	require.NoError(t, err)
	defer func() {
		inputfile.Close()
		os.Remove(inputfile.Name())
	}()

	inputpath := inputfile.Name()
	outputpath := gofakeit.Sentence(5)

	innerBuilder := mocks.NewMockBuilder(ctrl)
	innerBuilder.EXPECT().
		Build(inputpath).
		Return(outputpath, nil)

	sut := WrapBuilderWithSkipWhenNotExist(innerBuilder)

	result, err := sut.Build(inputpath)
	require.NoError(t, err)
	assert.Equal(t, outputpath, result)
}

func Test_BuilderWithSkipWhenNotExist_WhenProgramDoesNotExists_ShouldNotCallBuild(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	inputpath := gofakeit.Sentence(5)

	innerBuilder := mocks.NewMockBuilder(ctrl)
	sut := WrapBuilderWithSkipWhenNotExist(innerBuilder)

	result, err := sut.Build(inputpath)
	require.NoError(t, err)
	assert.Equal(t, inputpath, result)
}
