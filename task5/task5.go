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
	file, err := os.Open("task5/task5.txt")
	if err != nil {
		log.Fatalf("failed to open file task5.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task5.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	rulesRaw := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rulesRaw = append(rulesRaw, line)
	}
	rules := buildPre(rulesRaw)
	reverse := buildRev(rulesRaw)
	result1 := 0
	result2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Split(line, ",")
		val, ok := validate(elements, rules)
		result1 += val
		if !ok {
			result2 += fix(elements, reverse)
		}
	}
	fmt.Println(result1)
	fmt.Println(result2)
}

func buildPre(rulesRaw []string) map[string]map[string]struct{} {
	rules := make(map[string]map[string]struct{})
	for _, rule := range rulesRaw {
		x, y := strings.Split(rule, "|")[0], strings.Split(rule, "|")[1]

		if _, ok := rules[y]; !ok {
			rules[y] = make(map[string]struct{})
		}
		rules[y][x] = struct{}{}
	}
	return rules
}

func buildRev(rulesRaw []string) map[string]map[string]struct{} {
	rules := make(map[string]map[string]struct{})
	for _, rule := range rulesRaw {
		x, y := strings.Split(rule, "|")[0], strings.Split(rule, "|")[1]

		if _, ok := rules[x]; !ok {
			rules[x] = make(map[string]struct{})
		}
		rules[x][y] = struct{}{}
	}
	return rules
}

func validate(elements []string, rules map[string]map[string]struct{}) (int, bool) {
	set := make(map[string]struct{})
	for _, element := range elements {
		set[element] = struct{}{}
	}
	for _, element := range elements {
		for rule := range rules[element] {
			if _, ok := set[rule]; ok {
				return 0, false
			}
		}
		delete(set, element)
	}
	res, _ := strconv.Atoi(elements[len(elements)/2])
	return res, true
}

func fix(elements []string, rules map[string]map[string]struct{}) int {
	set := make(map[string]struct{})
	for _, element := range elements {
		set[element] = struct{}{}
	}
	inDegree := make(map[string]int)
	for from, toMap := range rules {
		if _, ok := set[from]; !ok {
			continue
		}
		for to := range toMap {
			if _, ok := set[to]; !ok {
				continue
			}
			inDegree[to]++
		}
	}
	res := make([]string, 0, len(elements))
	queue := make([]string, 0, len(elements))
	for _, element := range elements {
		if inDegree[element] == 0 {
			queue = append(queue, element)
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		res = append(res, cur)
		for next := range rules[cur] {
			if _, ok := set[next]; !ok {
				continue
			}
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	result, _ := strconv.Atoi(res[len(res)/2])
	return result
}
