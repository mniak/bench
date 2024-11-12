package toolchain

var cppToolchainFactories = make([]ToolchainFactory, 0)

func NewCPPLoader() ToolchainLoader {
	return NewLoaderFromFactories(
		cppToolchainFactories,
		[]string{".cpp", ".cxx", ".c++"},
		OSBinaryExtension,
	)
}
