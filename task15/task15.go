// Package main -?
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("task15/task15.txt")
	if err != nil {
		log.Fatalf("failed to open file task15.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task15.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	field := make([][]rune, 0)
	field2 := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		field = append(field, []rune(line))
		field2 = append(field2, []rune(line))
	}
	moves := make([]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		moves = append(moves, []rune(line)...)
	}
	simulate(field, moves)
	result := calcDist(field)

	fmt.Println(result)
	modified := modify(field2)
	simulateBulky(modified, moves)
	result2 := calcDist2(modified)
	fmt.Println(result2)
}

var dirs = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func simulate(field [][]rune, moves []rune) {
	x, y := locateRobot(field)
	for _, mv := range moves {
		switch mv {
		case '>':
			x, y = move(x, y, dirs[0][0], dirs[0][1], field)
		case 'v':
			x, y = move(x, y, dirs[1][0], dirs[1][1], field)
		case '<':
			x, y = move(x, y, dirs[2][0], dirs[2][1], field)
		case '^':
			x, y = move(x, y, dirs[3][0], dirs[3][1], field)
		}
	}
}

func move(x, y, i, j int, field [][]rune) (int, int) {
	nextX, nextY := x+i, y+j
	switch field[nextX][nextY] {
	case '.':
		field[x][y] = '.'
		field[nextX][nextY] = '@'
		return nextX, nextY
	case '#':
		return x, y
	case 'O':
		placeX, placeY := searchEmpty(nextX, nextY, i, j, field)
		if placeX == nextX && placeY == nextY {
			return x, y
		}
		field[x][y] = '.'
		field[placeX][placeY] = 'O'
		field[nextX][nextY] = '@'
		return nextX, nextY
	}
	return x, y
}

func searchEmpty(x, y, i, j int, field [][]rune) (int, int) {
	xx, yy := x, y
	for field[x][y] == 'O' {
		x += i
		y += j
	}
	if field[x][y] != '.' {
		return xx, yy
	}
	return x, y
}

func locateRobot(field [][]rune) (x, y int) {
	for i := range field {
		for j := range field[i] {
			if field[i][j] == '@' {
				return i, j
			}
		}
	}
	return -1, -1
}

func calcDist(field [][]rune) int {
	res := 0
	for i := range field {
		for j := range field[i] {
			if field[i][j] == 'O' {
				res += i*100 + j
			}
		}
	}
	return res
}

func printField(field [][]rune) {
	for i := range field {
		fmt.Println(string(field[i]))
	}
}

func modify(field [][]rune) [][]rune {
	modified := make([][]rune, 0, len(field))
	for i := 0; i < len(field); i++ {
		row := make([]rune, 0, len(field[i]))
		for _, val := range field[i] {
			switch val {
			case '.':
				row = append(row, '.')
				row = append(row, '.')
			case '#':
				row = append(row, '#', '#')
			case 'O':
				row = append(row, '[', ']')
			case '@':
				row = append(row, '@', '.')
			}
		}
		modified = append(modified, row)
	}
	return modified
}

func simulateBulky(field [][]rune, moves []rune) {
	x, y := locateRobot(field)
	for _, mv := range moves {
		switch mv {
		case '>':
			x, y = moveBulky(x, y, dirs[0][0], dirs[0][1], field)
		case 'v':
			x, y = moveBulky(x, y, dirs[1][0], dirs[1][1], field)
		case '<':
			x, y = moveBulky(x, y, dirs[2][0], dirs[2][1], field)
		case '^':
			x, y = moveBulky(x, y, dirs[3][0], dirs[3][1], field)
		}
	}
}

func moveBulky(x, y, i, j int, field [][]rune) (int, int) {
	nextX, nextY := x+i, y+j
	switch field[nextX][nextY] {
	case '.':
		field[x][y] = '.'
		field[nextX][nextY] = '@'
		return nextX, nextY
	case '#':
		return x, y
	case '[', ']':
		if !checkCascade(nextX, nextY, i, j, field) {
			return x, y
		}
		moveCascade(nextX, nextY, i, j, field)
		field[nextX][nextY] = '@'
		field[x][y] = '.'
		return nextX, nextY
	}
	return x, y
}

func checkCascade(x, y, i, j int, field [][]rune) bool {
	if field[x][y] == '#' {
		return false
	}
	if field[x][y] == '.' {
		return true
	}
	leftX, leftY := x, y
	rightX, rightY := x, y
	if field[x][y] == '[' {
		rightX, rightY = x, y+1
	} else {
		leftX, leftY = x, y-1
	}
	if j == 1 {
		return checkCascade(rightX+i, rightY+j, i, j, field)
	}
	if j == -1 {
		return checkCascade(leftX+i, leftY+j, i, j, field)
	}
	return checkCascade(leftX+i, leftY+j, i, j, field) && checkCascade(rightX+i, rightY+j, i, j, field)
}

func moveCascade(x, y, i, j int, field [][]rune) {
	if field[x][y] == '.' {
		return
	}
	leftX, leftY := x, y
	rightX, rightY := x, y
	if field[x][y] == '[' {
		rightX, rightY = x, y+1
	} else {
		leftX, leftY = x, y-1
	}
	switch j {
	case 1:
		moveCascade(rightX+i, rightY+j, i, j, field)
	case -1:
		moveCascade(leftX+i, leftY+j, i, j, field)
	default:
		moveCascade(leftX+i, leftY+j, i, j, field)
		moveCascade(rightX+i, rightY+j, i, j, field)
	}
	field[leftX][leftY] = '.'
	field[rightX][rightY] = '.'
	field[leftX+i][leftY+j] = '['
	field[rightX+i][rightY+j] = ']'
}

func calcDist2(field [][]rune) int {
	res := 0
	for i := range field {
		for j := range field[i] {
			if field[i][j] == '[' {
				res += i*100 + j
			}
		}
	}
	return res
}
