package main

type tree struct {
	value       int
	left, right *tree // define pointers to tree
}

func Sort(values []int) {
	var root *tree
	for _, val := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appends the elements of t to values in order and return the resulting slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = appendValues(values, t.value)
		values = appendValues(values, t.right)
	}

	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}
