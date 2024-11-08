package impl

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/bench/old_stuff/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuilderWithFileFinder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakepath := gofakeit.Sentence(5)
	fakefullpath := gofakeit.Sentence(5)
	fakebuilt := gofakeit.Sentence(5)

	filefinder := mocks.NewMockFileFinder(ctrl)
	filefinder.EXPECT().
		Find(fakepath).
		Return(fakefullpath, nil)

	builder := mocks.NewMockBuilder(ctrl)
	builder.EXPECT().
		Build(fakefullpath).
		Return(fakebuilt, nil)

	sut := WrapBuilderWithFileFinder(builder, filefinder)

	result, err := sut.Build(fakepath)
	require.NoError(t, err)
	assert.Equal(t, fakebuilt, result)
}
