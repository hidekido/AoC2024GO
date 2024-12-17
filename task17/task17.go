// Package main -?
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("task17/task17.txt")
	if err != nil {
		log.Fatalf("failed to open file task17.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task17.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	a, b, c := 0, 0, 0
	if scanner.Scan() {
		a, _ = strconv.Atoi(strings.Split(scanner.Text(), " ")[2])
		b, _ = strconv.Atoi(strings.Split(scanner.Text(), " ")[2])
		c, _ = strconv.Atoi(strings.Split(scanner.Text(), " ")[2])
	}
	if scanner.Scan() {
		b, _ = strconv.Atoi(strings.Split(scanner.Text(), " ")[2])
	}
	if scanner.Scan() {
		c, _ = strconv.Atoi(strings.Split(scanner.Text(), " ")[2])
	}
	if scanner.Scan() {
		scanner.Text()
	}
	var commands []int
	if scanner.Scan() {
		cs := strings.Split(strings.Split(scanner.Text(), " ")[1], ",")
		commands = make([]int, len(cs))
		for i, comm := range cs {
			commands[i], _ = strconv.Atoi(comm)
		}
	}
	fmt.Println(a, b, c, commands)
	res := eval(a, b, c, commands)
	fmt.Println(res)

	rec(0, len(commands)-1, commands)
}

func rec(a, index int, commands []int) {
	if index == -1 {
		fmt.Println(a)
		return
	}
	for add := 0; add < 8; add++ {
		next := a*8 + add
		if single(next, 0, 0, commands) == commands[index] {
			rec(next, index-1, commands)
		}
	}
}

func eval(a, b, c int, commands []int) []int {
	var rec func(index int)
	res := make([]int, 0)
	rec = func(index int) {
		if index >= len(commands)-1 {
			return
		}
		switch commands[index] {
		case 0:
			second := combo(commands[index+1], a, b, c)
			second = int(math.Pow(2, float64(second)))
			a = a / second
		case 1:
			b = b ^ commands[index+1]
		case 2:
			b = combo(commands[index+1], a, b, c) % 8
		case 3:
			if a != 0 {
				index = commands[index+1] - 2
			}
		case 4:
			b = b ^ c
		case 5:
			val := combo(commands[index+1], a, b, c) % 8
			res = append(res, val)
		case 6:
			second := combo(commands[index+1], a, b, c)
			second = int(math.Pow(2, float64(second)))
			b = a / second
		case 7:
			second := combo(commands[index+1], a, b, c)
			second = int(math.Pow(2, float64(second)))
			c = a / second
		}
		rec(index + 2)
	}
	rec(0)

	return res
}

func combo(num, a, b, c int) int {
	switch num {
	case 0, 1, 2, 3:
		return num
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	}
	return 0
}

func single(a, b, c int, commands []int) int {
	var rec func(index int) int
	rec = func(index int) int {
		switch commands[index] {
		case 0:
			second := combo(commands[index+1], a, b, c)
			second = int(math.Pow(2, float64(second)))
			a = a / second
		case 1:
			b = b ^ commands[index+1]
		case 2:
			b = combo(commands[index+1], a, b, c) % 8
		case 3:
			if a != 0 {
				index = commands[index+1] - 2
			}
		case 4:
			b = b ^ c
		case 5:
			return combo(commands[index+1], a, b, c) % 8
		case 6:
			second := combo(commands[index+1], a, b, c)
			second = int(math.Pow(2, float64(second)))
			b = a / second
		case 7:
			second := combo(commands[index+1], a, b, c)
			second = int(math.Pow(2, float64(second)))
			c = a / second
		}
		return rec(index + 2)
	}

	return rec(0)
}
