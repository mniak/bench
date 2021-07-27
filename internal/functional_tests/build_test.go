package functional_tests

import (
	"os"
	"testing"

	"github.com/mniak/bench/cmd"
	"github.com/stretchr/testify/require"
)

func Test_Build_WithFolderName(t *testing.T) {
	os.Args = []string{"cmd"}

	const FOLDER = "./build"

	err := cmd.Build(FOLDER)
	require.NoError(t, err)
}
