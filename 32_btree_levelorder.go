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

// var time = 1

func (t *Tree) levelorder() {
	var que []*Tree
	que = append(que, t)
	var time = 1
	for {
		if len(que) == 0 {
			break
		}
		q := que[0]
		q.Value = time
		time++
		que = que[1:]
		if q.Left != nil {
			que = append(que, q.Left)
		}
		if q.Right != nil {
			que = append(que, q.Right)
		}
	}
}

func main() {
	t := makeBtree()
	t.levelorder()
	fmt.Println(t.Right.Right.Right)
}
