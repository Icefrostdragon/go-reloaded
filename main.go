package main

import (
	"fmt"
	"go-reloaded/modifiers"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) != 2 {
		if len(arguments) == 0 {
			fmt.Println("No input or output files provided")
		} else {
			fmt.Println("Too many files provided")
		}
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		os.Exit(1)
	}

	inputFile := arguments[0]
	outputFile := arguments[1]

	input, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Failed to read input file: %v\n", err)
		os.Exit(1)
	}

	text := string(input)

	modifiedText := modifiers.ModifyText(text)

	err = ioutil.WriteFile(outputFile, []byte(modifiedText), 0644)
	if err != nil {
		fmt.Printf("Failed to write to output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Modifications completed successfully!")
}
