package bench

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Example struct {
	Name           string
	Input          string
	ExpectedOutput string
}

func FindExamples(directory string) ([]Example, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	examples := make(map[string]*Example)
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		filename := f.Name()
		bytes, err := os.ReadFile(path.Join(directory, filename))
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
