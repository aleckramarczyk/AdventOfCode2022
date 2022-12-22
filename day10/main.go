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
	m.spritePos = registerPos
}

func (m *monitor) drawPixel(pos int) {
	row := pos / 40
	index := pos % 40
	m.display[row][index] = m.sprite[index]
}

func (m *monitor) printDisplay() {
	for _, row := range m.display {
		for _, pixel := range row {
			if pixel {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func executeInstructions(intructions []*instruction) (noteableValues map[int]int) {
	cpu := &processor{
		currentCycle: 1,
		register:     1,
		cycles:       make(map[int]int),
	}
	mon := &monitor{
		spritePos: 1,
	}
	mon.moveSprite(mon.spritePos)
	for _, inst := range intructions {
		if inst.inst == "noop" {
			cpu.executeNoop(mon)
		} else if inst.inst == "addx" {
			cpu.executeAddx(inst.value, mon)
		}
	}
	noteableValues = cpu.cycles
	mon.printDisplay()
	return
}

func (cpu *processor) executeNoop(m *monitor) {
	cpu.currentCycle++
	cpu.cycles[cpu.currentCycle] = cpu.register
	//m.drawPixel(cpu.currentCycle - 1)
}

func (cpu *processor) executeAddx(value int, m *monitor) {
	//m.drawPixel(cpu.currentCycle - 1)
	cpu.currentCycle++
	cpu.cycles[cpu.currentCycle] = cpu.register
	//m.drawPixel(cpu.currentCycle - 1)
	cpu.currentCycle++
	cpu.register += value
	cpu.cycles[cpu.currentCycle] = cpu.register
	fmt.Println(cpu.register)
	//m.moveSprite(cpu.register - 1)
}
