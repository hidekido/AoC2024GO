// Package main -?
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("task23/task23.txt")
	if err != nil {
		log.Fatalf("failed to open file task23.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task23.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	graph := make(map[string]map[string]struct{})
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")
		first, second := line[0], line[1]
		if _, ok := graph[first]; !ok {
			graph[first] = make(map[string]struct{})
		}
		if _, ok := graph[second]; !ok {
			graph[second] = make(map[string]struct{})
		}
		graph[first][second] = struct{}{}
		graph[second][first] = struct{}{}
	}

	res := 0
	used := make(map[string]struct{})
	for k := range graph {
		if k[0] == 't' {
			res += search(k, graph, used)
			used[k] = struct{}{}
		}
	}
	fmt.Println(res)
	fmt.Println(getLargest(graph))
}

func getLargest(graph map[string]map[string]struct{}) string {
	res := make([]string, 0)
	var rec func(int, []string)
	comp := make([]string, len(graph))
	index := 0
	for k := range graph {
		comp[index] = k
		index++
	}
	rec = func(cur int, group []string) {
		if cur >= len(comp) {
			if len(group) > len(res) {
				res = make([]string, len(group))
				copy(res, group)

			}
			return
		}
		rec(cur+1, group)
		add := true
		for _, c := range group {
			if _, ok := graph[c][comp[cur]]; !ok {
				add = false
				break
			}
		}
		if add {
			group = append(group, comp[cur])
			rec(cur+1, group)
			group = group[:len(group)-1]
		}
	}

	rec(0, []string{})
	slices.Sort(res)
	return strings.Join(res, ",")
}

func search(k string, graph map[string]map[string]struct{}, used map[string]struct{}) int {
	groups := make(map[string]map[string]struct{})
	res := 0
	for first := range graph[k] {
		if _, ok := used[first]; ok {
			continue
		}
		for second := range graph[k] {
			if _, ok := used[second]; ok {
				continue
			}
			if first == second {
				continue
			}
			if _, ok := graph[first][second]; !ok {
				continue
			}
			if _, ok := groups[first][second]; ok {
				continue
			}
			res++
			if _, ok := groups[first]; !ok {
				groups[first] = make(map[string]struct{})
			}
			if _, ok := groups[second]; !ok {
				groups[second] = make(map[string]struct{})
			}
			groups[first][second] = struct{}{}
			groups[second][first] = struct{}{}
		}
	}
	return res
}
