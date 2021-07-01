package bench

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/bench/internal/mocks"
	"github.com/stretchr/testify/assert"
)

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
