package impl

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/mniak/bench/old_stuff/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTest(t *testing.T) {
	sentence := gofakeit.Sentence(5)

	test := domain.Test{
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

func cloneTest(test domain.Test) domain.Test {
	return test
}
