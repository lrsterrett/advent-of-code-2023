package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func solveDay2() int {
	f, errOpenFile := os.Open("./inputs/day2.txt")
	check(errOpenFile)
	defer f.Close()

	maxes := map[string]int{"red": 12, "green": 13, "blue": 14}
	
	reInt := regexp.MustCompile(`[0-9]+`)
	reStr := regexp.MustCompile(`[a-z]+`)
	sum := 0

	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	
	lineloop:
	for scanner.Scan() {
		line := scanner.Text()
		gameNumAsStr := reInt.FindString(line)
		gameNum, errGameNum := strconv.Atoi(gameNumAsStr)
		check(errGameNum)
		game := strings.Split(line, ":")[1]
		grabs := strings.Split(game, ";")
		
		for _, grab := range grabs {
			cubeInfos := strings.Split(grab, ",")

			for _, cubeInfo := range cubeInfos {
				cubeCountAsStr := reInt.FindString(cubeInfo)
				cubeCount, errCubeCount := strconv.Atoi(cubeCountAsStr)
				check(errCubeCount)
				cubeColor := reStr.FindString(cubeInfo)

				if cubeCount > maxes[cubeColor] {
					continue lineloop
				}
			}
		}

		sum += gameNum
	}
	
	return sum
}