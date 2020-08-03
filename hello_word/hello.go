package main

import (
	"fmt"
	"strings"
)

const delimiter = " "
const englishHelloPrefix = "Hello, "

func Hello(first_name, last_name, language string) string {

	if first_name == "" || last_name == "" {
		first_name = "World"
		last_name = "People"
	}

	return getGreetingsPrefix(language) + first_name + delimiter + last_name
}

func getGreetingsPrefix(language string) (prefix string) {
	language = strings.ToLower(language)

	switch language {
	case "french":
		prefix = "Bonjour, "
	case "spanish":
		prefix = "Hola, "
	case "hindi":
		prefix = "Hailo, "
	default:
		prefix = "Hello, "
	}

	return prefix
}

func main() {
	fmt.Println(Hello("Ige", "Adetokunbo", "French"))
}
