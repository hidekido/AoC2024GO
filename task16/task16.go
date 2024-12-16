// Package main -?
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("task16/task16.txt")
	if err != nil {
		log.Fatalf("failed to open file task16.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task16.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	field := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, []rune(line))
	}
	x, y := searchStart(field)
	res, res2 := searchPath(x, y, field)
	fmt.Println(res)
	fmt.Println(res2)
}

func searchStart(field [][]rune) (int, int) {
	for i := range field {
		for j := range field[0] {
			if field[i][j] == 'S' {
				return i, j
			}
		}
	}
	return -1, -1
}

var dirs = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

type Step struct {
	x, y  int
	score int
	dir   int
}

type Steps []Step

func (h Steps) Len() int {
	return len(h)
}

func (h Steps) Less(i, j int) bool {
	return h[i].score < h[j].score
}

func (h Steps) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Steps) Push(i any) {
	*h = append(*h, i.(Step))
}

func (h *Steps) Pop() any {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return res
}

func searchPath(x, y int, field [][]rune) (int, int) {
	steps := make(Steps, 0)
	steps.Push(Step{x, y, 0, 0})
	visited := make([][]int, len(field))
	for i := range visited {
		visited[i] = make([]int, len(field[0]))
		for j := range visited[i] {
			visited[i][j] = 1_000_000_000_000
		}
	}
	visited[x][y] = 0
	heap.Init(&steps)
	path, ex, ey := 0, 0, 0
	for steps.Len() > 0 {
		step := heap.Pop(&steps).(Step)

		if field[step.x][step.y] == 'E' {
			if path == 0 {
				path, ex, ey = step.score, step.x, step.y
			}
			continue
		}
		dir := step.dir
		for i := 0; i < 4; i++ {
			dir = (dir + i) % 4
			nextX, nextY := step.x+dirs[dir][0], step.y+dirs[dir][1]
			price := abs(step.dir-dir) % 2 * 1000
			if field[nextX][nextY] != '#' && visited[nextX][nextY] > step.score+1+price {
				heap.Push(&steps, Step{nextX, nextY, step.score + 1 + price, dir})
				visited[nextX][nextY] = step.score + 1 + price
			}
		}
	}
	spots := make([][]bool, len(field))
	for i := range spots {
		spots[i] = make([]bool, len(field[0]))
	}
	for dir := 0; dir < 4; dir++ {
		mark(ex, ey, dir, path, visited, field, spots)
	}

	res2 := 0
	for i := range spots {
		for j := range spots[i] {
			if spots[i][j] {
				res2++
			}
		}
	}

	return path, res2
}

func abs(x int) int {
	return max(x, -x)
}

func mark(x, y, dir, score int, visited [][]int, field [][]rune, spots [][]bool) bool {
	if field[x][y] == '#' || visited[x][y] > score {
		return false
	}
	if field[x][y] == 'S' {
		if score == 0 {
			spots[x][y] = true
			return true
		}
		return false
	}
	if score%1000 == 0 {
		return false
	}

	through := false
	nextX, nextY := x+dirs[dir][0], y+dirs[dir][1]
	for i := -1; i <= 1; i++ {
		nextDir := (dir + i + 4) % 4
		price := abs(i)*1000 + 1
		local := mark(nextX, nextY, nextDir, score-price, visited, field, spots)
		through = through || local
	}
	if through {
		spots[x][y] = true
	}
	return through
}
