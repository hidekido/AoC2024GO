// Package main -?
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("task19/task19.txt")
	if err != nil {
		log.Fatalf("failed to open file task19.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task19.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	trie := Node{false, make(map[rune]*Node)}
	for scanner.Scan() {
		line := scanner.Text()
		for _, pattern := range strings.Split(line, ", ") {
			trie.Add(pattern)
		}
		break
	}
	res1 := 0
	res2 := 0
	for scanner.Scan() {
		word := scanner.Text()
		if word == "" {
			continue
		}

		if search1(trie, word, 0, make([]bool, len(word))) {
			res1++
		}
		memo := make([]int, len(word))
		for i := range memo {
			memo[i] = -1
		}
		res2 += search2(trie, word, 0, memo)
	}
	fmt.Println(res1)
	fmt.Println(res2)
}

func search1(trie Node, word string, index int, memo []bool) bool {
	if index == len(word) {
		return true
	}
	if memo[index] {
		return false
	}
	options := trie.Search(word, index)
	for _, option := range options {
		if search1(trie, word, index+option, memo) {
			return true
		}
	}
	memo[index] = true
	return false
}

func search2(trie Node, word string, index int, memo []int) int {
	if index == len(word) {
		return 1
	}
	if memo[index] != -1 {
		return memo[index]
	}
	options := trie.Search(word, index)
	res := 0
	for _, option := range options {
		res += search2(trie, word, index+option, memo)
	}
	memo[index] = res
	return res
}

type Node struct {
	isWord bool
	ch     map[rune]*Node
}

func (n *Node) Add(pattern string) {
	p := n
	for _, char := range pattern {
		if _, ok := p.ch[char]; !ok {
			p.ch[char] = &Node{false, make(map[rune]*Node)}
		}
		p = p.ch[char]
	}
	p.isWord = true
}

func (n *Node) Search(word string, start int) []int {
	res := make([]int, 0)
	p := n
	size := 0
	for index := start; index < len(word); index++ {
		char := rune(word[index])
		if next, ok := p.ch[char]; !ok {
			break
		} else {
			p = next
		}
		size++
		if p.isWord {
			res = append(res, size)
		}
	}
	return res
}
