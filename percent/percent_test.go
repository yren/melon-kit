package percent

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestPercent(t *testing.T) {
	threshold := 0.8

	f1 := float64(8)/float64(10)
	fmt.Println(f1)
	assert.Equal(t, f1, threshold)

	f2 := float64(8)/float64(11)
	fmt.Println(f2)
	assert.Equal(t, f2 < threshold, true)

	f3 := float64(9)/float64(11)
	fmt.Println(f3)
	assert.Equal(t, f3 < threshold, false)
}
