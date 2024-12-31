package main

import (
	"image"
	"log"
	"os"
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

	c1, c2 := 0, 0
	for p := range grid {
		c1 += strings.Count(strings.Join(adjancency(p, 4), " "), "XMAS")
		c2 += strings.Count("AMAMASASAMAMAS", strings.Join(adjancency(p, 2)[4:], ""))
	}

	log.Printf("Count 1: %d", c1)
	log.Printf("Count 2: %d", c2)
}
