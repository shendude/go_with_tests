package hello_world

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return getPrefix(language) + name
}

func getPrefix(language string) (retval string) {
	switch language {
	case french:
		retval = frenchHelloPrefix
	case spanish:
		retval = spanishHelloPrefix
	default:
		retval = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
