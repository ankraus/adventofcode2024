package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	counter := 0
	listA := make([]int, 0)
	listB := make([]int, 0)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		val, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Fatal(err)
		}
		if counter%2 == 0 {
			listA = insert(listA, val)
		} else {
			listB = insert(listB, val)
		}
		counter++
	}
	log.Printf("Total numbers: %d\n", counter)

	if len(listA) != len(listB) {
		log.Fatal("Slices are not the same length")
	}

	sum := 0

	for i := range len(listA) {
		diff := listA[i] - listB[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	log.Printf("Total sum of distances: %d", sum)
}

func insert(list []int, n int) []int {
	i := sort.SearchInts(list, n)
	result := slices.Insert(list, i, n)
	return result
}
