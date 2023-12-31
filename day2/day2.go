package main

import (
	"bufio"
	// "errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileLines := getLines("input2")

	fmt.Println("Part 1:", part1(fileLines))
	fmt.Println("Part 2:", part2(fileLines))
}

func part1(lines []string) int {
	// red, green, blue
	maxValues := []int{12, 13, 14}
	total := 0

	for pos, line := range lines {
		line = strings.Split(line, ":")[1]
		games := strings.Split(line, ";")
		gamePossible := true

		for _, game := range games {
			colors := parseGame(game)
			if !possible(colors, maxValues) {
				gamePossible = false
			}
		}

		if gamePossible {
			total += (pos + 1)
		}
	}

	return total
}

func part2(lines []string) int {
	total := 0

	for _, line := range lines {
		line = strings.Split(line, ":")[1]
		games := strings.Split(line, ";")
		values := []int{0, 0, 0}

		for _, game := range games {
			colors := parseGame(game)
			values = findValues(values, colors)

		}
		total += multiplyColors(values)
	}
	return total
}

func multiplyColors(colors []int) int {
	return colors[0] * colors[1] * colors[2]
}

func findValues(previous []int, current []int) []int {
	updatedValues := []int{0, 0, 0}
	for i := 0; i < 3; i++ {
		if previous[i] < current[i] {
			updatedValues[i] = current[i]
		} else {
			updatedValues[i] = previous[i]
		}
	}
	return updatedValues
}

func possible(values []int, maxValues []int) bool {
	for i := 0; i < 3; i++ {
		if values[i] > maxValues[i] {
			return false
		}
	}
	return true
}

func parseGame(game string) []int {
	colors := []int{0, 0, 0}
	grabs := strings.Split(game, ",")
	for _, grab := range grabs {
		components := strings.Split(grab, " ")
		count := components[1]
		color := components[2]
		switch color {
		case "red":
			colors[0], _ = strconv.Atoi(count)
		case "green":
			colors[1], _ = strconv.Atoi(count)
		case "blue":
			colors[2], _ = strconv.Atoi(count)
		}
	}
	return colors
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
