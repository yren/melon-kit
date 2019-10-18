package defaulttype

import (
	"fmt"
	"testing"
)

func TestDefaultType(t *testing.T) {
	numint := 10
	fmt.Printf("type is %T\n", numint)

	numf := 10.1
	fmt.Printf("type is %T\n", numf)
}
