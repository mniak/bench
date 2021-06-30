package impl

import (
	"testing"

	"github.com/mniak/bench/domain"
	"github.com/stretchr/testify/require"
)

// func TestShouldHaveExtensionsAndFilenames(t *testing.T) {
// 	programFinder := DefaultProgramFinder.(*_FinderOnDirByFilenameAndExtensions)

// 	assert.Contains(t, programFinder.filenames, "main")
// 	assert.Contains(t, programFinder.extensions, ".py")

// 	switch runtime.GOOS {
// 	case "windows":
// 		assert.Contains(t, programFinder.extensions, ".exe")
// 		assert.Contains(t, programFinder.extensions, ".bat")
// 		assert.Contains(t, programFinder.extensions, ".cmd")
// 		assert.Contains(t, programFinder.extensions, ".ps1")

// 		assert.NotContains(t, programFinder.extensions, "", "(none)")
// 		assert.NotContains(t, programFinder.extensions, ".sh")
// 	default:
// 		assert.NotContains(t, programFinder.extensions, ".exe")
// 		assert.NotContains(t, programFinder.extensions, ".bat")
// 		assert.NotContains(t, programFinder.extensions, ".cmd")
// 		assert.NotContains(t, programFinder.extensions, ".ps1")

// 		assert.Contains(t, programFinder.extensions, "", "(none)")
// 		assert.Contains(t, programFinder.extensions, ".sh")
// 	}
// }

func Test_DefaultBuilder_Composition(t *testing.T) {
	var builder domain.Builder = DefaultBuilder

	require.IsType(t, &_BuilderWithFileFinder{}, builder)
	builder = builder.(*_BuilderWithFileFinder).Builder

	require.IsType(t, &_BaseBuilder{}, builder)
	builder = builder.(*_BaseBuilder)
}
