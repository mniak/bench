package main

//go:generate mockgen -package=bench -destination=lib/bench/mock_bench_test.go     github.com/mniak/bench/lib/bench     FileFinder,ToolchainProducer,Builder,Tester
//go:generate mockgen -package=mocks -destination=internal/mocks/toolchain.go github.com/mniak/bench/lib/toolchain Toolchain
