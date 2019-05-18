package main

import (
	"fmt"
)

func main() {

	// arrays
	var arr []string
	arr1 := []string{"item1", "item2"}

	// var arr := []string{"Club", "Heart", "Diamond", "Spade"}
	// maps
	var map1 map[string]string
	map2 := make(map[string]string)
	map3 := map[string]string{
		"item": "value",
	}

	fmt.Println(map1, map2, map3)
	fmt.Println(arr, arr1)
}
