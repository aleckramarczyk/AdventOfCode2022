package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var Args struct {
	InputFile *string
}

func readData() string {
	file, err := os.Open(*Args.InputFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

type datastreamBuffer struct {
	packetBuffer  [4]rune
	messageBuffer [14]rune
}

func (b *datastreamBuffer) shiftPacketBuffer(newChar rune) {
	var newBuffer [4]rune
	newBuffer[3] = newChar
	copy(newBuffer[:3], b.packetBuffer[1:])
	b.packetBuffer = newBuffer
}

func (b *datastreamBuffer) shiftMessageBuffer(newChar rune) {
	var newBuffer [14]rune
	newBuffer[13] = newChar
	copy(newBuffer[:13], b.messageBuffer[1:])
	b.messageBuffer = newBuffer
}

func (b *datastreamBuffer) initializePacketBuffer(startingBuffer string) error {
	if len(startingBuffer) != 4 {
		return nil //TODO add error type
	} else {
		var newBuffer [4]rune
		for i, char := range startingBuffer {
			newBuffer[i] = char
		}
		b.packetBuffer = newBuffer
		return nil
	}
}

func (b *datastreamBuffer) initializeMessageBuffer(startingBuffer string) error {
	if len(startingBuffer) != 4 {
		return nil
	} else {
		var newBuffer [14]rune
		for i, char := range startingBuffer {
			newBuffer[i] = char
		}
		b.messageBuffer = newBuffer
		return nil
	}
}

func packetBufferContainsAllUniqueCharacters(buffer [4]rune) bool {
	visited := make(map[rune]bool, 0)
	for i := 0; i < len(buffer); i++ {
		if visited[buffer[i]] {
			return false
		} else {
			visited[buffer[i]] = true
		}
	}
	return true
}

func messageBufferContainsAllUniqueCharacters(buffer [14]rune) bool {
	visited := make(map[rune]bool, 0)
	for i := 0; i < len(buffer); i++ {
		if visited[buffer[i]] {
			return false
		} else {
			visited[buffer[i]] = true
		}
	}
	return true
}

func processDatastream(datastream string) (startOfPacketIndex int, startOfMessageIndex int) {
	buffer := new(datastreamBuffer)
	buffer.initializePacketBuffer(datastream[:4])
	for index, char := range datastream[4:] {
		buffer.shiftPacketBuffer(char)
		if packetBufferContainsAllUniqueCharacters(buffer.packetBuffer) {
			startOfPacketIndex = index + 5
			break
		}
	}

	buffer.initializeMessageBuffer(datastream[startOfPacketIndex:])
	for index, char := range datastream[startOfPacketIndex:] {
		buffer.shiftMessageBuffer(char)
		if messageBufferContainsAllUniqueCharacters(buffer.messageBuffer) {
			startOfMessageIndex = index + startOfPacketIndex + 1
			break
		}
	}
	return startOfPacketIndex, startOfMessageIndex
}

func main() {
	Args.InputFile = flag.String("path", "input", "file path to read from")
	flag.Parse()

	datastream := readData()
	startOfPacketMarkerIndex, startOfMessageIndex := processDatastream(datastream)
	fmt.Println(startOfPacketMarkerIndex)
	fmt.Println(startOfMessageIndex)
}
