package main

//go:generate mockgen -package=mocks -destination=internal/mocks/bench.go     github.com/mniak/bench/lib/bench     ProgramFinder,ToolchainProducer,Builder
//go:generate mockgen -package=mocks -destination=internal/mocks/toolchain.go github.com/mniak/bench/lib/toolchain Toolchain
