package main

import "fmt"

// Hello ...
func Hello(name string) string {
	return fmt.Sprintf("hello, %s", name)
}
func main() {
	fmt.Println(Hello("You"))
}
