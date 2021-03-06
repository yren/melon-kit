package defensive

import (
	"fmt"
	"testing"
)

func TestNormalize(t *testing.T) {
	var pluto *Dog
	fmt.Printf("pluto type %T\n", pluto)
	var nancy Runner = pluto
	fmt.Printf("nancy type %T\n", nancy)
	var jim Runner
	fmt.Printf("jim type %T\n", jim)
	fmt.Printf("nancy value %v\n", nancy)
	fmt.Printf("nancy equal nil %v\n", nancy == nil)
	var jerry Runner
	fmt.Printf("jerry type %v\n", jerry)
	fmt.Printf("jerry equal nil %v\n", jerry == nil)

	// normalize nancy from type nil to normal nil
	v, isDog := nancy.(*Dog)
	if isDog {
		if v == nil {
			nancy = nil
		}
	}
	fmt.Printf("nancy equal nil %v\n", nancy == nil)
}
