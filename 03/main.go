package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const doRegexStr = `^(.*?)(don't\(\))|do\(\)(.*?)(don't\(\))`
const mulRegexStr = `mul\(\d{1,3}\,\d{1,3}\)`
const paramsRegexStr = `\d{1,3}`

func main() {

	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	mulStrings := findTerms(string(f[:]))
	sum := 0
	for i := 1; i < len(mulStrings); i += 2 {
		a, errA := strconv.Atoi(mulStrings[i-1])
		b, errB := strconv.Atoi(mulStrings[i])
		if errA != nil || errB != nil {
			log.Fatal(errA, errB)
		}
		sum += a * b
	}
	log.Printf("%d", sum)
}

func findTerms(input string) []string {
	doRegex := regexp.MustCompile(doRegexStr)
	mulRegex := regexp.MustCompile(mulRegexStr)
	paramRegex := regexp.MustCompile(paramsRegexStr)
	input = strings.ReplaceAll(input, "\n", "")
	doStrs := doRegex.FindAllString(input, -1)
	mulStrs := mulRegex.FindAllString(strings.Join(doStrs, " "), -1)
	paramStrs := paramRegex.FindAllString(strings.Join(mulStrs, " "), -1)
	return paramStrs
}
