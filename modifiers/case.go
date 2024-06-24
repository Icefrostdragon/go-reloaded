package modifiers

import (
	"strconv"
	"strings"
)

func ConvertUpLowCap(text string) string {
	words := strings.Split(text, " ")

	for i, word := range words {
		if word == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if word == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if word == "(cap)" {
			words[i-1] = strings.Title(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if word == "(up," {
			b := strings.Trim(string(words[i+1]), words[i+1][1:])
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			// lower with number
		} else if word == "(low," {
			b := strings.Trim(string(words[i+1]), words[i+1][1:])
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number; j++ {
				words[i-j] = strings.ToLower(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			// capitalize with num
		} else if word == "(cap," {
			b := strings.Trim(string(words[i+1]), words[i+1][1:])
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number; j++ {
				words[i-j] = strings.Title(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		}
	}
	return strings.Join(words, " ")

}
