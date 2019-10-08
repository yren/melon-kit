package defensive

import (
	"fmt"
	"testing"
)

func TestNormalize(t *testing.T) {
	var pluto *Dog
	fmt.Printf("pluto is %v\n", pluto)
	var nancy Runner = pluto
	fmt.Printf("nancy is %v\n", nancy)
	fmt.Printf("nancy equal nil %v\n", nancy == nil)
	var jerry Runner
	fmt.Printf("jerry is %v\n", jerry)
	fmt.Printf("jerry equal nil %v\n", jerry == nil)

	v, isDog := nancy.(*Dog)
	if isDog {
		if v == nil {
			nancy = nil
		}
	}
	fmt.Printf("nancy equal nil %v\n", nancy == nil)
}
