package bench

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTest(t *testing.T) {
	sentence := gofakeit.Sentence(5)

	test := Test{
		Program:        "cat",
		Input:          sentence,
		ExpectedOutput: sentence,
	}

	started, err := StartTest(test)
	require.NoError(t, err, "start")

	result, err := WaitTest(started)
	require.NoError(t, err, "wait")

	assert.Equal(t, sentence, result.Output)
}
