package main

import "fmt"

// Hello ...
func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("hello, %s", name)
}
func main() {
	fmt.Println(Hello("You"))
}
