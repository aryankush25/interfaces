package main

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type hindiBot struct{}

func main() {
	eb := englishBot{}
	hb := hindiBot{}

	printGreeting(eb)
	printGreeting(hb)
}

func printGreeting(b bot) {
	println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi there!"
}

func (hb hindiBot) getGreeting() string {
	return "Namestey!"
}
