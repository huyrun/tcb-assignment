package rbtree

func insertFixup(root *node, node *node) *node {
	parent := node.parent
	if parent == nil {
		node.color = black
		return root
	}

	if parent.color == red {
		grandparent := parent.parent
		if parent == grandparent.left {
			uncle := grandparent.right
			if uncle != nil && uncle.color == red {
				parent.color = black
				uncle.color = black
				grandparent.color = red
				return insertFixup(root, grandparent)
			}
		} else { // parent == grandparent.right
			uncle := grandparent.left
			if uncle != nil && uncle.color == red {
				parent.color = black
				uncle.color = black
				grandparent.color = red
				return insertFixup(root, grandparent)
			}
		}

		if parent == grandparent.left && node == parent.left { // left left case
			root = rightRotate(root, grandparent)
			grandparent.color, parent.color = parent.color, grandparent.color
		} else if parent == grandparent.left && node == parent.right { // left right case
			root = leftRotate(root, parent)
			root = rightRotate(root, grandparent)
			grandparent.color, node.color = node.color, grandparent.color
		} else if parent == grandparent.right && node == parent.right { // right right case
			root = leftRotate(root, grandparent)
			grandparent.color, parent.color = parent.color, grandparent.color
		} else if parent == grandparent.right && node == parent.left { // right left case
			root = rightRotate(root, parent)
			root = leftRotate(root, grandparent)
			grandparent.color, node.color = node.color, grandparent.color
		}
	}

	root.color = black
	return root
}

func leftRotate(root *node, node *node) *node {
	x := node.right
	y := x.left

	x.left = node
	node.right = y
	x.parent = node.parent
	if x.parent == nil {
		root = x
	} else if x.parent.left == node {
		x.parent.left = x
	} else {
		x.parent.right = x
	}
	node.parent = x
	if y != nil {
		y.parent = node
	}

	node.size = 1
	if node.left != nil {
		node.size += node.left.size
	}
	if node.right != nil {
		node.size += node.right.size
	}

	x.size = 1
	if x.left != nil {
		x.size += x.left.size
	}
	if x.right != nil {
		x.size += x.right.size
	}

	return root
}

func rightRotate(root *node, node *node) *node {
	x := node.left
	y := x.right

	x.right = node
	node.left = y
	x.parent = node.parent
	if x.parent == nil {
		root = x
	} else if x.parent.left == node {
		x.parent.left = x
	} else {
		x.parent.right = x
	}
	node.parent = x
	if y != nil {
		y.parent = node
	}

	node.size = 1
	if node.left != nil {
		node.size += node.left.size
	}
	if node.right != nil {
		node.size += node.right.size
	}

	x.size = 1
	if x.left != nil {
		x.size += x.left.size
	}
	if x.right != nil {
		x.size += x.right.size
	}

	return root
}
