package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fileLines := getLines("input1")

	part1 := 0
	for pos, line := range fileLines {
		first, err := findFirstDigit(line)
		if err != nil {
			log.Fatal(pos, err)
		}

		last, err := findLastDigit(line)
		if err != nil {
			log.Fatal(pos, err)
		}

		combined := fmt.Sprintf("%s%s", string(first), string(last))

		value, err := strconv.Atoi(combined)
		if err != nil {
			log.Fatal(pos, err)
		}

		part1 += value
	}

	fmt.Println("Part 1:", part1)

	part2 := 0
	for pos, line := range fileLines {
		line = replaceDigitWords(line)

		first, err := findFirstDigit(line)
		if err != nil {
			log.Fatal(pos, err)
		}

		last, err := findLastDigit(line)
		if err != nil {
			log.Fatal(pos, err)
		}

		combined := fmt.Sprintf("%s%s", string(first), string(last))

		value, err := strconv.Atoi(combined)
		if err != nil {
			log.Fatal(pos, err)
		}

		part2 += value
	}

	fmt.Println("Part 2:", part2)
}

func getLines(file string) []string {
	readFile, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	return fileLines
}

func findFirstDigit(line string) (int, error) {
	runeLine := []rune(line)
	for i := 0; i < len(runeLine); i++ {
		if unicode.IsDigit(runeLine[i]) {
			return int(runeLine[i]), nil
		}
	}
	return 0, errors.New("no digit found")
}

func findLastDigit(line string) (int, error) {
	runeLine := []rune(line)
	for i := len(runeLine) - 1; i >= 0; i-- {
		if unicode.IsDigit(runeLine[i]) {
			return int(runeLine[i]), nil
		}
	}
	return 0, errors.New("no digit found")
}

func replaceDigitWords(line string) string {
	fixedLine := ""

	fixedLine = strings.Replace(line, "one", "o1e", -1)
	fixedLine = strings.Replace(fixedLine, "two", "t2o", -1)
	fixedLine = strings.Replace(fixedLine, "three", "t3e", -1)
	fixedLine = strings.Replace(fixedLine, "four", "f4r", -1)
	fixedLine = strings.Replace(fixedLine, "five", "f5e", -1)
	fixedLine = strings.Replace(fixedLine, "six", "s6x", -1)
	fixedLine = strings.Replace(fixedLine, "seven", "s7n", -1)
	fixedLine = strings.Replace(fixedLine, "eight", "e8t", -1)
	fixedLine = strings.Replace(fixedLine, "nine", "n9e", -1)

	return fixedLine
}
