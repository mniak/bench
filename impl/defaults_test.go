package impl

import (
	"testing"

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
	require.IsType(t, &_BuilderWithFileFinder{}, DefaultBuilder)
	bff := DefaultBuilder.(*_BuilderWithFileFinder)
	require.Same(t, DefaultProgramFinder, bff.fileFinder)

	require.IsType(t, &_BaseBuilder{}, bff.Builder)
	bb := bff.Builder.(*_BaseBuilder)

	require.Same(t, DefaultToolchainFinder, bb.toolchainFinder)
}

func Test_DefaultTester_Composition(t *testing.T) {
	// DefaultTester =
	//   decorator WithFullPath over
	require.IsType(t, &_TesterWithFileFinder{}, DefaultTester)
	tw_finder := DefaultTester.(*_TesterWithFileFinder)

	//   decorator WithBuilder(TestBuilder) over
	require.IsType(t, &_TesterWithBuilder{}, tw_finder.Tester)
	tw_builder := tw_finder.Tester.(*_TesterWithBuilder)

	//             BaseTester
	require.IsType(t, &_BaseTester{}, tw_builder.Tester)

	// TestBuilder =
	//   decorator SkipIfDoesNotExist over
	require.IsType(t, &_BuilderWithSkipWhenNotExist{}, tw_builder.builder)
	bw_skip := tw_builder.builder.(*_BuilderWithSkipWhenNotExist)

	//             BaseBuilder
	require.IsType(t, &_BaseBuilder{}, bw_skip.Builder)
	base_builder := bw_skip.Builder.(*_BaseBuilder)

	require.Equal(t, DefaultToolchainFinder, base_builder.toolchainFinder)
}
