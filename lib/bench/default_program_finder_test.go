package bench

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldHaveExtensionsAndFilenames(t *testing.T) {
	assert.Contains(t, defaultProgramFinder.filenames, "main")

	assert.Contains(t, defaultProgramFinder.extensions, ".py")

	switch runtime.GOOS {
	case "windows":
		assert.Contains(t, defaultProgramFinder.extensions, ".exe")
		assert.Contains(t, defaultProgramFinder.extensions, ".bat")
		assert.Contains(t, defaultProgramFinder.extensions, ".cmd")
		assert.Contains(t, defaultProgramFinder.extensions, ".ps1")

		assert.NotContains(t, defaultProgramFinder.extensions, "", "(none)")
		assert.NotContains(t, defaultProgramFinder.extensions, ".sh")
	default:
		assert.NotContains(t, defaultProgramFinder.extensions, ".exe")
		assert.NotContains(t, defaultProgramFinder.extensions, ".bat")
		assert.NotContains(t, defaultProgramFinder.extensions, ".cmd")
		assert.NotContains(t, defaultProgramFinder.extensions, ".ps1")

		assert.Contains(t, defaultProgramFinder.extensions, "", "(none)")
		assert.Contains(t, defaultProgramFinder.extensions, ".sh")
	}
}

func TestUppercaseDefault_ShouldBeTheSameAsLowercaseDefault(t *testing.T) {
	assert.Same(t, defaultProgramFinder, defaultProgramFinder)
}
