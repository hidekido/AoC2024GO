// Package main ->?
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("task6/task6.txt")
	if err != nil {
		log.Fatalf("failed to open file task6.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task6.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	field := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, []rune(line))
	}
	result1, result2 := track(field)

	fmt.Println(result1)
	fmt.Println(result2)
}

var dirs = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func makeVisited(n, m int) [][][]bool {
	visited := make([][][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([][]bool, m)
		for j := 0; j < m; j++ {
			visited[i][j] = make([]bool, 4)
		}
	}
	return visited
}

func track(field [][]rune) (int, int) {
	n := len(field)
	m := len(field[0])
	visited := makeVisited(n, m)
	i, j := searchStart(field)
	dir := 0
	result := 0
	loops := make([][]bool, n)
	loopsRes := 0
	for i := 0; i < n; i++ {
		loops[i] = make([]bool, m)
	}
	for i < n && i >= 0 && j < m && j >= 0 {
		for i < n && i >= 0 && j < m && j >= 0 {
			if field[i][j] == '#' {
				i -= dirs[dir][0]
				j -= dirs[dir][1]
				break
			}
			if isNotVisited(visited[i][j]) {
				result++
			}
			visited[i][j][dir] = true
			i += dirs[dir][0]
			j += dirs[dir][1]
			if i < n && i >= 0 && j < m && j >= 0 && field[i][j] != '#' && !loops[i][j] && isNotVisited(visited[i][j]) {
				prev := field[i][j]
				field[i][j] = '#'
				if checkLoop(i-dirs[dir][0], j-dirs[dir][1], (dir+1)%4, field) {
					loops[i][j] = true
					loopsRes++
				}
				field[i][j] = prev
			}
		}
		dir = (dir + 1) % 4
	}
	return result, loopsRes
}

func checkLoop(i, j, dir int, field [][]rune) bool {
	n := len(field)
	m := len(field[0])
	visited := makeVisited(n, m)
	for i < n && i >= 0 && j < m && j >= 0 {
		for i < n && i >= 0 && j < m && j >= 0 {
			if visited[i][j][dir] {
				return true
			}
			visited[i][j][dir] = true
			if field[i][j] == '#' {
				i -= dirs[dir][0]
				j -= dirs[dir][1]
				break
			}
			i += dirs[dir][0]
			j += dirs[dir][1]
		}
		dir = (dir + 1) % 4
	}
	return false
}

func searchStart(field [][]rune) (int, int) {
	for i := range field {
		for j := range field[i] {
			if field[i][j] == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}

func isNotVisited(vis []bool) bool {
	count := 0
	for _, ok := range vis {
		if ok {
			count++
		}
	}
	return count == 0
}
