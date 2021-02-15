package ch04

type node struct {
	data        int
	left, right *node
}

func Sort(values []int) {
	var tree *node
	// insert data to the binary tree.
	for _, v := range values {
		tree = add(tree, v)
	}
	appendValue(values[:0], tree)
}

func appendValue(data []int, n *node) []int {
	if n != nil {
		data = appendValue(data, n.left)
		data = append(data, n.data)
		data = appendValue(data, n.right)
	}
	return data
}

func add(n *node, data int) *node {
	if n == nil {
		n = new(node)
		n.data = data
		return n
	}
	if data < n.data {
		n.left = add(n.left, data)
	} else {
		n.right = add(n.right, data)
	}
	return n
}
