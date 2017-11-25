// [_Variadic functions_](http://en.wikipedia.org/wiki/Variadic_function)
// can be called with any number of trailing arguments.
// For example, `fmt.Println` is a common variadic
// function.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("strings = ", []string{"A", "B", "C"})
}
