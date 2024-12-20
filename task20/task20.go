// Package main -?
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("task20/task20.txt")
	if err != nil {
		log.Fatalf("failed to open file task20.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task20.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	field := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, []rune(line))
	}
	x, y := searchStart(field)
	baseTime := bfs(field, x, y, 'E')
	shortEnd := make([][]int, len(field))
	for i := range shortEnd {
		shortEnd[i] = make([]int, len(field[0]))
	}
	for i := range field {
		for j := range field[i] {
			if field[i][j] != '#' {
				shortEnd[i][j] = bfs(field, i, j, 'E')
			}
		}
	}
	shortStart := make([][]int, len(field))
	for i := range shortStart {
		shortStart[i] = make([]int, len(field[0]))
	}
	for i := range field {
		for j := range field[i] {
			if field[i][j] != '#' {
				shortStart[i][j] = bfs(field, i, j, 'S')
			}
		}
	}
	fmt.Println(baseTime)
	solve(field, baseTime, 2, 100, shortStart, shortEnd)
	solve(field, baseTime, 20, 100, shortStart, shortEnd)
}

func solve(field [][]rune, baseTime, dist, limit int, shortStart, shortEnd [][]int) {
	res1 := 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] == '#' {
				continue
			}
			for x := -dist; x <= dist; x++ {
				for y := -dist; y <= dist; y++ {
					if abs(x)+abs(y) > dist {
						continue
					}
					xx := i + x
					yy := j + y
					if xx < 0 || xx >= len(field) || yy < 0 || yy >= len(field[0]) || field[xx][yy] == '#' {
						continue
					}
					localDist := shortStart[i][j] + shortEnd[xx][yy] + abs(xx-i) + abs(yy-j)
					if baseTime-localDist >= limit {
						res1++
					}
				}
			}
		}
	}
	fmt.Println(res1)
}

var dirs = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func bfs(field [][]rune, i, j int, target rune) int {
	queue := make([][]int, 0)
	queue = append(queue, []int{i, j, 0})
	visited := make([][]bool, len(field))
	for i = 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(field[i]))
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if field[cur[0]][cur[1]] == target {
			return cur[2]
		}
		for _, dir := range dirs {
			x, y := cur[0]+dir[0], cur[1]+dir[1]
			if x < 0 || x >= len(field) || y < 0 || y >= len(field[0]) || visited[x][y] || field[x][y] == '#' {
				continue
			}
			visited[x][y] = true
			queue = append(queue, []int{x, y, cur[2] + 1})

		}
	}
	return 1_000_000_000
}

func abs(i int) int {
	return max(i, -i)
}

func searchStart(field [][]rune) (int, int) {
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] == 'S' {
				return i, j
			}
		}
	}
	return 0, 0
}
