package runner

import "github.com/mniak/bench/domain"

var loaders = []domain.RunnerLoader{
	NewBinaryLoader(),
	NewPythonLoader(),
}

func RunnerFor(filename string) (domain.Runner, error) {
	for _, loader := range loaders {
		can := loader.CanRun(filename)
		if !can {
			continue
		}
		r, err := loader.Load()
		if err == nil {
			return r, nil
		}
		if err == ErrRunnerNotFound {
			continue
		}
	}
	return nil, ErrRunnerNotFound
}
