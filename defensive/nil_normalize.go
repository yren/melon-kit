package defensive

import "fmt"

type Runner interface {
	Run()
}

type Dog struct {
	Name string
}

func (d *Dog) Run() {
	fmt.Printf("%s running...\n", d.Name)
}
