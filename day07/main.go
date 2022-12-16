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

var root *directory

func executeCommands(commands []command) {
	//Create PWD object. Starts at root
	pwd := root
	for _, com := range commands {
		if com.command == "cd" {
			if com.argument == "/" {
				pwd = root
			} else if com.argument == ".." {
				pwd = pwd.parent
			} else {
				pwd = pwd.subDirectories[com.argument]
			}
		} else if com.command == "ls" {
			files, dirs := com.parseOutput()
			for name, size := range files {
				pwd.createFile(name, size)
			}
			for _, name := range dirs {
				pwd.createSubDirectory(name)
			}
		}
	}
}

func (d *directory) findDirectoriesWithSizeOfLessThan100000() int {
	var totalSize int
	for _, file := range d.files {
		totalSize += file.size
	}
	for _, dir := range d.subDirectories {
		totalSize += dir.findDirectoriesWithSizeOfLessThan100000()
	}
	if totalSize <= 100000 {
		directoriesMatchingSearchCriteria[d.name] = totalSize
	}
	return totalSize
}

func (d *directory) getSubdirectoryPath(name string) (path string) {
	path += d.name
	return
}

var directoriesMatchingSearchCriteria map[string]int

func main() {
	commands := parseCommands(readInput())
	root = initializeFileSystem()
	executeCommands(commands)
	directoriesMatchingSearchCriteria = make(map[string]int)
	root.findDirectoriesWithSizeOfLessThan100000()
	output := 0
	for name, size := range directoriesMatchingSearchCriteria {
		output += size
		fmt.Printf("%s : %d\n", name, size)
	}
	fmt.Println(output)
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
