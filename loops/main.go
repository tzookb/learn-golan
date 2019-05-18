package main

import (
	"fmt"
)

func main() {

	// loop over a list/slice
	strs := []string{"str1", "str2", "str3"}
	for idx, str := range strs {
		fmt.Println(idx, str)
	}

	// loop n times
	n := 5
	for i := 0; i < n; i++ {
		fmt.Println(n, "times loop", i)
	}

	// endless loop
	for {
		fmt.Println("this could be endless")
		break
	}
}
