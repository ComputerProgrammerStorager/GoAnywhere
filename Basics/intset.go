package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers
// Its zero value represents the empty set
type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// set s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String is the defined method that returns the int array as a string in the form of "{1,2,3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var data1 IntSet
	data1.Add(1)
	data1.Add(2)
	fmt.Println(data1.String())
	var data2 IntSet
	data2.Add(3)
	data2.Add(4)
	fmt.Println(data2.String())
	data1.UnionWith(&data2)
	fmt.Println(data1.String())
	fmt.Println(&data1)
	fmt.Println(data1)

}
