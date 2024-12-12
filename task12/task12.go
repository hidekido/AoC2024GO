// Package main -?
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("task12/task12.txt")
	if err != nil {
		log.Fatalf("failed to open file task12.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task12.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	field := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, []rune(line))
	}
	res := 0
	for i := range field {
		for j := range field[i] {
			if field[i][j] == ' ' {
				continue
			}
			visited := makeVisited(len(field), len(field[0]))
			area := mark(i, j, field[i][j], field, visited)
			sides := countSides(visited)
			res += area * sides / 2
		}
	}
	fmt.Println(res)
}

var dirs = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func markAndReturn(i, j int, kind rune, field [][]rune, visited [][]bool) (int, int) {
	if field[i][j] != kind {
		return 0, 1
	}
	area, per := 1, 0
	visited[i][j] = true
	field[i][j] = ' '
	for _, dir := range dirs {
		x := i + dir[0]
		y := j + dir[1]
		if x < 0 || x >= len(field) || y < 0 || y >= len(field[0]) {
			per++
			continue
		}
		if visited[x][y] {
			continue
		}
		a, p := markAndReturn(x, y, kind, field, visited)
		area += a
		per += p
	}
	return area, per
}

func mark(i, j int, kind rune, field [][]rune, visited [][]bool) int {
	if field[i][j] != kind {
		return 0
	}
	res := 1
	visited[i][j] = true
	field[i][j] = ' '
	for _, dir := range dirs {
		x := i + dir[0]
		y := j + dir[1]
		if x < 0 || x >= len(field) || y < 0 || y >= len(field[0]) {
			continue
		}
		if visited[x][y] {
			continue
		}
		res += mark(x, y, kind, field, visited)
	}
	return res
}

func countSides(visited [][]bool) int {
	res := 0
	for i := range visited {
		for j := range visited[i] {
			if !visited[i][j] {
				continue
			}
			res += count(i, j, visited)
		}
	}
	return res
}

func count(i, j int, visited [][]bool) int {
	siLocal := 0
	for index, dir := range dirs {
		x := i + dir[0]
		y := j + dir[1]
		if endOrDiff(x, y, visited) {
			xx := i + dirs[(index+1)%4][0]
			yy := j + dirs[(index+1)%4][1]
			if endOrDiff(xx, yy, visited) {
				siLocal += 2
			}
		}
	}

	for index, dir := range dirs {
		x := i + dir[0]
		y := j + dir[1]
		if endOrDiff(x, y, visited) {
			continue
		}
		xx := i + dirs[(index+1)%4][0]
		yy := j + dirs[(index+1)%4][1]
		if endOrDiff(xx, yy, visited) {
			continue
		}
		xxx := i + dirs[(index+1+4)%4][0] + dir[0]
		yyy := j + dirs[(index+1+4)%4][1] + dir[1]
		if visited[xxx][yyy] {
			continue
		}
		siLocal += 2
	}
	return siLocal
}

func endOrDiff(x, y int, visited [][]bool) bool {
	return x < 0 || x >= len(visited) || y < 0 || y >= len(visited[0]) || !visited[x][y]
}

func makeVisited(n, m int) [][]bool {
	visited := make([][]bool, 0, n)
	for i := 0; i < n; i++ {
		visited = append(visited, make([]bool, m))
	}
	return visited
}
