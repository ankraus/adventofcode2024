package main

import (
	"image"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	grid := map[image.Point]rune{}
	for y, s := range strings.Fields(string(f)) {
		for x, r := range s {
			grid[image.Point{x, y}] = r
		}
	}

	adjancency := func(p image.Point, size int) []string {
		directions := []image.Point{
			{0, -1}, {1, 0}, {0, 1}, {-1, 0},
			{-1, -1}, {1, -1}, {1, 1}, {-1, 1},
		}

		words := make([]string, len(directions))
		for i, d := range directions {
			for n := range size {
				words[i] += string(grid[p.Add(d.Mul(n))])
			}
		}
		return words
	}

	count := 0
	for p := range grid {
		count += strings.Count(strings.Join(adjancency(p, 4), " "), "XMAS")
	}

	log.Printf("Count: %d", count)
}

func findHorizontal(input []string) int {
	r := regexp.MustCompile(`(?:XMAS)|(?:SAMX)`)
	sum := 0
	for _, s := range input {
		sum += len(r.FindAllStringIndex(s, -1))
	}
	return sum
}
