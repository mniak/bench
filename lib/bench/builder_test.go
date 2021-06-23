package bench

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/bench/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuilderBuild(t *testing.T) {
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

	builder := _Builder{
		toolchainProducer: toolchainProducer,
	}

	builtPath, err := builder.Build(fakepath)
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

	DefaultBuilder = builder
	result, err := Build(fakepath)

	assert.NoError(t, err)
	assert.Equal(t, fakebuilt, result)
}
