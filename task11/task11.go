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
	file, err := os.Open("task11/task11.txt")
	if err != nil {
		log.Fatalf("failed to open file task11.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task11.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	result := 0
	memo := make(map[string]map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		for _, part := range parts {
			result += calc(part, 75, memo)
		}
	}
	fmt.Println(result)
}

func calc(part string, count int, memo map[string]map[int]int) int {
	if count == 0 {
		return 1
	}
	if _, ok := memo[part][count]; ok {
		return memo[part][count]
	}
	if _, ok := memo[part]; !ok {
		memo[part] = make(map[int]int)
	}
	switch {
	case part == "0":
		memo[part][count] = calc("1", count-1, memo)
	case len(part)%2 == 0:
		val, _ := strconv.Atoi(part[len(part)/2:])
		memo[part][count] = calc(part[:len(part)/2], count-1, memo) + calc(strconv.Itoa(val), count-1, memo)
	default:
		val, _ := strconv.Atoi(part)
		memo[part][count] = calc(strconv.Itoa(val*2024), count-1, memo)
	}
	return memo[part][count]
}
