package defensive

import "fmt"

// Runner ...
type Runner interface {
	Run()
}

// Dog ...
type Dog struct {
	Name string
}

// Run ...
func (d *Dog) Run() {
	fmt.Printf("%s running...\n", d.Name)
}
