package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numsAsStrs = []string{"one","two","three","four","five","six","seven","eight","nine"}

func check(e error) {
	if e != nil {
			fmt.Println(e)
			panic(e)
	}
}

func solveDay1() int {
	f, errOpenFile := os.Open("./inputs/day1.txt")
	check(errOpenFile)
	defer f.Close()
	
	sum := 0
	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		firstIntIndex, firstInt, lastIntIndex, lastInt := len(line) + 1, 0, -1, 0

		for i, char := range line {
			parsedInt, errInt := strconv.Atoi(string(char))
			
			if errInt != nil {
				continue
			}
			
			firstIntIndex = i
			firstInt = parsedInt
			break
		}
		
		for i := len(line) - 1; i >= 0; i-- {
			char := line[i]
			parsedInt, errInt := strconv.Atoi(string(char))
			
			if errInt != nil {
				continue
			}
			
			lastIntIndex = i
			lastInt = parsedInt
			break
		}

		for i, numAsStr := range numsAsStrs {
			indexInLine := strings.Index(line, numAsStr)

			if indexInLine == -1 {
				continue
			}
			
			if indexInLine < firstIntIndex {
				firstIntIndex = indexInLine
				firstInt = i + 1
			}

			lastIndexInLine := strings.LastIndex(line, numAsStr)

			if lastIndexInLine > lastIntIndex {
				lastIntIndex = lastIndexInLine
				lastInt = i + 1
			}
		}

		num, errAtoi := strconv.Atoi(strconv.Itoa(firstInt) + strconv.Itoa(lastInt))
		check(errAtoi)
		fmt.Println(num)
		
		sum += num
	}
	
	return sum
}