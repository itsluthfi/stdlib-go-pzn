package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Paul", "George", "Adam", "John"}
	values := []int{54, 34, 756, 114}

	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Index(names, "Adam"))
	fmt.Println(slices.Index(names, "Alan")) // -1 berarti gaada
	fmt.Println(slices.Index(values, 43))    // -1 berarti gaada
	fmt.Println(slices.Index(values, 34))
	fmt.Println(slices.Contains(names, "Adam"))
}
