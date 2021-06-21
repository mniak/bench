package toolchain

var cppToolchainFactories = make([]ToolchainFactory, 0)

func NewCPP() (Toolchain, error) {
	for _, factory := range cppToolchainFactories {
		tc, err := factory()
		if err == nil {
			return tc, nil
		}
		if err == ErrToolchainNotFound {
			continue
		}
	}
	return nil, ErrToolchainNotFound
}
