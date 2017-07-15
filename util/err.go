package util

import (
	"fmt"
	"os"
)

// ExitIfError log and exit if err exists
func ExitIfError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
