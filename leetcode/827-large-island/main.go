package main

import "fmt"

const N = 500

type sol struct {
	cols    int
	rows    int
	grid    [][]int
	visited [N][N]bool
	ids     [N][N]int
}

var (
	cc = [4]int{-1, 0, 0, 1} // col directions
	rr = [4]int{0, 1, -1, 0} // row directions
)

func (s *sol) solve() int {
	color := 0
	for r := 0; r < s.rows; r++ {
		for c := 0; c < s.cols; c++ {
			s.fill(c, r, color)
			color++
		}
	}

	res := -1
	mp := make(map[int]int)
	for _, vv := range s.ids[:s.rows] {
		for _, v := range vv[:s.cols] {
			mp[v]++
			if v != -1 {
				res = max(res, mp[v])
			}
		}
	}

	for r := 0; r < s.rows; r++ {
		for c := 0; c < s.cols; c++ {
			if s.ids[r][c] == -1 {
				// look at 4 sides
				sum := 0
				m := map[int]int{}
				for i := range cc {
					newC := c + cc[i]
					newR := r + rr[i]
					if s.inGrid(newC, newR) && s.ids[newR][newC] != -1 {
						m[s.ids[newR][newC]] = mp[s.ids[newR][newC]]
					}
				}
				for _, v := range m {
					sum += v
				}
				res = max(res, sum+1)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (s *sol) inGrid(c, r int) bool {
	return (c >= 0 && c < s.cols) && (r >= 0 && r < s.cols)
}

func (s *sol) fill(c, r, color int) {
	if s.visited[r][c] {
		return
	}
	s.visited[r][c] = true

	// edge
	if s.grid[r][c] == 0 {
		s.ids[r][c] = -1
		return
	}

	s.ids[r][c] = color

	for i := range cc {
		newC := c + cc[i]
		newR := r + rr[i]
		if s.inGrid(newC, newR) {
			s.fill(newC, newR, color)
		}
	}
}

func largestIsland(grid [][]int) int {
	s := sol{
		grid: grid,
		rows: len(grid),
		cols: len(grid[0]),
	}
	return s.solve()
}

func main() {
	fmt.Println(largestIsland([][]int{{0, 0}, {0, 0}}))
}
