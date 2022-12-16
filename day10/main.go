package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() (inputLines []string) {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}
	return
}

type instruction struct {
	inst  string
	value int
}

func parseInstructions(inputLines []string) (instructions []*instruction) {
	for _, line := range inputLines {
		newInstruction := new(instruction)
		strs := strings.Split(line, " ")
		newInstruction.inst = strs[0]
		if len(strs) == 2 {
			newInstruction.value, _ = strconv.Atoi(strs[1])
		}
		instructions = append(instructions, newInstruction)
	}
	return
}

func main() {
	inputLines := readInput()
	instructions := parseInstructions(inputLines)
	cycles := executeInstructions(instructions)
	total := 0
	for i := 20; i <= 220; i += 40 {
		total += cycles[i] * i
		fmt.Println(cycles[i])
	}
	fmt.Println(total)
}

type processor struct {
	currentCycle int
	register     int
	cycles       map[int]int
}

type monitor struct {
	display   [6][40]bool
	sprite    [40]bool
	spritePos int
}

func (m *monitor) moveSprite(registerPos int) {
	for i := -1; i < 2; i++ {
		m.sprite[m.spritePos+i] = false
		m.sprite[registerPos+i] = true
	}
}

func (m *monitor) drawPixel(pos int) {
	row := pos / 40
	index := pos % 40
	lit := m.determineIfPixelIsLit(index)
}

func (m *monitor) determineIfPixelIsLit(index int) bool {
	if m.sprite[index] == true {
		return true
	} else {
		return false
	}
}

func executeInstructions(intructions []*instruction) (noteableValues map[int]int) {
	cpu := &processor{
		currentCycle: 1,
		register:     1,
		cycles:       make(map[int]int),
	}
	for _, inst := range intructions {
		if inst.inst == "noop" {
			cpu.executeNoop()
		} else if inst.inst == "addx" {
			cpu.executeAddx(inst.value)
		}
	}
	noteableValues = cpu.cycles
	return
}

func (cpu *processor) executeNoop() {
	cpu.currentCycle++
	cpu.cycles[cpu.currentCycle] = cpu.register
}

func (cpu *processor) executeAddx(value int) {
	cpu.currentCycle++
	cpu.cycles[cpu.currentCycle] = cpu.register
	cpu.currentCycle++
	cpu.register += value
	cpu.cycles[cpu.currentCycle] = cpu.register
}
