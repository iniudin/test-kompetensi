package main

import (
	"fmt"
	"regexp"
)

func validateEmail(email string) {
	var pattern = `^([\w-\.]){0,20}@([\w-]+\.)+(?:co.id|id)$`
	match, _ := regexp.MatchString(pattern, email)
	if match {
		fmt.Printf("Email %s valid\n", email)
	} else {
		fmt.Printf("Email %s tidak valid\n", email)
	}
}

func main() {
	var email string
	fmt.Println("Masukkan Email: ")
	fmt.Scanf("%s", &email)
	validateEmail(email)
}
