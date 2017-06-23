package util

import (
	"fmt"
	"os"
)

func ExitIfError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
