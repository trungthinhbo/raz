package quote

import (
	"math/rand/v2"
)

type Quote struct {
	Author string
	Quote  string
}

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (m *Service) GetQuote() Quote {
	quotes := []Quote{
		{Author: "Alice Programmer", Quote: "Good things come to those who write Go"},
		{Author: "Bob Coder", Quote: "Code is poetry; Go is its rhythm"},
		{Author: "Charlie Developer", Quote: "The best code is written with Go, not haste"},
		{Author: "Dana Architect", Quote: "Go boldly where no bugs have gone before"},
		{Author: "Eve Engineer", Quote: "Go fast, Go simple, Go everywhere"},
		{Author: "Frank Hacker", Quote: "In Go we trust, with bugs we fight"},
		{Author: "Grace Programmer", Quote: "Writing Go is like writing dreams that work"},
		{Author: "Henry Developer", Quote: "A clean code is a happy code, especially in Go"},
		{Author: "Ivy Architect", Quote: "Every line of Go is a step towards clarity"},
		{Author: "Jack Engineer", Quote: "Go is simplicity in a complex world"},
		{Author: "Kara Coder", Quote: "The best thing about Go? Its simplicity"},
		{Author: "Liam Developer", Quote: "A bug-free day starts with Go"},
		{Author: "Maya Architect", Quote: "Go: where performance meets elegance"},
		{Author: "Noah Engineer", Quote: "The Go developer's journey is paved with clean code"},
		{Author: "Olivia Programmer", Quote: "The power of Go is in its simplicity"},
		{Author: "Paul Developer", Quote: "Write Go, write the future"},
		{Author: "Quinn Architect", Quote: "In Go, we code with purpose and precision"},
		{Author: "Rachel Hacker", Quote: "Go is not just a language, it's a way of thinking"},
		{Author: "Sam Engineer", Quote: "Good code starts with Go, and ends with clarity"},
		{Author: "Tina Programmer", Quote: "Go is the tool, the developer is the artist"},
		{Author: "Ursula Developer", Quote: "The best way to predict the future is to write Go"},
		{Author: "Victor Coder", Quote: "Every project begins with Go and ends in success"},
		{Author: "Wendy Architect", Quote: "Go builds bridges, not just software"},
		{Author: "Xander Engineer", Quote: "In the world of code, Go is the silent hero"},
		{Author: "Yara Developer", Quote: "A Go developer's mind is always at peace"},
		{Author: "Zane Hacker", Quote: "Go doesn't just solve problems, it defines them"},
		{Author: "Amelia Programmer", Quote: "Code more, worry less, Go more"},
		{Author: "Ben Developer", Quote: "Go is like a compass, always pointing to simplicity"},
		{Author: "Caitlyn Architect", Quote: "In Go, we trust; in bugs, we debug"},
		{Author: "Dylan Engineer", Quote: "Go is the calm in the storm of programming"},
		{Author: "Ella Coder", Quote: "Go is the perfect balance between speed and simplicity"},
		{Author: "Felix Developer", Quote: "To Go is to code with clarity in mind"},
		{Author: "Gina Architect", Quote: "A well-written Go function is worth a thousand lines of code"},
		{Author: "Hugo Engineer", Quote: "Writing Go is like writing the future one function at a time"},
		{Author: "Iris Programmer", Quote: "Go: the language that writes itself"},
		{Author: "Jason Developer", Quote: "Go is the key to unlocking modern software design"},
		{Author: "Kelsey Architect", Quote: "Go is where simplicity meets power"},
		{Author: "Leo Engineer", Quote: "Go lets you write clean code without compromise"},
		{Author: "Mason Coder", Quote: "Simplicity is the soul of Go"},
		{Author: "Nina Developer", Quote: "Go: the art of writing elegant solutions"},
		{Author: "Oscar Architect", Quote: "Go teaches you to think before you code"},
		{Author: "Penny Engineer", Quote: "The future of software is written in Go"},
		{Author: "Quincy Coder", Quote: "Go is not just for today, it’s for tomorrow’s problems"},
		{Author: "Riley Developer", Quote: "Go is for developers who value precision and efficiency"},
		{Author: "Sophie Architect", Quote: "Go is the future; make sure you’re coding it"},
		{Author: "Tyler Engineer", Quote: "Go isn't just fast, it’s thoughtful"},
		{Author: "Uma Programmer", Quote: "In Go, every function is a step toward excellence"},
		{Author: "Vera Developer", Quote: "Go is what happens when simplicity meets productivity"},
		{Author: "Warren Architect", Quote: "Go: because the future can’t wait for bugs"},
		{Author: "Ximena Engineer", Quote: "With Go, you don’t just write code, you craft solutions"},
	}

	return quotes[rand.IntN(len(quotes))]
}
