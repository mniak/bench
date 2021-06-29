package domain

type ToolchainProducer interface {
	Produce(string) (Toolchain, error)
}
