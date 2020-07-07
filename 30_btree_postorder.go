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

func (t *Tree) postorder() {
	if t == nil {
		return
	}
	t.Left.postorder()
	t.Right.postorder()
	t.Value = time
	time++
}

func main() {
	t := makeBtree()
	t.postorder()
	fmt.Println(t.Left.Left)
}
