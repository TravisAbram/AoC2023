package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Card struct {
	winning string
	chosen  string
	matches []int
	points  int
	copies  int
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

func part2(lines []string) int {
	cards := parseTickets(lines)
	cards = findMatches(cards)
	cards = calculateCopies(cards)
	return sumCopies(cards)
}

func calculateCopies(cards []Card) []Card {
	for i := 0; i < len(cards); i++ {
		cards[i].copies = 1
	}

	for i, card := range cards {
		wins := len(card.matches)

		for k := 0; k < cards[i].copies; k++ {
			for j := 0; j < wins; j++ {
				cards[i+j+1].copies += 1
			}
		}
	}
	return cards
}

func sumCopies(cards []Card) int {
	var sum int
	for _, card := range cards {
		sum += card.copies
	}
	return sum
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
