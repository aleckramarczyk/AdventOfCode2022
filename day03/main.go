package main

import (
	"bufio"
	"flag"
	"os"
)

var Args struct {
	fileName *string
}

func readInput() []string {
	input, err := os.Open(*Args.fileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = input.Close(); err != nil {
			panic(err)
		}
	}()

	var lines []string

	s := bufio.NewScanner(input)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	err = s.Err()
	if err != nil {
		panic(err)
	}

	return lines
}

func makePriorityList() map[byte]int {
	priorityList := make(map[byte]int)
	for i, p := 95, 1; i <= 122; i++ {
		priorityList[byte(i)] = p
		p++
	}
	for i, p := 65, 27; i <= 90; i++ {
		priorityList[byte(i)] = p
		p++
	}
	return priorityList
}

var PriorityList map[byte]int

type rucksack struct {
	compartment0 string
	compartment1 string
}

func part1(rucksacks []rucksack) {
	var commonItems []byte
	for i := 95; i <= 122; i++ {
		commonItems = append(commonItems, byte(i))
	}
	for i := 65; i <= 90; i++ {
		commonItems = append(commonItems, byte(i))
	}

}

func main() {
	//Parse arguments
	Args.fileName = flag.String("fpath", "input", "file path to read from")
	flag.Parse()

	//Read lines from input file into inputLines
	inputLines := readInput()

	//Create a list of all rucksacks from inputLines, splitting their contents into two compartments
	var rucksacks []rucksack
	for _, line := range inputLines {
		var currentRucksack rucksack
		currentRucksack.compartment0 = line[0:(len(line) / 2)]
		currentRucksack.compartment1 = line[(len(line) / 2):]
		rucksacks = append(rucksacks, currentRucksack)
	}

	PriorityList = makePriorityList()

	part1(rucksacks)
}
