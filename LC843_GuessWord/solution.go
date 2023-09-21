package lc843

import (
	"math/rand"
)

type Master struct {
	secret string
}

func (this *Master) Guess(word string) int {
	match := 0
	for i := range word {
		if word[i] == this.secret[i] {
			match++
		}
	}
	return match
}

// here is my code ...

func findSecretWord(words []string, master *Master) {
	for i := 0; i < 10; i++ {
		randIndex := rand.Intn(len(words))
		candidate := words[randIndex]
		matchLevel := master.Guess(candidate)
		if matchLevel == 6 {
			return
		}
		possibleWords := make([]string, 0)
		for _, word := range words {
			if match(candidate, word) == matchLevel {
				possibleWords = append(possibleWords, word)
			}
		}
		words = possibleWords
	}

}

func match(str1, str2 string) int {
	var matchLevel int
	for i := range str1 {
		if str1[i] == str2[i] {
			matchLevel++
		}
	}
	return matchLevel
}
