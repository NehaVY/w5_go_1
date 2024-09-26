package main

import "fmt"

func main() {
	Greeting()
	advancedGreeting("Neha")
	advancedGreeting("Valsa")

}

func Greeting() {
	fmt.Println("Hey there!")

}

func advancedGreeting(name string) {
	fmt.Printf(" Hello, %s!\n", name)

}
