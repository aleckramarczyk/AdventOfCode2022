package main

import (
	"bufio"
	"flag"
	"fmt"
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
	return inputLines
}

func parseInput(input []string) {
	var cargoDiagram []string
	var onInstructions bool = false
	var rawInstructions []string
	for _, line := range input {
		if !onInstructions {
			if line == " 1   2   3   4   5   6   7   8   9 " {
				onInstructions = true
				continue
			} else {
				cargoDiagram = append(cargoDiagram, line)
			}
		} else {
			if line != "" {
				rawInstructions = append(rawInstructions, line)
			}
		}
	}
	parseCargo(cargoDiagram)
	parseInstructions(rawInstructions)
}

func parseCargo(cargoDiagram []string) {
	//Build stacks
	for i := 0; i < 9; i++ {
		cargoArea.stacks = append(cargoArea.stacks, new(stack))
	}
	for i := len(cargoDiagram) - 1; i >= 0; i-- {
		line := cargoDiagram[i]
		for n, s := range cargoArea.stacks {
			if n == 0 {
				crate := string(line[1])
				if crate != " " {
					s.crates = append(s.crates, crate)
				}
			} else {
				crate := string(line[(n*4)+1])
				if crate != " " {
					s.crates = append(s.crates, string(line[(n*4)+1]))
				}
			}
		}
	}
}

func parseInstructions(rawInstructions []string) {
	for _, line := range rawInstructions {
		newInstruction := new(instruction)
		line = strings.Replace(line, "move ", "", 1)
		line = strings.Replace(line, "from ", "", 1)
		line = strings.Replace(line, "to ", "", 1)
		values := (strings.Split(line, " "))
		newInstruction.numberOfCrates, _ = strconv.Atoi(values[0])
		newInstruction.fromStack, _ = strconv.Atoi(values[1])
		newInstruction.toStack, _ = strconv.Atoi(values[2])
		instructions = append(instructions, newInstruction)
	}
}

type stack struct {
	crates []string
}

func (s *stack) moveCratesFromStack(numberOfCrates int) (crates []string) {
	//Replace the uncommented code with commented to solve part 1
	/*
		for i := 0; i < numberOfCrates; i++ {
			crates = append(crates, s.crates[len(s.crates)-1])
			s.crates = s.crates[:len(s.crates)-1]
		}
		return crates
	*/
	crates = s.crates[(len(s.crates))-numberOfCrates:]
	s.crates = s.crates[:(len(s.crates))-numberOfCrates]
	return crates
}

func (s *stack) addCratesToStack(crates []string) {
	s.crates = append(s.crates, crates...)
}

type stackStructure struct {
	stacks []*stack
}

func (c *stackStructure) moveCargo(numberOfCrates int, fromStack int, toStack int) {
	fromStack--
	toStack--
	crates := c.stacks[fromStack].moveCratesFromStack(numberOfCrates)
	c.stacks[toStack].addCratesToStack(crates)
}

type instruction struct {
	fromStack      int
	toStack        int
	numberOfCrates int
}

var instructions []*instruction

var cargoArea stackStructure

func main() {
	Args.InputFile = flag.String("path", "input", "file path to read from")
	flag.Parse()
	inputLines := readInput()
	parseInput(inputLines)
	executeInstructions()
	fmt.Printf("%s\n", getOutput())
}

func executeInstructions() {
	for _, inst := range instructions {
		cargoArea.moveCargo(inst.numberOfCrates, inst.fromStack, inst.toStack)
	}
	for _, s := range cargoArea.stacks {
		fmt.Println(s.crates)
	}
}

func getOutput() (output string) {
	for _, s := range cargoArea.stacks {
		output += s.crates[len(s.crates)-1]
	}
	return output
}
