// Package main -?
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("task4/task4.txt")
	if err != nil {
		log.Fatalf("failed to open file task4.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task4.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	space := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		space = append(space, []rune(line))
	}
	fmt.Println(searchSimple(space))
	fmt.Println(searchHard(space))
}

func searchSimple(space [][]rune) int {
	result := 0
	word := []rune("XMAS")
	for i := range space {
		for j := range space[i] {
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if x == 0 && y == 0 {
						continue
					}
					result += check(i, j, 0, x, y, space, word)
				}
			}
		}
	}
	return result
}

func check(i, j, index, x, y int, space [][]rune, word []rune) int {
	if i < 0 || i >= len(space) || j < 0 || j >= len(space[0]) || word[index] != space[i][j] {
		return 0
	}
	if len(word)-1 == index {
		return 1
	}
	return check(i+x, j+y, index+1, x, y, space, word)
}

func searchHard(space [][]rune) int {
	result := 0
	for i := range space {
		for j := range space[i] {
			if space[i][j] != 'A' {
				continue
			}
			if i == 0 || j == 0 || i == len(space)-1 || j == len(space[0])-1 {
				continue
			}
			conditions := 0
			if space[i-1][j-1] == 'M' && space[i+1][j+1] == 'S' {
				conditions++
			}
			if space[i-1][j-1] == 'S' && space[i+1][j+1] == 'M' {
				conditions++
			}
			if space[i-1][j+1] == 'M' && space[i+1][j-1] == 'S' {
				conditions++
			}
			if space[i-1][j+1] == 'S' && space[i+1][j-1] == 'M' {
				conditions++
			}
			if conditions == 2 {
				result++
			}
		}
	}
	return result
}
