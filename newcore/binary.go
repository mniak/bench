package newcore

import "sync"

var (
	binaryRunnerOnce        sync.Once
	binaryRunnerInitialized Runner
)

func BinaryRunner() Runner {
	binaryRunnerOnce.Do(func() {
		binaryRunnerInitialized = loadPlatformSpecificBinaryRunner()
	})
	return binaryRunnerInitialized
}
