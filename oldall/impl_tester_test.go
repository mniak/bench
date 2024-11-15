package oldall

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTest(t *testing.T) {
	sentence := gofakeit.Sentence(5)

	test := oldall.Test{
		Program:        "cat",
		Input:          sentence,
		ExpectedOutput: sentence,
	}

	sut := NewTester()

	started, err := sut.Start(test)
	require.NoError(t, err, "start")

	result, err := sut.Wait(started)
	require.NoError(t, err, "wait")

	assert.Equal(t, sentence, result.Output)
}

func cloneTest(test Test) Test {
	return test
}