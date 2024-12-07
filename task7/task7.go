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
	file, err := os.Open("task7/task7.txt")
	if err != nil {
		log.Fatalf("failed to open file task7.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task7.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	result1 := 0
	result2 := 0
	for scanner.Scan() {
		eq := scanner.Text()
		f, s := verify(eq)
		result1 += f
		result2 += s
	}
	fmt.Println(result1)
	fmt.Println(result2)
}

func verify(eq string) (int, int) {
	split := strings.Split(eq, ":")
	target, _ := strconv.Atoi(split[0])
	right := strings.Split(split[1], " ")[1:]
	vals := make([]int, 0, len(right))
	for _, v := range right {
		num, _ := strconv.Atoi(v)
		vals = append(vals, num)
	}
	res1, res2 := 0, 0
	if rec(target, vals[0], vals[1:]) {
		res1 = target
	}
	if rec2(target, vals[0], vals[1:]) {
		res2 = target
	}
	return res1, res2
}

func rec(target, cur int, vals []int) bool {
	if len(vals) == 0 {
		return target == cur
	}
	if cur > target {
		return false
	}
	return rec(target, cur*vals[0], vals[1:]) || rec(target, cur+vals[0], vals[1:])
}

func rec2(target, cur int, vals []int) bool {
	if len(vals) == 0 {
		return target == cur
	}
	if cur > target {
		return false
	}
	val3, _ := strconv.Atoi(strconv.Itoa(cur) + strconv.Itoa(vals[0]))
	return rec2(target, cur*vals[0], vals[1:]) || rec2(target, cur+vals[0], vals[1:]) || rec2(target, val3, vals[1:])
}
