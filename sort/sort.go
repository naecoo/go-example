package main

import (
	"fmt"
	"slices"
)

func main() {

	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)
	fmt.Println("Str Sorted: ", slices.IsSorted(strs))

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)
	fmt.Println("Ints Sorted: ", slices.IsSorted(ints))
}
