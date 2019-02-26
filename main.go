package main

import "fmt"

// every different type
// now has a custom type, in this case is called bot
// if you are a type in this program with a function called getGreeting()
// and return a string then you are now an honorary member of type "bot"
//
// Now that a specific type is an honorary member of type "bot", you can now call printGreeting()
type bot interface {
	// anything that matches this interface, has access to type bot
	getGreeting( /*what we expect this function to be called with string etc */ ) string // ( I can specify what the function should return within these parenthesis)
	// ^ function name

}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// imagine this to have custom logic that cant be reused
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	// imagine this to have custom logic that cant be reused
	return "Hola!"
}
