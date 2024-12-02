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
	file, _ := os.Open("task2.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task2.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		if check(parse(line), 1) {
			result++
		}
	}
	fmt.Println(result)
}

func parse(report string) []int {
	raw := strings.Split(report, " ")
	result := make([]int, len(raw))
	for i, v := range raw {
		result[i], _ = strconv.Atoi(v)
	}
	return result
}

func check(vals []int, tolerate int) bool {
	inc := vals[0] - vals[1]
	for index := 0; index < len(vals)-1; index++ {
		if testVals(vals[index], vals[index+1], inc) {
			continue
		}
		if tolerate > 0 {
			res := check(removeAt(vals, index), tolerate-1)
			res = res || check(removeAt(vals, index+1), tolerate-1)
			res = res || check(removeAt(vals, 0), tolerate-1)
			res = res || check(removeAt(vals, 1), tolerate-1)
			return res
		}
		return false
	}
	return true
}

func removeAt(vals []int, index int) []int {
	res := make([]int, 0, len(vals)-1)
	res = append(res, vals[:index]...)
	res = append(res, vals[index+1:]...)
	return res
}

func testVals(first, second, inc int) bool {
	val := first - second
	abs := max(val, -val)
	return inc*val > 0 && abs >= 1 && abs <= 3
}
