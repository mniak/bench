package impl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
