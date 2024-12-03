// Package main -?
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("task3/task3.txt")
	if err != nil {
		log.Fatalf("failed to open file task3.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task3.txt %v", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	result1 := 0
	result2 := 0
	take := true
	for scanner.Scan() {
		line := scanner.Text()
		exps := parse(line)
		for _, exp := range exps {
			result1 += calc(exp)
		}
		exps2 := parse2(line)
		val := 0
		val, take = calcWhole(exps2, take)
		result2 += val
	}
	fmt.Println(result1)
	fmt.Println(result2)
}

func parse(line string) []string {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	return re.FindAllString(line, -1)
}

func parse2(line string) []string {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	return re.FindAllString(line, -1)
}

func calc(exp string) int {
	re := regexp.MustCompile(`\d{1,3}`)
	parts := re.FindAllString(exp, -1)
	first, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("Failed to parse %v - %v", first, err)
		return 0
	}
	second, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("Failed to parse %v - %v", second, err)
		return 0
	}
	return first * second
}

func calcWhole(exps []string, take bool) (int, bool) {
	result := 0
	for _, exp := range exps {
		switch {
		case exp == "do()":
			take = true
		case exp == "don't()":
			take = false
		case take:
			result += calc(exp)
		}
	}
	return result, take
}
