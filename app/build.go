package app

import (
	"fmt"

	"github.com/mniak/bench/lib/bench"
)

func Build(folder string) error {
	builtPath, err := bench.Build(folder)
	if err == nil {
		fmt.Println(builtPath)
	}
	return err
}
