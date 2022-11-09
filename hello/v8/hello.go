package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func worldLang(name, lang string) string {
	world := "World"
	if name == "" {
		switch lang {

		case spanish:
			world = "mundo"
			return world

		case french:
			world = "monde"
			return world

		default:
			return world
		}
	}

	return name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix

	case french:
		prefix = frenchHelloPrefix

	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

// Hello returns a personalised greeting, considering the language chosen
func Hello(name, lang string) string {
	return greetingPrefix(lang) + worldLang(name, lang)
}

func main() {
	fmt.Println(Hello("world", ""))
}
