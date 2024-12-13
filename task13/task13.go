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
	file, err := os.Open("task13/task13.txt")
	if err != nil {
		log.Fatalf("failed to open file task13.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task13.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		input = append(input, line)
	}
	res := 0
	res2 := 0
	for i := 0; i < len(input)/3; i++ {
		ax, ay := parseButton(input[i*3])
		bx, by := parseButton(input[i*3+1])
		px, py := parseCoordinates(input[i*3+2])
		tokens := findMinCostSolution(ax, ay, bx, by, px, py)
		val := findMinCostSolution(ax, ay, bx, by, px+10000000000000, py+10000000000000)
		res2 += val
		res += tokens
	}
	fmt.Println(res)
	fmt.Println(res2)
}

func parseCoordinates(input string) (int, int) {
	parts := strings.Split(input, ", ")
	xPart := strings.Split(parts[0], "=")[1]
	yPart := strings.Split(parts[1], "=")[1]

	x, _ := strconv.Atoi(xPart)
	y, _ := strconv.Atoi(yPart)

	return x, y
}

func parseButton(input string) (int, int) {
	parts := strings.Split(input, ", ")
	x, _ := strconv.Atoi(strings.Split(parts[0], "+")[1])
	y, _ := strconv.Atoi(strings.Split(parts[1], "+")[1])

	return x, y
}

func findMinCostSolution(x1, y1, x2, y2, x3, y3 int) int {
	// x1 * a + x2 * b = x3
	// y1 * a + y2 * b = y3
	// a = (x3 - x2 * b) / x1
	// b = (y3 - y1 * a) / y2
	// y3 = y1 * (x3 - x2 * b) / x1 + y2 * b
	// y3 = y1 * x3 / x1 - y1 * x2 * b / x1 + y2 * b
	// b = (y3 - y1 * x3 / x1) / (y2 - y1 * x2 / x1)
	// b = (y3 * x1 - y1 * x3) / (y2 * x1 - y1 * x2)

	if (y3*x1-y1*x3)%(y2*x1-y1*x2) != 0 {
		return 0
	}
	b := (y3*x1 - y1*x3) / (y2*x1 - y1*x2)
	if (x3-x2*b)%x1 != 0 {
		return 0
	}
	a := (x3 - x2*b) / x1
	return a*3 + b
}
