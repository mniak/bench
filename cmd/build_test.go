package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mniak/bench/impl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Build_WithFolderName(t *testing.T) {
	const SOURCE_CODE = `
#include <strings>

using namespace std;

void main() {
    string name;6
    name << cin;
    cout << "BUILD: ";
    cout << "Hello, " << name << endl;
}`

	tempdir, err := os.MkdirTemp("", "build_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempdir)

	err = os.WriteFile(filepath.Join(tempdir, "main.cpp"), []byte(SOURCE_CODE), 0644)
	require.NoError(t, err)

	err = Build(tempdir)
	require.NoError(t, err)

	assert.FileExists(t, filepath.Join(tempdir, "main"+impl.DEFAULT_BINARY_OS_EXTENSION))
}
