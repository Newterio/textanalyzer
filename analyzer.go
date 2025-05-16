// analyzer.go
package analyzer

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// Hello анализирует текст, введённый пользователем
func Hello() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	chars, digits, spaces, vowels, special_chars, letters, sentences, words := 0, 0, 0, 0, 0, 0, 0, 0

	for i := 0; i < len(text); i++ {
		if unicode.IsLetter(rune(text[i])) {
			words++
			for j := i + 1; j < len(text); j++ {
				if !unicode.IsLetter(rune(text[j])) && !unicode.IsDigit(rune(text[j])) {
					i = j
					break
				}
			}
		}
	}
	for _, c := range text {
		if c == '!' || c == '.' || c == '?' {
			sentences++
		}
		if c <= '9' && c >= '0' {
			digits++
		} else if (c <= 'Z' && c >= 'A') || (c <= 'z' && c >= 'a') {
			if c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
				vowels++
			}
			if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
				vowels++
			}
			letters++
		} else if c == ' ' {
			spaces++
		} else {
			special_chars++
		}
		chars++
	}
	if len(text) > 2 && text[len(text)-2] != '!' && text[len(text)-2] != '?' && text[len(text)-2] != '.' {
		sentences++
	}
	fmt.Printf("Characters: %d\nSentences: %d\nWords: %d\nLetters: %d\nVowels: %d\nConsonants: %d\nDigits: %d\nSpaces: %d\nSpecial Characters: %d\n", chars-1, sentences, words, letters, vowels, letters-vowels, digits, spaces, special_chars-1)
}
