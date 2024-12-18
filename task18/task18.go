// Package main -?
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("task18/task18.txt")
	if err != nil {
		log.Fatalf("failed to open file task18.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task18.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	grid := make([][]bool, 71)
	for i := range grid {
		grid[i] = make([]bool, 71)
	}
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		grid[x][y] = true
		if solve(grid) == -1 {
			fmt.Println("Broken: ", x, y)
		}
	}
	fmt.Println(solve(grid))
}

var dirs = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func solve(grid [][]bool) int {
	queue := make([][]int, 0)
	queue = append(queue, []int{0, 0, 0})
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}
	visited[0][0] = true
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur[0] == len(grid)-1 && cur[1] == len(grid[0])-1 {
			return cur[2]
		}
		for _, dir := range dirs {
			x, y := cur[0]+dir[0], cur[1]+dir[1]
			if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || visited[x][y] || grid[x][y] {
				continue
			}
			visited[x][y] = true
			queue = append(queue, []int{x, y, cur[2] + 1})
		}
	}
	return -1
}
