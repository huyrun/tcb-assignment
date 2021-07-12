package rbtree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRbtree_AddOne(t *testing.T) {
	rbtree := NewRbtree()
	rbtree.addOne(1)
	rbtree.addOne(4)
	rbtree.addOne(6)
	rbtree.addOne(3)
	rbtree.addOne(5)
	rbtree.addOne(7)
	rbtree.addOne(8)
	rbtree.addOne(2)

	tests := []struct {
		name  string
		value int
	}{
		// TODO: Add test cases.
		{
			name:  "test",
			value: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rbtree.addOne(tt.value)
			fmt.Println(rbtree.root.value)
			fmt.Println(rbtree.Inorder())
		})
	}
}

func TestRbtree_AddMany(t *testing.T) {
	rbtree := NewRbtree()

	tests := []struct {
		name   string
		values []int
	}{
		// TODO: Add test cases.
		{
			name:   "test",
			values: []int{1, 4, 6, 3, 5, 7, 8, 2, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rbtree.AddMany(tt.values)
			str := "RedBlackTree\n"
			output(rbtree.root, "", true, &str)
			fmt.Println(str)
		})
	}
}

func Test_rank(t *testing.T) {
	rbt := NewRbtree()

	type args struct {
		values []int
		x      int
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				values: []int{1, 9, 6, 0},
				x:      1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rbt.AddMany(tt.args.values)
			rbt.visualize()
			got := rank(rbt.root, tt.args.x)
			fmt.Println("got: ", got.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRbtree_Rank(t *testing.T) {
	rbt := NewRbtree()

	type args struct {
		values []int
		x      int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				values: []int{1, 9, 6, 0},
				x:      4,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rbt.AddMany(tt.args.values)
			rbt.Visualize()
			if got := rbt.Rank(tt.args.x); got != tt.want {
				t.Errorf("Rank() = %v, want %v", got, tt.want)
			}
		})
	}
}
