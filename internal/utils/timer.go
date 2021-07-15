package utils

import (
	"fmt"
	"time"
)

type Finisher func()

func StartTimer(context string) Finisher {
	t := time.Now()
	fmt.Printf("STARTED - %s\n", context)
	return func() {
		fmt.Printf("FINISHED - %s - %s\n", context, time.Now().Sub(t))
	}
}
