package rbtree

type color int

const (
	red color = iota
	black
)

type node struct {
	value  int
	left   *node
	right  *node
	parent *node
	color  color
	size   int
}

func insert(root *node, value int) *node {
	newNode := &node{
		value: value,
		size:  1,
		color: red,
	}

	root = doInsert(root, newNode)
	root = insertFixup(root, newNode)
	return root
}

func doInsert(root *node, newNode *node) *node {
	if root == nil {
		return newNode
	}

	if newNode.value <= root.value {
		root.left = doInsert(root.left, newNode)
		root.left.parent = root
		root.size++
	}

	if newNode.value > root.value {
		root.right = doInsert(root.right, newNode)
		root.right.parent = root
		root.size++
	}

	return root
}

func inorder(root *node) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)

	if root.left != nil {
		res = inorder(root.left)
	}

	res = append(res, root.value)

	if root.right != nil {
		res = append(res, inorder(root.right)...)
	}

	return res
}

func rank(root *node, x int) *node {
	if root == nil {
		return nil
	}

	var r int
	var leftSize int
	if root.left != nil {
		leftSize = root.left.size
	}

	r = leftSize + 1
	if r == x {
		return root
	}

	if x < r {
		return rank(root.left, x)
	}

	return rank(root.right, x-r)

	return nil
}
