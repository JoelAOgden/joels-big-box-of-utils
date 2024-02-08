package error_util

import "fmt"

// Throws a panic with error e if NOT nil
func PanicCheck(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}
