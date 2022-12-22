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
	return
}

type graph struct {
	vertices    map[coordinate]*vertex
	startCoors  coordinate
	endCoors    coordinate
	xDimension  int
	yDimensions int
}

type vertex struct {
	height    int
	isStart   bool
	isEnd     bool
	neighbors []*vertex
}

type coordinate struct {
	x int
	y int
}

func makeHeightMap() map[rune]int {
	//Returns a mapping of characters to their corresponding heights for easy lookup when building the graph
	heightMap := make(map[rune]int)
	height := 0
	//Lowercase characters
	for i := 97; i <= 122; i++ {
		heightMap[rune(i)] = height
		height++
	}
	//S
	heightMap[rune(83)] = 0
	//E
	heightMap[rune(69)] = 25

	return heightMap
}

func buildGraph(inputLines []string) *graph {
	//Creates a graph given a slice of strings. Used to turn the input of the challenge into a datastructure
	heightMap := makeHeightMap() //Heightmap is used to quickly get the height of a vertex based on a character
	mapGraph := &graph{
		vertices: make(map[coordinate]*vertex),
	}
	for y, row := range inputLines {
		for x, point := range row {
			coordinates := coordinate{
				x: x,
				y: y,
			}
			vert := &vertex{
				height:  heightMap[point],
				isStart: (point == 'S'),
				isEnd:   (point == 'E'),
			}
			if vert.isStart {
				mapGraph.startCoors = coordinates
			}
			if vert.isEnd {
				mapGraph.endCoors = coordinates
			}
			mapGraph.vertices[coordinates] = vert
		}
	}

	//Set the dimensions of mapGraph
	mapGraph.yDimensions = len(inputLines)
	mapGraph.xDimension = len(inputLines[0])

	mapRelationships(mapGraph)
	return mapGraph
}

func mapRelationships(mapGraph *graph) {
	for coors, vert := range mapGraph.vertices {
		//If the vertex is not on the left edge of the map, add the vertex to the left to this vertex's list of neighbors
		if coors.x-1 != -1 {
			leftNeighborCoors := coordinate{
				x: coors.x - 1,
				y: coors.y,
			}
			//Only add if the height of the neighbor is at most one higher than the current vertex's height
			if vert.height+1 >= mapGraph.vertices[leftNeighborCoors].height {
				vert.neighbors = append(vert.neighbors, mapGraph.vertices[leftNeighborCoors])
			}
		}
		//If the vertex is not on the right edge of the map, add the vertex to the left to this vertex's list of neighbors
		if coors.x+1 < mapGraph.xDimension {
			rightNeighborCoors := coordinate{
				x: coors.x + 1,
				y: coors.y,
			}
			if vert.height+1 >= mapGraph.vertices[rightNeighborCoors].height {
				vert.neighbors = append(vert.neighbors, mapGraph.vertices[rightNeighborCoors])
			}
		}
		//If the vertex is not on the upper edge of the map, add teh vertex above it to this vertex's list of neighbors
		if coors.y-1 != -1 {
			upNeighborCoors := coordinate{
				x: coors.x,
				y: coors.y - 1,
			}
			if vert.height+1 >= mapGraph.vertices[upNeighborCoors].height {
				vert.neighbors = append(vert.neighbors, mapGraph.vertices[upNeighborCoors])
			}
		}
		//If the vertex is not on the bottom edge of the map, add the vertex below it to this vertex's list of neighbors
		if coors.y+1 < mapGraph.yDimensions {
			bottomNeighborCoors := coordinate{
				x: coors.x,
				y: coors.y + 1,
			}
			if vert.height+1 >= mapGraph.vertices[bottomNeighborCoors].height {
				vert.neighbors = append(vert.neighbors, mapGraph.vertices[bottomNeighborCoors])
			}
		}
	}
}

func main() {
	inputLines := readInput()
	mapGraph := buildGraph(inputLines)
	//part1(mapGraph)
	vert := mapGraph.vertices[coordinate{x: 5, y: 2}]
	fmt.Println(mapGraph, vert)
}
