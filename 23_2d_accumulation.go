package main

import (
	"fmt"
)

type xy struct {
	X int
	Y int
}

type rectangle struct {
	leftT  xy
	rightB xy
}

func accumulation2d(in [][]int, set []rectangle) [][]int {
	for _, v := range set {
		x1 := v.leftT.X
		y1 := v.leftT.Y
		x2 := v.rightB.X
		y2 := v.rightB.Y
		in[x1][y1] = in[x1][y1] + 1
		in[x2][y2] = in[x2][y2] + 1
		in[x1][y2] = in[x1][y2] - 1
		in[x2][y1] = in[x2][y1] - 1
	}

	lx := len(in)
	ly := len(in[0])

	for i := 1; i <= (lx - 1); i++ {
		for j := 0; j <= (ly - 1); j++ {
			in[i][j] = in[i][j] + in[i-1][j]
		}
	}
	for j := 1; j <= (ly - 1); j++ {
		for i := 0; i <= (lx - 1); i++ {
			in[i][j] = in[i][j] + in[i][j-1]
		}
	}
	return in

}

func main() {
	rec1 := rectangle{xy{1, 1}, xy{5, 4}}
	rec2 := rectangle{xy{3, 2}, xy{7, 6}}
	rec3 := rectangle{xy{3, 3}, xy{6, 6}}
	set := []rectangle{rec1, rec2, rec3}
	list := make([][]int, 8)
	for i := 0; i < len(list); i++ {
		list[i] = make([]int, 7)
	}
	fmt.Println(accumulation2d(list, set))
}
