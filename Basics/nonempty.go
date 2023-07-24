package main

import "fmt"

// remove the empty-string from the slice and return the modified slice
func removeempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}

// remove the i-th element from a slice
func remove_element(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	var strings = []string{"1", "3", "4", "", "5", "6", ""}
	fmt.Printf("%q\n", strings)
	strings = removeempty(strings)
	fmt.Printf("%q\n", strings)
	s := []int{5, 6, 7, 8, 9}
	fmt.Printf("Before removing %v\n", s)
	result := remove_element(s, 2)
	fmt.Printf("after removing %v\n", s)
	fmt.Printf("result is %v\n", result)
}
