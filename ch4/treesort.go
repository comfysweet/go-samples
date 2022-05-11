package main

type tree struct {
	value       int
	left, right *tree
}

// Sort сортирует значения "на места"
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// добавляет элементы t к values в требуемом
// порядке и возвращает результирующий срез
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, v int) *tree {
	if t == nil {
		// эквивалентно возврату &tree{value:value}
		t := new(tree)
		t.value = v
		return t
	}
	if v < t.value {
		t.left = add(t.left, v)
	} else {
		t.right = add(t.right, v)
	}
	return t
}
