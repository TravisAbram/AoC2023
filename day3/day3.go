package main

import (
	"bufio"
	"unicode"
	// "errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	// "unicode"
	// "strings"
)

type Part struct {
	xStart     int
	xEnd       int
	y          int
	value      int
	scanYStart int
	scanYEnd   int
	scanXStart int
	scanXEnd   int
	valid      bool
}

func main() {
	fileLines := getLines("input")

	fmt.Println("Part 1:", part1(fileLines))
}

func part1(lines []string) int {
	parts := buildParts(lines)
	parts = setScanRanges(parts, lines)
	parts = validate(parts, lines)
	total := calculateTotal(parts)
	return total
}

func calculateTotal(parts []Part) int {
	total := 0
	for _, part := range parts {
		if part.valid {
			total += part.value
		}
	}
	return total
}

func setScanRanges(parts []Part, lines []string) []Part {
	for index, part := range parts {
		part.scanYStart = 0
		part.scanXStart = 0
		part.scanYEnd = len(lines) - 1
		part.scanXEnd = 140

		if part.y != 0 {
			part.scanYStart = part.y - 1
		}
		if part.y != len(lines)-1 {
			part.scanYEnd = part.y + 1
		}
		if part.xStart != 0 {
			part.scanXStart = part.xStart - 1
		}
		if part.xEnd != 140 {
			part.scanXEnd = part.xEnd + 1
		}
		parts[index] = part

	}
	return parts

}

func validate(parts []Part, lines []string) []Part {
	for index, part := range parts {
		for i := part.scanYStart; i <= part.scanYEnd; i++ {
			for j := part.scanXStart; j < part.scanXEnd; j++ {
				if unicode.IsDigit(rune(lines[i][j])) {
					continue
				}
				if string(lines[i][j]) == "." {
					continue
				}
				part.valid = true
			}
		}
		parts[index] = part
	}
	return parts
}

func buildParts(lines []string) []Part {
	parts := []Part{}

	for y, line := range lines {
		re := regexp.MustCompile("[0-9]+")
		idx := re.FindAllStringIndex(line, -1)
		for _, j := range idx {
			match, _ := strconv.Atoi(line[j[0]:j[1]])
			part := Part{xStart: j[0], xEnd: j[1], y: y, value: match}
			parts = append(parts, part)
		}
	}
	return parts
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
