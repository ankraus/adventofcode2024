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

	listA, listB := readListsFromFile()

	if len(listA) != len(listB) {
		log.Fatal("Slices are not the same length")
	}

	differenceSum := calculateDifferenceSum(listA, listB)

	log.Printf("Total sum of distances: %d", differenceSum)

	similarityScore := calculateSimilarityScore(listA, listB)

	log.Printf("Similarity Score: %d", similarityScore)
}

func readListsFromFile() ([]int, []int) {
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
	return listA, listB
}

func calculateDifferenceSum(listA []int, listB []int) int {
	sum := 0

	for i := range len(listA) {
		diff := listA[i] - listB[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}
	return sum
}

func calculateSimilarityScore(listA []int, listB []int) int {
	index := make(map[int]int)
	for _, n := range listB {
		index[n]++
	}

	sum := 0
	for _, n := range listA {
		sum += n * index[n]
	}
	return sum
}

func insert(list []int, n int) []int {
	i := sort.SearchInts(list, n)
	result := slices.Insert(list, i, n)
	return result
}
