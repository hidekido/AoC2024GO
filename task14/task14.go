// Package main -?
package main

import (
	"bufio"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("task14/task14.txt")
	if err != nil {
		log.Fatalf("failed to open file task14.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task14.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	l1, l2, l3, l4 := 0, 0, 0, 0
	robots := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		sx, sy, vx, vy := parse(line)
		robots = append(robots, []int{sx, sy, vx, vy})
		r1 := predict(sx, sy, vx, vy, 101, 103)
		switch r1 {
		case 1:
			l1++
		case 2:
			l2++
		case 3:
			l3++
		case 4:
			l4++
		}
	}

	fmt.Println(l1 * l2 * l3 * l4)
	display(robots)
}

func parse(line string) (int, int, int, int) {
	parts := strings.Split(line, " ")
	start := strings.Split(strings.Split(parts[0], "=")[1], ",")
	velo := strings.Split(strings.Split(parts[1], "=")[1], ",")
	startx, _ := strconv.Atoi(start[0])
	starty, _ := strconv.Atoi(start[1])
	velox, _ := strconv.Atoi(velo[0])
	veloy, _ := strconv.Atoi(velo[1])
	return startx, starty, velox, veloy
}

func predict(sx, sy, vx, vy, n, m int) int {
	for i := 0; i < 100; i++ {
		sx, sy = next(sx, sy, vx, vy, n, m)
	}
	middleX, middleY := n/2, m/2
	switch {
	case sx < middleX && sy < middleY:
		return 1
	case sx < middleX && sy > middleY:
		return 3
	case sx > middleX && sy < middleY:
		return 2
	case sx > middleX && sy > middleY:
		return 4
	}
	return 0
}

func next(sx, sy, vx, vy, n, m int) (int, int) {
	sx += vx
	sy += vy
	if sx < 0 {
		sx += n
	}
	if sy < 0 {
		sy += m
	}
	if sx >= n {
		sx -= n
	}
	if sy >= m {
		sy -= m
	}
	return sx, sy
}

type Image struct {
	data [][]uint8
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(im.data), len(im.data[0]))
}

func (im Image) At(x, y int) color.Color {
	if im.data[x][y] == 0 {
		return color.Black
	}
	return color.RGBA{0, 255, 0, 255}
}

func display(robots [][]int) {
	times := 0
	var field [][]uint8
	for times < 11000 {
		field = make([][]uint8, 103)
		for i := range field {
			field[i] = make([]uint8, 101)
		}
		for _, robot := range robots {
			robot[0], robot[1] = next(robot[0], robot[1], robot[2], robot[3], 101, 103)
			field[robot[1]][robot[0]]++
		}
		line := countLine(field)
		if line > 10 {
			m := Image{field}
			imaging.Save(m, fmt.Sprintf("./result/task14_%d.png", times))
			fmt.Println(times)
		}

		times++
	}
}

func countLine(field [][]uint8) int {
	res := 0
	for j := 0; j < 101; j++ {
		localRes := 0
		for i := 0; i < 103; i++ {
			if field[i][j] == 0 {
				localRes = 0
				continue
			}
			localRes++
			res = max(res, localRes)
		}
	}
	return res
}
