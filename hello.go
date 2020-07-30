package main

import "fmt"

const englishHelloPrefix = "Hello, "
const delimiter = " "

func Hello(first_name, last_name string) string {

	if first_name == "" || last_name == "" {
		first_name = "World"
		last_name = "People"
	}

	return englishHelloPrefix + first_name + delimiter + last_name
}

func main() {
	fmt.Println(Hello("Ige", "Adetokunbo"))
}
