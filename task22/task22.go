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
	file, err := os.Open("task22/task22.txt")
	if err != nil {
		log.Fatalf("failed to open file task22.txt %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file task22.txt %v", err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	res := 0
	secrets := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		secret, _ := strconv.Atoi(line)
		val := simulate(secret, 2000)
		secrets = append(secrets, secret)
		res += val
	}
	res2 := buy(secrets, 2000)
	fmt.Println(res)
	fmt.Println(res2)
}

var MODULO = 16777216

func buy(secrets []int, m_ax int) int {
	res := make(map[[4]int]int)
	preCalc := prices(secrets, m_ax)
	for s := range secrets {
		pr := preCalc[s]
		visited := make(map[[4]int]struct{})
		for i := 4; i < m_ax; i++ {
			key := [4]int{pr[i-3] - pr[i-4], pr[i-2] - pr[i-3], pr[i-1] - pr[i-2], pr[i] - pr[i-1]}
			if _, ok := visited[key]; ok {
				continue
			}
			visited[key] = struct{}{}
			res[key] += pr[i]
		}
	}
	result := 0
	for _, v := range res {
		result = max(result, v)
	}
	return result
}

func prices(secrets []int, num int) [][]int {
	result := make([][]int, len(secrets))
	for i, secret := range secrets {
		result[i] = make([]int, num+1)
		for j := 0; j <= num; j++ {
			result[i][j] = secret % 10
			secret = next(secret)
		}
	}
	return result
}

func simulate(secret int, number int) int {
	for i := 0; i < number; i++ {
		secret = next(secret)
	}
	return secret
}

func next(secret int) int {
	secret = ((secret * 64) ^ secret) % MODULO
	secret = ((secret / 32) ^ secret) % MODULO
	secret = ((secret * 2048) ^ secret) % MODULO
	return secret
}
