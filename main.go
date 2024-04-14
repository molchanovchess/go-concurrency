package main

import (
	"fmt"
	"time"
)

func main() {
	greet("Nice to meet you!")
	greet("How are you?")
	slowGreet("How ... are ... you ... ?")
	greet("I hope OK!")
}

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
}
