package main

import (
	"fmt"
	"sort"
)

// to list he key/value pairs in order, we need to first sort keys and then use them to index map

func sort_map(data map[string]int) {
	var keys []string
	for key := range data {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s \t %d\n", key, data[key])
	}
}

func concatenate_string_slice(list []string) string {
	return fmt.Sprintf("%q", list)
}

func main() {
	data := map[string]int{
		"b": 5,
		"e": 6,
		"a": 7,
	}
	sort_map(data)
	var nil_map map[string]int
	non_nil_map := make(map[string]int)
	fmt.Println(nil_map == nil)
	fmt.Println(non_nil_map == nil)

	strings := [3]string{"str1", "str2", "str3"}
	result := concatenate_string_slice(strings[:])
	fmt.Printf("list of strings is %v\n", result)
}
