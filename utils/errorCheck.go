package utils

import "fmt"

// Check prints the desc of the error throughout
func Check(err error, desc string) {
	if err == nil {
		return
	}
	errStr := fmt.Sprintf("%s: %s", desc, err)
	panic(errStr)
}
