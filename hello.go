package main

import "fmt"

// godoc -http :8000
// If you go to localhost:8000/pkg you will see all the packages installed on your system

// If you don't have godoc command, then maybe you are using the newer version of Go (1.14 or later) which is no longer including godoc.
// You can manually install it with go get golang.org/x/tools/cmd/godoc
const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}


// The function name starts with a lowercase letter.
// In Go public functions start with a capital letter and private ones start with a lowercase.

// In our function signature we have made a named return value (prefix string).
// This will create a variable called prefix in your function.
// It will be assigned the "zero" value. This depends on the type, for example ints are 0 and for strings it is "".
// You can return whatever it's set to by just calling return rather than return prefix.
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Ege", ""))
}
