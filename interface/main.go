package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type hindiBot struct{}

func main() {

	english := englishBot{}
	hindi := hindiBot{}

	printGreeting(english)
	printGreeting(hindi)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//Custom logic for generating an english greeting
	return "Hi there"
}

func (hindiBot) getGreeting() string {
	return "Namaste kaise ho"
}
