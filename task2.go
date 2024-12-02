// Package main ?
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type counter struct {
	mu    sync.Mutex
	value int
}

func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	file, _ := os.Open("task2.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task2.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	wg := sync.WaitGroup{}
	result := &counter{}
	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go func() {
			if check(parse(line), 1) {
				result.increment()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(result.getValue())
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
	return slices.Concat(vals[:index], vals[index+1:])
}

func testVals(first, second, inc int) bool {
	val := first - second
	abs := max(val, -val)
	return inc*val > 0 && abs >= 1 && abs <= 3
}
