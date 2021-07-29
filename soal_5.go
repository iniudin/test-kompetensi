package main

import (
	"fmt"
)

func reverseWord(word string) (newWord string) {
	for _, char := range word {
		newWord = string(char) + newWord
	}
	return newWord
}

func isPalindrome(word string) bool {
	tempWord := reverseWord(word)

	if tempWord == word {
		return true
	} else {
		return false
	}

}

func main() {
	var word string
	fmt.Println("Masukkan Kata: ")
	fmt.Scanf("%s", &word)
	print(isPalindrome(word))
}
