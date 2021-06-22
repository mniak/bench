package internal

//go:generate mockgen -destination=mock_bench/build.go github.com/mniak/bench ProgramFinder,ToolchainProducer
//go:generate mockgen -destination=mock_toolchain/build.go github.com/mniak/bench/toolchain Toolchain
