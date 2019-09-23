package flag

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestFlag(t *testing.T) {
	flags := int64(0)
	Set(&flags, EnableRun)
	assert.True(t, IsSet(flags, EnableRun))
	assert.False(t, IsSet(flags, EnableSwim))
	fmt.Printf("flag: %d\n", flags)
	Set(&flags, EnableSwim)
	fmt.Printf("flag: %d\n", flags)
	Set(&flags, EnableFly)
	fmt.Printf("flag: %d\n", flags)

	Clear(&flags, EnableRun)
	Clear(&flags, EnableSwim)
	Clear(&flags, EnableFly)

	for i := uint(0); i < 64; i++ {
		Set(&flags, i)
		assert.True(t, IsSet(flags, i))
	}

	for i := uint(0); i < 64; i++  {
		Clear(&flags, i)
		assert.False(t, IsSet(flags, i))
	}
}

