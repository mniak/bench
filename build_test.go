package bench

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/bench/internal/mock_bench"
	"github.com/mniak/bench/internal/mock_toolchain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuilderBuild(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakepath := gofakeit.Sentence(5)
	fakesource := gofakeit.Sentence(5)
	fakebuilt := gofakeit.Sentence(5)

	programFinder := mock_bench.NewMockProgramFinder(ctrl)
	programFinder.EXPECT().
		Find(fakepath).
		Return(fakesource, nil)

	tchain := mock_toolchain.NewMockToolchain(ctrl)
	tchain.EXPECT().
		Build(fakesource).
		Return(fakebuilt, nil)

	toolchainProducer := mock_bench.NewMockToolchainProducer(ctrl)
	toolchainProducer.EXPECT().
		Produce(fakesource).
		Return(tchain, nil)

	builder := Builder{
		programFinder:     programFinder,
		toolchainProducer: toolchainProducer,
	}

	builtPath, err := builder.Build(fakepath)
	require.NoError(t, err, "build")
	assert.Equal(t, fakebuilt, builtPath, "built path")
}
