package main

import "fmt"

func reverseWord(word string) (newWord string) {
	for _, char := range word {
		newWord = string(char) + newWord
	}
	return newWord
}

func main() {
	var kata string
	fmt.Println("Masukkan kata:")
	fmt.Scanf("%s", &kata)
	fmt.Println("Hasil: ")
	fmt.Println(reverseWord(kata))
}
