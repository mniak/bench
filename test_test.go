package bench

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTest(t *testing.T) {
	sentence := gofakeit.Sentence(5)

	test := NewTest("cat", sentence, sentence)

	err := test.Start()
	require.NoError(t, err, "start")

	result, err := test.Wait()
	require.NoError(t, err, "wait")

	assert.Equal(t, sentence, result.Output)
}
