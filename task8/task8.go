// Package main -?
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("task8/task8.txt")
	if err != nil {
		log.Fatalf("failed to open file task8.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task8.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	field := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, []rune(line))
	}
	freqs := calcFreq(field)
	fmt.Println(search(len(field), len(field[0]), freqs))
	fmt.Println(search2(len(field), len(field[0]), freqs))
}

func calcFreq(field [][]rune) map[rune][][]int {
	freqs := make(map[rune][][]int)
	for i := range field {
		for j := range field[i] {
			if field[i][j] == '.' {
				continue
			}
			freqs[field[i][j]] = append(freqs[field[i][j]], []int{i, j})
		}
	}
	return freqs
}

func search(n, m int, freqs map[rune][][]int) int {
	marked := make([][]bool, n)
	for i := range marked {
		marked[i] = make([]bool, m)
	}
	res := 0
	for _, freq := range freqs {
		for i, first := range freq {
			for _, second := range freq[i+1:] {
				xOff := first[0] - second[0]
				yOff := first[1] - second[1]
				x := first[0] + xOff
				y := first[1] + yOff
				if x >= 0 && y >= 0 && x < len(marked) && y < len(marked[0]) && !marked[x][y] {
					marked[x][y] = true
					res++
				}
				x = second[0] - xOff
				y = second[1] - yOff
				if x >= 0 && y >= 0 && x < len(marked) && y < len(marked[0]) && !marked[x][y] {
					marked[x][y] = true
					res++
				}
			}
		}
	}
	return res
}

func search2(n, m int, freqs map[rune][][]int) int {
	marked := make([][]bool, n)
	for i := range marked {
		marked[i] = make([]bool, m)
	}
	res := 0
	for _, freq := range freqs {
		if len(freq) < 2 {
			continue
		}
		for i, first := range freq {
			if !marked[first[0]][first[1]] {
				res++
				marked[first[0]][first[1]] = true
			}
			for _, second := range freq[i+1:] {
				if !marked[second[0]][second[1]] {
					res++
					marked[second[0]][second[1]] = true
				}
				xOff := first[0] - second[0]
				yOff := first[1] - second[1]
				x := first[0] + xOff
				y := first[1] + yOff
				for x >= 0 && y >= 0 && x < len(marked) && y < len(marked[0]) {
					if !marked[x][y] {
						marked[x][y] = true
						res++
					}
					x += xOff
					y += yOff
				}
				x = second[0] - xOff
				y = second[1] - yOff
				for x >= 0 && y >= 0 && x < len(marked) && y < len(marked[0]) {
					if !marked[x][y] {
						marked[x][y] = true
						res++
					}
					x -= xOff
					y -= yOff
				}
			}
		}
	}
	return res
}
