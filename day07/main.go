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
	return inputLines
}

func initializeFileSystem() (root *directory) {
	return &directory{
		name:           "/",
		files:          make(map[string]*file),
		subDirectories: make(map[string]*directory),
		parent:         nil,
	}
}

type command struct {
	command  string
	argument string
	output   []string
}

func parseCommands(inputLines []string) (commands []command) {
	for _, line := range inputLines {
		var newCommand command
		if line[0] == '$' {
			newCommand.command = line[2:4]
			if newCommand.command == "cd" {
				newCommand.argument = line[5:]
				newCommand.output = nil
			} else if newCommand.command == "ls" {
				newCommand.argument = ""
			}
			commands = append(commands, newCommand)
		} else {
			previousCommand := &commands[(len(commands) - 1)]
			previousCommand.output = append(previousCommand.output, line)
		}
	}
	return commands
}

func (c *command) parseOutput() (files map[string]int, dirs []string) {
	files = make(map[string]int)
	for _, outputLine := range c.output {
		outputParts := strings.Split(outputLine, " ")
		if outputParts[0] == "dir" {
			dirs = append(dirs, outputParts[1])
		} else {
			files[outputParts[1]], _ = strconv.Atoi(outputParts[0])
		}
	}
	return
}

func executeCommands(commands []command) {
	//root := initializeFileSystem()

	//Create PWD object. Starts at root
	//pwd := root
	files, dirs := commands[1].parseOutput()
	fmt.Println(files, dirs)
}

func main() {
	commands := parseCommands(readInput())
	executeCommands(commands)
}

type directory struct {
	name           string
	files          map[string]*file
	subDirectories map[string]*directory
	parent         *directory
}

type file struct {
	name string
	size int
}

func (d *directory) createSubDirectory(name string) {
	newDirectory := &directory{
		name:           name,
		files:          make(map[string]*file),
		subDirectories: make(map[string]*directory),
		parent:         d,
	}
	d.subDirectories[name] = newDirectory
}

func (d *directory) createFile(name string, size int) {
	newFile := &file{
		name: name,
		size: size,
	}
	d.files[name] = newFile
}
