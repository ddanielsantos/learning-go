package main

import "fmt"

const SaluteInEnglish = "Hello, "
const SaluteInSpanish = "Hola, "
const SaluteInFrench = "Salut, "

func salutePrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = SaluteInSpanish
	case "French":
		prefix = SaluteInFrench
	default:
		prefix = SaluteInEnglish
	}

	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return salutePrefix(language) + name + "!"
}

func main() {
	fmt.Println(Hello("World", ""))
}
