package main

import "fmt"

const (
	englishPrefix = "hello"
	chinesePrefix = "你好"
	spanishPrefix = "Hola"
	chineseLanguage = "Chinese"
	spanishLanguage = "Spanish"
)

// Hello ...
func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	prefix := getPrefix(language)
	return fmt.Sprintf("%s, %s", prefix, name)
}

func getPrefix(language string) (prefix string) {
switch language {
case spanishLanguage: {
	prefix = spanishPrefix
}
case chineseLanguage: {
	prefix = chinesePrefix
}
default:
	prefix = englishPrefix
}
return
}
func main() {
	fmt.Println(Hello("You", "cn"))
}
