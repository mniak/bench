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

func TestTest(t *testing.T) {
	sentence := gofakeit.Sentence(5)

	test := bench.Test{
		Program:        "cat",
		Input:          sentence,
		ExpectedOutput: sentence,
	}

	sut := bench.NewTester()

	started, err := sut.Start(test)
	require.NoError(t, err, "start")

	result, err := sut.Wait(started)
	require.NoError(t, err, "wait")

	assert.Equal(t, sentence, result.Output)
}

func cloneTest(test bench.Test) bench.Test {
	return test
}

func TestWrapWithProgramFinder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var faketest bench.Test
	var fakestarted bench.StartedTest
	var fakeresult bench.TestResult

	require.NoError(t, gofakeit.Struct(&faketest))
	require.NoError(t, gofakeit.Struct(&fakestarted))
	require.NoError(t, gofakeit.Struct(&fakeresult))

	fakeprogram := gofakeit.Sentence(5)
	faketestWithFakeprogram := cloneTest(faketest)
	faketestWithFakeprogram.Program = fakeprogram

	tester := mocks.NewMockTester(ctrl)
	tester.EXPECT().
		Start(faketestWithFakeprogram).
		Return(fakestarted, nil)
	tester.EXPECT().
		Wait(fakestarted).
		Return(fakeresult, nil)

	programFinder := mocks.NewMockFileFinder(ctrl)
	programFinder.EXPECT().
		Find(faketest.Program).
		Return(fakeprogram, nil)

	sut := bench.WrapTesterWithProgramFinder(tester, programFinder)

	started, err := sut.Start(faketest)
	assert.NoError(t, err)

	result, err := sut.Wait(started)
	assert.NoError(t, err)

	assert.Equal(t, fakeresult, result)
}

func TestWrapWithBuilder(t *testing.T) {
}
