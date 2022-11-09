package main

import (
	"fmt"
	"strings"
)

// Hello returns a personalised greeting.
// func Hello() string { <old>
func Hello(names ...string) string {
	// return "Hello world" <old>
	return "Hello, " + strings.Join(names, ", ")
}

func main() {
	// fmt.Println(Hello()) <old>
	fmt.Println(Hello("world"))
}
