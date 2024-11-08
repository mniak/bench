package impl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_DefaultBuilder_Composition(t *testing.T) {
	require.IsType(t, &_BuilderWithFileFinder{}, DefaultBuilder)
	bff := DefaultBuilder.(*_BuilderWithFileFinder)
	require.Same(t, DefaultSourceFinder, bff.fileFinder)

	require.IsType(t, &_BaseBuilder{}, bff.Builder)
	bb := bff.Builder.(*_BaseBuilder)

	require.Same(t, DefaultToolchainFinder, bb.toolchainFinder)
}

func Test_DefaultTester_Composition(t *testing.T) {
	require.IsType(t, &_TesterWithFileFinder{}, DefaultTester)

	tester_with_filefinder := DefaultTester.(*_TesterWithFileFinder)
	require.IsType(t, &_TesterWithBuilder{}, tester_with_filefinder.Tester)
	require.Same(t, DefaultSourceFinder, tester_with_filefinder.fileFinder)

	tester_with_builder := tester_with_filefinder.Tester.(*_TesterWithBuilder)
	require.IsType(t, &_BaseTester{}, tester_with_builder.Tester)
	require.IsType(t, &_BuilderWithSkipWhenNotExist{}, tester_with_builder.builder)

	builder_with_skip := tester_with_builder.builder.(*_BuilderWithSkipWhenNotExist)
	require.IsType(t, &_BaseBuilder{}, builder_with_skip.Builder)

	base_builder := builder_with_skip.Builder.(*_BaseBuilder)
	require.Same(t, DefaultToolchainFinder, base_builder.toolchainFinder)
}
