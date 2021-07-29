package main

import "fmt"

func HelloWorld(number int) {
	if number%3 == 0 && number%5 == 0 {
		fmt.Println("Hello")
	} else if number%3 == 0 {
		fmt.Println("World")
	} else if number%5 == 0 {
		fmt.Println("Hello World")
	}
}

func main() {
	for n := 0; n <= 15; n++ {
		HelloWorld(n)
	}
}
