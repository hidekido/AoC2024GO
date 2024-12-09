// Package main -?
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("task9/task9.txt")
	if err != nil {
		log.Fatalf("failed to open file task9.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task9.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(compress(line))
		fmt.Println(compress2(line))
	}
}

func compress(line string) int {
	values := make([]int, 0, len(line))
	for i := 0; i < len(line); i++ {
		value, _ := strconv.Atoi(string(line[i]))
		values = append(values, value)
	}
	right := len(values) - 1
	if len(values)%2 == 0 {
		right--
	}

	cells := make([]int, 0, len(values))
	sum := 0
	for i := 0; i < len(values); i++ {
		cells = append(cells, sum)
		sum += values[i]
	}
	left := 1
	memory := cells[1]
	checksum := 0
	for left < right {
		fill := min(values[left], values[right])
		for i := 0; i < fill; i++ {
			checksum += memory * right / 2
			memory++
		}
		values[left] -= fill
		values[right] -= fill
		if values[left] == 0 {
			left += 2
			memory = cells[left]
		}
		if values[right] == 0 {
			right -= 2
		}
	}
	for i := 0; i < left; i += 2 {
		memory = cells[i]
		for j := 0; j < values[i]; j++ {
			checksum += memory * i / 2
			memory++
		}
	}

	return checksum
}

func compress2(line string) int {
	values := make([]int, 0, len(line))
	for i := 0; i < len(line); i++ {
		value, _ := strconv.Atoi(string(line[i]))
		values = append(values, value)
	}
	right := len(values) - 1
	if len(values)%2 == 0 {
		right--
	}

	cells := make([]int, 0, len(values))
	sum := 0
	for i := 0; i < len(values); i++ {
		cells = append(cells, sum)
		sum += values[i]
	}
	memory := cells[1]
	checksum := 0
	for right > 1 {
		for i := 1; i < right; i += 2 {
			if values[i] >= values[right] {
				memory = cells[i]
				values[i] -= values[right]
				for j := 0; j < values[right]; j++ {
					checksum += memory * right / 2
					memory++
				}
				cells[i] = memory
				values[right] = 0
				break
			}
		}
		right -= 2
	}
	for i := 0; i < len(values); i += 2 {
		memory = cells[i]
		for j := 0; j < values[i]; j++ {
			checksum += memory * i / 2
			memory++
		}
	}
	return checksum
}
