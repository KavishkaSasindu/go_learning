package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func openFileHere(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("There is no error while processing")
		fmt.Println("File is opening")
		return
	}
	defer file.Close()
}

func createFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("File is creating")
		fmt.Println("There is no error creating file")
		return
	}
	defer file.Close()
}

func writingFile(filename string, writeData string) {

	// open file
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal("There is an error when opening file")
		fmt.Println("File is not opening")
		return
	}

	defer file.Close()

	// write in file
	file.WriteString(writeData)
}

func readFileHere(filename string) {
	// open file
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("There is an error when opening file")
		fmt.Println(err)
	}

	defer file.Close()

	// read file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("There is an error when reading file %w", err)
	}

}

func main() {
	readFileHere("mytext.txt")
}
