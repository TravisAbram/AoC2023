package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fileLines := getLines("input1")

	total := 0
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

		total += value
	}

	fmt.Println("Part 1:", total)

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
