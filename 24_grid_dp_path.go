package main

import "fmt"

func gridDpPath(in [][]int) int {
	xL := len(in)
	yL := len(in[0])

	for j := 0; j < yL; j++ {
		for i := 0; i < xL; i++ {
			if in[i][j] == -1 {
				continue
			}
			if j == 0 {
				in[i][j] = 1
				continue
			}
			if i == 0 {
				in[i][j] = in[i][j-1]
				continue
			}
			if in[i][j-1] == -1 && in[i-1][j] == -1 {
				in[i][j] = 0
				continue
			} else if in[i-1][j] == -1 {
				in[i][j] = in[i][j-1]
				continue
			} else if in[i][j-1] == -1 {
				in[i][j] = in[i-1][j]
				continue
			}
			in[i][j] = in[i-1][j] + in[i][j-1]
		}
	}
	return in[xL-1][yL-1]
}

func main() {
	in := [][]int{{0, 0, 0, 0}, {0, 0, -1, 0}, {0, 0, 0, 0}, {0, -1, 0, 0}, {0, 0, 0, 0}}
	fmt.Println(gridDpPath(in))
}
