package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile() {
	file, err := os.Open("/src/nand2tetris/projects/06/add/Add.asm")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Read the data of the file
	fileReader := bufio.NewReader(file)

	if err != nil {
		log.Fatal(err)
	}

	for {
		line, err := fileReader.ReadString('\n')

		if err != nil {
			break
		}

		// Remove lines that contain comments or only new lines (two bytes)
		if strings.HasPrefix(line, "//") || len(line) == 2 {
			continue
		}

		fmt.Printf(line)
	}
}

func main() {
	readFile()
}
