package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func worldLang(name, lang string) string {
	world := "World"
	if name == "" {
		switch lang {

		case "Spanish":
			world = "mundo"
			return world

		case "French":
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
	case "Spanish":
		prefix = spanishHelloPrefix

	case "French":
		prefix = frenchHelloPrefix

	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

// Hello returns a personalised greeting, defaulting to Hello, world if an empty name is passed.
func Hello(name, lang string) string {
	return greetingPrefix(lang) + worldLang(name, lang)
}

func main() {
	fmt.Println(Hello("world", ""))
}
