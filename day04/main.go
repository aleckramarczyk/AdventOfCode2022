package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var Args struct {
	InputFile *string
}

func readInput() []string {
	file, err := os.Open(*Args.InputFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()
	var inputLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		panic(err)
	}
	return inputLines
}

type partner struct {
	startSection int
	endSection   int
}

func parseRange(input string) (int, int, error) {
	var err error
	sections := strings.Split(input, "-")
	startSection, err := strconv.Atoi(sections[0])
	if err != nil {
		return 0, 0, err
	}
	endSection, err := strconv.Atoi(sections[1])
	if err != nil {
		return 0, 0, err
	}
	return startSection, endSection, nil
}

type pair struct {
	partner0             *partner
	partner1             *partner
	partiallyOverlapping bool
	fullyOverlapping     bool
}

func (p *pair) determineOverlaps() {
	if math.Max(float64(p.partner0.startSection), float64(p.partner1.startSection)) <= math.Min(float64(p.partner0.endSection), float64(p.partner1.endSection)) {
		p.fullyOverlapping = false
		p.partiallyOverlapping = true
	}
	if p.partner0.startSection <= p.partner1.startSection && p.partner0.endSection >= p.partner1.endSection {
		p.fullyOverlapping = true
		p.partiallyOverlapping = false
	}
}

func part1(inputLines []string) {
	var pairs []*pair
	var numberOfFullyOverlappingPairs int
	var numberOfPartiallyOverlappingPairs int
	for _, line := range inputLines {
		partnerSections := strings.Split(line, ",")
		pair := &pair{
			partner0: new(partner),
			partner1: new(partner),
		}
		pair.partner0.startSection, pair.partner0.endSection, _ = parseRange(partnerSections[0])
		pair.partner1.startSection, pair.partner1.endSection, _ = parseRange(partnerSections[1])
		pair.determineOverlaps()
		pairs = append(pairs, pair)
	}
	for _, pair := range pairs {
		if pair.fullyOverlapping {
			numberOfFullyOverlappingPairs++
		} else if pair.partiallyOverlapping {
			numberOfPartiallyOverlappingPairs++
		}
	}
	fmt.Printf("The number of pairs where the two section ranges overlap is %d\n", numberOfFullyOverlappingPairs)
	fmt.Printf("The total number of pairs with an overlap is %d\n", numberOfFullyOverlappingPairs+numberOfPartiallyOverlappingPairs)
}

func main() {
	Args.InputFile = flag.String("path", "input", "file path to read from")
	flag.Parse()

	inputLines := readInput()

	part1(inputLines)
}
