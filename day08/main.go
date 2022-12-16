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

func convertInput(inputLines []string) {
	forest = make([][]*tree, len(inputLines))
	for i, line := range inputLines {
		rawTreeHeights := strings.Split(line, "")
		var trees []*tree
		for _, theight := range rawTreeHeights {
			height, _ := strconv.Atoi(theight)
			newTree := &tree{
				height: height,
			}
			trees = append(trees, newTree)
		}
		forest[i] = trees
	}
	mapRelationships()
}

func mapRelationships() {
	for rowIndex, row := range forest {
		for treeIndex, t := range row {
			if rowIndex == 0 {
				t.up = nil
				t.down = forest[1][treeIndex]
			} else if rowIndex == len(forest)-1 {
				t.down = nil
				t.up = forest[rowIndex-1][treeIndex]
			} else {
				t.up = forest[rowIndex-1][treeIndex]
				t.down = forest[rowIndex+1][treeIndex]
			}
			if treeIndex == 0 {
				t.left = nil
				t.right = forest[rowIndex][treeIndex+1]
			} else if treeIndex == len(row)-1 {
				t.right = nil
				t.left = forest[rowIndex][treeIndex-1]
			} else {
				t.right = forest[rowIndex][treeIndex+1]
				t.left = forest[rowIndex][treeIndex-1]
			}
		}
	}
}

type tree struct {
	height  int
	left    *tree
	right   *tree
	up      *tree
	down    *tree
	visited bool
	visible bool
}

var forest [][]*tree

func main() {
	inputLines := readInput()
	convertInput(inputLines)
	numberOfTreesVisibleFromEdges := getNumberOfTreesVisibleFromEdges()
	printForest()
	fmt.Println(numberOfTreesVisibleFromEdges)
}

func getNumberOfTreesVisibleFromEdges() (visibleTrees int) {
	for _, row := range forest {
		//Check from left
		visibleHeight := -1
		for _, t := range row {
			if t.height > visibleHeight {
				if !t.visited {
					visibleTrees++
					t.visible = true
				}
				visibleHeight = t.height
			}
		}
		//Check from right
		visibleHeight = -1
		for tIndex := len(row) - 1; tIndex >= 0; tIndex-- {
			if row[tIndex].height > visibleHeight {
				if !row[tIndex].visited {
					visibleTrees++
					row[tIndex].visible = true
				}
				visibleHeight = row[tIndex].height
			}
		}
	}
	for column := 0; column < len(forest[0]); column++ {
		//Check from top
		visibleHeight := -1
		for t := 0; t < len(forest); t++ {
			if forest[t][column].height > visibleHeight {
				if !forest[t][column].visited {
					visibleTrees++
					forest[t][column].visible = true
				}
				visibleHeight = forest[t][column].height
			}
		}
		//Check from bottom
		for t := len(forest) - 1; t >= 0; t-- {
			if forest[t][column].height > visibleHeight {
				if !forest[t][column].visited {
					visibleTrees++
					forest[t][column].visible = true
				}
				visibleHeight = forest[t][column].height
			}
		}
	}
	return
}

func printForest() {
	for _, row := range forest {
		for _, t := range row {
			if t.visible {
				colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, strconv.Itoa(t.height))
				fmt.Printf(colored + " ")
			} else {
				fmt.Printf("%d ", t.height)
			}
		}
		fmt.Printf("\n")
	}
}
