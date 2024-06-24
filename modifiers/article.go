package modifiers

import "strings"

func ConvertAToAn(text string) string {
	words := strings.Split(text, " ")
	modifiedWords := make([]string, len(words))

	for i, word := range words {
		if word == "a" && i < len(words)-1 {
			if isVowel(words[i+1]) {
				modifiedWords[i] = "an"
			} else {
				modifiedWords[i] = "a"
			}
		} else {
			modifiedWords[i] = word
		}
	}

	return strings.Join(modifiedWords, " ")
}

func isVowel(word string) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'h'}
	for _, vowel := range vowels {
		if strings.HasPrefix(strings.ToLower(word), string(vowel)) {
			return true
		}
	}
	return false
}
