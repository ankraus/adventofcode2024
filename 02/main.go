package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	list := readReportsFromFile()
	count := 0
	for _, r := range list {
		if isSafe(r) {
			count++
		}
	}
	log.Printf("Safe count: %v", count)
}

func readReportsFromFile() [][]int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	list := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")

		result := make([]int, 0)
		for _, s := range nums {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, n)
		}
		list = append(list, result)
	}
	return list
}

func isSafe(report []int) bool {
	isDecreasing := report[0]-report[1] >= 1

	for i := 1; i < len(report); i++ {
		if report[i-1] == report[i] {
			return false
		}
		if report[i]-report[i-1] < 0 && !isDecreasing {
			return false
		}
		if report[i-1]-report[i] < 0 && isDecreasing {
			return false
		}
		difference := report[i-1] - report[i]
		if difference < 0 {
			difference = -difference
		}
		if difference > 3 {
			return false
		}
	}
	return true
}
