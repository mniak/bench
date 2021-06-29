package main

//go:generate mockgen -package=mocks -destination=internal/mocks/builder.go				github.com/mniak/bench/domain Builder
//go:generate mockgen -package=mocks -destination=internal/mocks/filefinder.go			github.com/mniak/bench/domain FileFinder
//go:generate mockgen -package=mocks -destination=internal/mocks/mocks.go				github.com/mniak/bench/domain Tester
//go:generate mockgen -package=mocks -destination=internal/mocks/toolchain_producer.go	github.com/mniak/bench/domain ToolchainProducer
//go:generate mockgen -package=mocks -destination=internal/mocks/toolchain.go			github.com/mniak/bench/domain Toolchain
