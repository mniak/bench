package bench

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mniak/bench/old_stuff/internal/utils"
)

type Example struct {
	Name           string
	Input          string
	ExpectedOutput string
}

func FindExamples(dirOrProgram string, examplesDir string) ([]Example, error) {
	dir, _, err := utils.SplitDirAndProgram(dirOrProgram)
	if err != nil {
		return nil, err
	}

	fullExamplesDir := filepath.Join(dir, examplesDir)
	files, err := os.ReadDir(fullExamplesDir)
	if err != nil {
		return nil, err
	}

	examples := make(map[string]*Example)
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		filename := f.Name()
		bytes, err := os.ReadFile(path.Join(fullExamplesDir, filename))
		if err != nil {
			return nil, err
		}
		if strings.HasSuffix(filename, ".input.txt") {
			name := strings.TrimSuffix(filename, ".input.txt")
			ex, ok := examples[name]
			if !ok {
				ex = &Example{
					Name: name,
				}
				examples[name] = ex
			}
			ex.Input = string(bytes)
		} else if strings.HasSuffix(filename, ".output.txt") {
			name := strings.TrimSuffix(filename, ".output.txt")
			ex, ok := examples[name]
			if !ok {
				ex = &Example{
					Name: name,
				}
				examples[name] = ex
			}
			ex.ExpectedOutput = string(bytes)
		}
	}
	examplesSlice := make([]Example, 0)
	for _, ex := range examples {
		examplesSlice = append(examplesSlice, *ex)
	}
	return examplesSlice, nil
}
