package main

import "fmt"

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func New() *Tree {
	tree := &Tree{}
	tree.Value = 0
	return tree
}

func (t *Tree) newLeft() {
	t.Left = New()
}

func (t *Tree) newRight() {
	t.Right = New()
}
func (t *Tree) newLR() {
	t.Left = New()
	t.Right = New()
}

func makeBtree() *Tree {
	t := New()
	t.newLR()
	t.Left.newLR()
	t.Right.newLR()
	t.Left.Right.newLR()
	t.Right.Right.newRight()
	return t
}

var time = 1

func (t *Tree) preorder() {
	if t == nil {
		return
	}
	t.Value = time
	time++
	t.Left.preorder()
	t.Right.preorder()
}

func main() {
	t := makeBtree()
	t.preorder()
	fmt.Println(t.Right)
}
