package main

import (
	"bufio"
	"regexp"

	// "unicode"
	// "errors"
	"fmt"
	"log"
	"os"
	// "regexp"
	"math"
	"strconv"
	// "unicode"
	// "strings"
)

type Card struct {
	winning string
	chosen  string
	matches []int
	points  int
}

func main() {
	fileLines := getLines("input")

	fmt.Println("Part 1:", part1(fileLines))
	fmt.Println("Part 2:", part2(fileLines))
}

func part1(lines []string) int {
	cards := parseTickets(lines)
	cards = findMatches(cards)
	cards = calculatePoints(cards)
	return sumPoints(cards)
}

func sumPoints(cards []Card) int {
	var sum int
	for _, card := range cards {
		sum += card.points
	}
	return sum
}

func calculatePoints(cards []Card) []Card {
	for i, card := range cards {
		if len(card.matches) == 0 {
			card.points = 0
		} else {
			card.points = int(math.Pow(2, float64(len(card.matches))-1))
		}
		cards[i] = card
	}
	return cards
}

func findMatches(cards []Card) []Card {
	winningPattern := regexp.MustCompile(`\s+`)
	for i, card := range cards {
		winningRegex := `\b` + winningPattern.ReplaceAllString(card.winning, `\b|\b`) + `\b`
		pattern := regexp.MustCompile(winningRegex)
		matches := pattern.FindAllString(card.chosen, -1)
		for _, match := range matches {
			matchInt, _ := strconv.Atoi(match)
			card.matches = append(card.matches, matchInt)
		}
		cards[i] = card
	}
	return cards
}

func parseTickets(lines []string) []Card {
	var cards = make([]Card, len(lines))
	pattern := regexp.MustCompile(`^.*:\s+(.*)\s\|\s+(.*)$`)
	for i, line := range lines {
		matches := pattern.FindStringSubmatch(line)
		cards[i].winning = matches[1]
		cards[i].chosen = matches[2]
	}
	return cards
}

func part2(lines []string) int {
	total := 0
	return total
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
