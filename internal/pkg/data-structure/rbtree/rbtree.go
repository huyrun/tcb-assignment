package rbtree

import "fmt"

type Rbtree struct {
	root        *node
	len         int
	isVisualize bool
}

type Option func(*Rbtree)

func NewRbtree(opts ...Option) *Rbtree {
	r := &Rbtree{}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func Visualize() Option {
	return func(rbtree *Rbtree) {
		rbtree.isVisualize = true
	}
}

func (r *Rbtree) visualize() {
	if r.root == nil {
		return
	}
	str := "RedBlackTree\n"
	output(r.root, "", true, &str)
	fmt.Println(str)
}

func (r *Rbtree) Len() int {
	return r.len
}

func (r *Rbtree) Rank(x int) int {
	node := rank(r.root, x)
	if node == nil {
		return -1
	}

	return node.value
}

func (r *Rbtree) AddMany(values []int) {
	for _, v := range values {
		r.addOne(v)
	}

	if r.isVisualize {
		go r.visualize()
	}
}

func (r *Rbtree) addOne(value int) {
	r.root = insert(r.root, value)
	r.len++
}

func (r *Rbtree) Inorder() []int {
	return inorder(r.root)
}
