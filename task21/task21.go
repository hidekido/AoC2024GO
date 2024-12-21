// Package main -?
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("task21/task21.txt")
	if err != nil {
		log.Fatalf("failed to open file task21.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task21.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	codes := make([]string, 0)
	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}
	res := 0
	for _, code := range codes {
		path := solve(code, 25)
		val, _ := strconv.Atoi(code[:len(code)-1])
		fmt.Println(code, val, path, path*val)
		res += path * val
	}
	fmt.Println(res)
}

var nums = map[rune][]int{
	'7': {0, 0}, '8': {0, 1}, '9': {0, 2},
	'4': {1, 0}, '5': {1, 1}, '6': {1, 2},
	'1': {2, 0}, '2': {2, 1}, '3': {2, 2},
	'0': {3, 1}, 'A': {3, 2},
}

//379A

func solve(code string, post int) int {
	x, y := nums['A'][0], nums['A'][1]
	res := 0
	for _, num := range code {
		p := moveCode(x, y, nums[num][0], nums[num][1])
		local := 9223372036854775807
		for _, attempt := range p {
			attempt = append(attempt, 'A')
			val := postProcess(attempt, post)
			local = min(local, val)
		}
		res += local
		x, y = nums[num][0], nums[num][1]
	}
	return res
}

//v<A<AA>>^AAvA<^A>AvA^Av<<A>>^AAvA^Av<A>^AA<A>Av<A<A>>^AAAvA<^A>A
//<v<A>>^A<vA<A>>^AAvAA<^A>A<v<A>>^AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A   <<<

//<<^A^^A>>AvvvA

//^A <<^^A >>A vvvA <<<
//<A>A v<<AA>^AA>A vAA^A <vAAA>^A  <<<<<<
//<v<A>>^AvA^A <vA<AA>>^AAvA<^A>AAvA^A <vA>^AA<A>A<v<A>A>^AAAvA<^A>A  <<<<<<

//^A ^^<<A >>A vvvA
//<A>A <AAv<AA>>^A vAA^A v<AAA>^A
//v<<A>>^AvA^Av<<A>>^AAv<A<A>>^AAvAA<^A>Av<A>^AA<A>Av<A<A>>^AAAvA<^A>A

//	  +---+---+
//	  | ^ | A |
//+---+---+---+
//| < | v | > |
//+---+---+---+

var controls = map[rune][]int{
	'^': {0, 1}, 'A': {0, 2},
	'<': {1, 0}, 'v': {1, 1}, '>': {1, 2},
}

var memo = make(map[string]map[int]int)

func postProcess(path []rune, times int) int {
	if times == 0 {
		return len(path)
	}
	key := string(path)
	if memo[key] == nil {
		memo[key] = make(map[int]int)
	}
	if memo[key][times] != 0 {
		return memo[key][times]
	}
	x, y := controls['A'][0], controls['A'][1]
	res := 0
	for _, num := range path {
		p := moveControl(x, y, controls[num][0], controls[num][1])
		localRes := 9223372036854775807
		for _, attempt := range p {
			attempt = append(attempt, 'A')
			val := postProcess(attempt, times-1)
			localRes = min(localRes, val)
		}
		res += localRes
		x, y = controls[num][0], controls[num][1]
	}
	memo[key][times] = res
	return res
}

// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+
// | 4 | 5 | 6 |
// +---+---+---+
// | 1 | 2 | 3 |
// +---+---+---+
//
//		  | 0 | A |
//	   +---+---+
func moveCode(x, y, i, j int) [][]rune {
	paths := make([][]rune, 0)
	moves := make(map[rune]int, 0)
	// move right
	if y < j {
		moves['>'] += j - y
	}
	// move up
	if x > i {
		moves['^'] += x - i
	}
	// move down
	if x < i {
		moves['v'] += i - x
	}
	// move left
	if y > j {
		moves['<'] += y - j
	}
	if len(moves) == 1 {
		paths = append(paths, make([]rune, 0))
		for k, v := range moves {
			for t := 0; t < v; t++ {
				paths[0] = append(paths[0], k)
			}
		}
		return paths
	}
	for move, times := range moves {
		if move == '<' && y-times == 0 && x == 3 {
			continue
		}
		if move == 'v' && x+times == 3 && y == 0 {
			continue
		}
		first := make([]rune, 0, times)
		for t := 0; t < times; t++ {
			first = append(first, move)
		}
		for move2, times2 := range moves {
			if move == move2 {
				continue
			}
			second := make([]rune, 0, times2)
			for t := 0; t < times2; t++ {
				second = append(second, move2)
			}
			path := append(first, second...)
			paths = append(paths, path)
		}
	}
	return paths
}

func moveControl(x, y, i, j int) [][]rune {
	paths := make([][]rune, 0)
	moves := make(map[rune]int, 0)
	// move right
	if y < j {
		moves['>'] += j - y
	}
	// move up
	if x > i {
		moves['^'] += x - i
	}
	// move down
	if x < i {
		moves['v'] += i - x
	}
	// move left
	if y > j {
		moves['<'] += y - j
	}
	if len(moves) == 1 {
		paths = append(paths, make([]rune, 0))
		for k, v := range moves {
			for t := 0; t < v; t++ {
				paths[0] = append(paths[0], k)
			}
		}
		return paths
	}
	for move, times := range moves {
		if move == '<' && y-times == 0 && x == 0 {
			continue
		}
		if move == '^' && x-times == 0 && y == 0 {
			continue
		}
		first := make([]rune, 0, times)
		for t := 0; t < times; t++ {
			first = append(first, move)
		}
		for move2, times2 := range moves {
			if move == move2 {
				continue
			}
			second := make([]rune, 0, times2)
			for t := 0; t < times2; t++ {
				second = append(second, move2)
			}
			path := append(first, second...)
			paths = append(paths, path)
		}
	}
	if len(paths) == 0 {
		paths = append(paths, make([]rune, 0))
	}
	return paths
}
