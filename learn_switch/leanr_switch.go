package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]
	greet := "greetings"

	switch exprl := len(word); word {
	 case "hello":
		fmt.Println("Hi yourself")
		fallthrough
	 case "farewell":
	 case greet:
		 fmt.Println("Salude")
	 case "goodbye","bye":
		fmt.Println("Bye bye")
	 default:
		 fmt.Println("I dont understand, but it was", l, "chars long")
	}

	c := "crackerjack"
	switch l := len(word); {
		case word == "hello":
			fmt.Println("Hi yourself")
		case l == 1:
			fmt.Println("One letter word")
		case 1 < l && l < 10, word == c:
			fmt.Println("This word is either", c, "or it is 2-9 chars long")
	}
	if word == "hello" {
		fmt.Println("Hi yourself")
	} else if word == "goodbye" {
		fmt.Println("Bye bye")
	} else if word == "greetings" {
		fmt.Println("Salude")
	} else {
		fmt.Println("I dont understand")
	}
}