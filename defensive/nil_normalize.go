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

func Normalize() {
	var ab *Dog
	var jim Runner
	jim = ab
	fmt.Printf("jim equal nil %v", jim == nil)

	v, ok := jim.(*Dog)
	if ok {
		if v == nil {
			jim = nil
		}
		return
	}

}
