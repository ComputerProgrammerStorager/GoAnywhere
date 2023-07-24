package main

import "fmt"

// reverse a slice of ints
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// rotate a slice left by n elements can be achieved in the following three steps
// 1.reverse the first n elements
// 2. reverse the remaining elements
// 3. reverse the entire array
func left_rotate_2(s []int) {
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
}

// append function appends items to slices
func append_slice() {
	var runes []rune
	for _, r := range "Hello 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}

func main() {
	months := [...]string{1: "January", 2: "Feb", 3: "Mar", 4: "April", 5: "May", 6: "June", 7: "July", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Printf("len %d cap %d type %T \n", len(months), cap(months), months)
	fmt.Printf("Q2: %v len: %d cap: %d\n", Q2, len(Q2), cap(Q2))
	fmt.Printf("summer: %v, len: %d cap %d \n", summer, len(summer), cap(summer))

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("s: %v, len %d cap: %d\n", s, len(s), cap(s))
	left_rotate_2(s)
	fmt.Printf("s: %v, len %d cap: %d\n", s, len(s), cap(s))
	append_slice()
}
