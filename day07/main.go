package main

import (
	"bufio"
	"fmt"
	"os"
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

type command struct {
}

func parseCommands(inputLines []string) []command {
	return nil
}

func main() {
	//inputLines := readInput()
	//commands := parseCommands(inputLines)
	root := &directory{
		name:           "/",
		files:          make(map[string]*file),
		subDirectories: make(map[string]*directory),
		parent:         nil,
	}
	root.createSubDirectory("chch")
	root.subDirectories["chch"].createSubDirectory("ghgh")
	root.subDirectories["chch"].createFile("test", 1000)
	root.subDirectories["chch"].subDirectories["ghgh"].createFile("bigif", 2000)
	root.subDirectories["chch"].subDirectories["ghgh"].createFile("bigger", 3000)
	var pwd *directory
	pwd = root
	pwd = root.subDirectories["chch"]
	for _, f := range pwd.files {
		fmt.Println(f.name, f.size)
	}
	for _, d := range pwd.subDirectories {
		for _, f := range d.files {
			fmt.Println(f.name, f.size)
		}
	}
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
