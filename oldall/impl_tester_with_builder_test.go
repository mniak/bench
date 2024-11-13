package oldall

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/bench/old_stuff/internal/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_TesterWithBuilder_Start_ShouldCallBuilder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var fakeTest Test
	gofakeit.Struct(&fakeTest)
	var fakeStarted StartedTest
	gofakeit.Struct(&fakeStarted)

	fakeProgram := gofakeit.Sentence(5)
	fakeTestWithFakeProgram := cloneTest(fakeTest)
	fakeTestWithFakeProgram.Program = fakeProgram

	builder := mocks.NewMockBuilder(ctrl)
	builder.EXPECT().Build(fakeTest.Program).Return(fakeProgram, nil)

	tester := mocks.NewMockTester(ctrl)
	tester.EXPECT().Start(fakeTestWithFakeProgram).Return(fakeStarted, nil)

	sut := WrapTesterWithBuilder(tester, builder)
	result, err := sut.Start(fakeTest)
	require.NoError(t, err)
	assert.Equal(t, fakeStarted, result)
}

func Test_TesterWithBuilder_Wait_ShouldBypass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var fakeStarted StartedTest
	gofakeit.Struct(&fakeStarted)
	var fakeResult TestResult
	gofakeit.Struct(&fakeResult)
	fakeError := errors.New(gofakeit.Sentence(5))

	tester := mocks.NewMockTester(ctrl)
	builder := mocks.NewMockBuilder(ctrl)

	tester.EXPECT().Wait(fakeStarted).Return(fakeResult, fakeError)

	sut := WrapTesterWithBuilder(tester, builder)
	result, err := sut.Wait(fakeStarted)
	require.Equal(t, fakeError, err)
	assert.Equal(t, fakeResult, result)
}
