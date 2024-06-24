package modifiers

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ConvertHexAndBinToDecimal(text string) string {
	words := strings.Split(text, " ")

	for i := 1; i < len(words); i++ {
		switch words[i] {
		case "(hex)":
			value, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				fmt.Printf("Failed to convert %s: %v\n", words[i-1], err)
				os.Exit(1)
			}
			words[i-1] = strconv.FormatInt(value, 10)
			words = append(words[:i], words[i+1:]...)
			i-- // Decrement i to re-check the current index
		case "(bin)":
			value, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				fmt.Printf("Failed to convert %s: %v\n", words[i-1], err)
				os.Exit(1)
			}
			words[i-1] = strconv.FormatInt(value, 10)
			words = append(words[:i], words[i+1:]...)
			i-- // Decrement i to re-check the current index
		}
	}

	return strings.Join(words, " ")
}
