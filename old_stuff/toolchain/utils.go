package toolchain

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mniak/bench/old_stuff/domain"
)

func findBinaryPath(exe string, paths ...string) (string, error) {
	if paths == nil || len(paths) == 0 {
		paths = strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))
	}
	for _, p := range paths {
		abs, err := filepath.Abs(p)
		if err != nil {
			continue
		}
		filename := filepath.Join(abs, exe+domain.OSBinaryExtension)

		info, err := os.Stat(filename)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return "", err
		}
		if !info.IsDir() {
			return filename, nil
		}

	}
	return "", ErrToolchainNotFound
}
