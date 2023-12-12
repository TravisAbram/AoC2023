package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	fileLines := strings.Split(input, "\n")

	fmt.Println("Part 1:", part1(fileLines))
	fmt.Println("Part 2:", part2(fileLines))

}

func part1(lines []string) int {
	seeds, instructions := parseInput(lines)
	lowest := seeds[0]
	for _, seed := range seeds {
		for i := 1; i <= 8; i++ {
			for _, item := range instructions[i] {
				if seed >= item[1] && seed < (item[1]+item[2]) {
					seed = item[0] + (seed - item[1])
					break
				}
			}
		}
		if seed < lowest {
			lowest = seed
		}
	}
	return lowest
}

func part2(lines []string) int {
	return 0
}

func parseInput(lines []string) ([]int, map[int][][]int) {
	whichMap := 0
	instructions := make(map[int][][]int)
	pattern := regexp.MustCompile(`\b\d+\b`)
	seeds := []int{}
	seedsStr := pattern.FindAllString(lines[0], -1)
	for _, seedStr := range seedsStr {
		seed, _ := strconv.Atoi(seedStr)
		seeds = append(seeds, seed)
	}

	for _, line := range lines {
		if strings.Contains(line, "map") {
			whichMap++
			continue
		} else if whichMap == 0 || len(line) == 0 {
			continue
		}
		match := pattern.FindAllString(line, -1)
		destination, _ := strconv.Atoi(match[0])
		source, _ := strconv.Atoi(match[1])
		length, _ := strconv.Atoi(match[2])
		instruction := []int{destination, source, length}

		instructions[whichMap] = append(instructions[whichMap], instruction)
	}
	return seeds, instructions
}
