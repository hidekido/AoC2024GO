// Package main -?
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("task10/task10.txt")
	if err != nil {
		log.Fatalf("failed to open file task10.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task10.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	field := make([][]rune, 0)
	for scanner.Scan() {
		field = append(field, []rune(scanner.Text()))
	}

	result := 0
	result2 := 0
	n := len(field)
	m := len(field[0])
	memoAcc := memoInt(n, m)
	for i := range field {
		for j := range field[i] {
			if field[i][j] == '0' {
				result += traverse(i, j, field[i][j], field, memo(n, m))
				result2 += traverse2(i, j, field[i][j], field, memoAcc)
			}
		}
	}
	fmt.Println(result)
	fmt.Println(result2)
}

var dirs = [4][2]int{
	{0, 1},
	{1, 0},
	{-1, 0},
	{0, -1},
}

func traverse(i, j int, num rune, field [][]rune, memo [][]bool) int {
	if i < 0 || i >= len(field) || j < 0 || j >= len(field[0]) || field[i][j] != num {
		return 0
	}
	if memo[i][j] {
		return 0
	}
	memo[i][j] = true
	if num == '9' {
		return 1
	}
	res := 0
	for _, dir := range dirs {
		x := i + dir[0]
		y := j + dir[1]
		res += traverse(x, y, num+1, field, memo)
	}
	memo[i][j] = true
	return res
}

func traverse2(i, j int, num rune, field [][]rune, memo [][]int) int {
	if i < 0 || i >= len(field) || j < 0 || j >= len(field[0]) || field[i][j] != num {
		return 0
	}
	if num == '9' {
		return 1
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	res := 0
	for _, dir := range dirs {
		x := i + dir[0]
		y := j + dir[1]
		res += traverse2(x, y, num+1, field, memo)
	}
	memo[i][j] = res
	return res
}

func memo(n, m int) [][]bool {
	res := make([][]bool, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, make([]bool, m))
	}
	return res
}

func memoInt(n, m int) [][]int {
	res := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, make([]int, m))
		for j := 0; j < m; j++ {
			res[i][j] = -1
		}
	}
	return res
}
