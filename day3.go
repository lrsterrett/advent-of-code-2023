package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solveDay3Part1() int {
	f, errOpenFile := os.Open("./inputs/day3.txt")
	check(errOpenFile)
	
	var (
		data [][]string
		sum int
	)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, strings.Split(scanner.Text(), ""))
	}
	
	for i, row := range data {
		for j, char := range row {
			if char != "." && (char < "0" || char > "9") {
				sum += sumNeighboringNums(data, i, j)
			}
		}
	}

	return sum
}

func sumNeighboringNums(data [][]string, i int, j int) int {
	var (
		matchingNeighbors []string
		sum int
	)

	if i > 1 {
		matchingNeighbors = append(matchingNeighbors, data[i-1][j])
		if j > 1 {
			matchingNeighbors = append(matchingNeighbors, data[i-1][j-1])
		}
		if j < len(data[i]) - 1 {
			matchingNeighbors = append(matchingNeighbors, data[i-1][j+1])
		}
	}
	if i < len(data) - 1 {
		matchingNeighbors = append(matchingNeighbors, data[i+1][j])
		if j > 1 {
			matchingNeighbors = append(matchingNeighbors, data[i+1][j-1])
		}
		if j < len(data[i]) - 1 {
			matchingNeighbors = append(matchingNeighbors, data[i+1][j+1])
		}
	}

	for _, char := range matchingNeighbors {
		num, errConvertToNum := strconv.Atoi(char)
		if errConvertToNum != nil {
			continue
		}
		sum += num
	}

	return sum
}