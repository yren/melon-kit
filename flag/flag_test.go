package flag

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestFlag(t *testing.T) {
	flags := int64(0)
	Set(&flags, EnableRun)
	fmt.Printf("flags: %d\n", flags)

	flags = int64(0)
	Set(&flags, EnableSwim)
	fmt.Printf("flags: %d\n", flags)

	flags = int64(0)
	Set(&flags, EnableFly)
	fmt.Printf("flags: %d\n", flags)

	flags = int64(0)
	Set(&flags, Enable_3)
	fmt.Printf("flags: %d\n", flags)

	flags = int64(0)
	Set(&flags, Enable_4)
	fmt.Printf("flags: %d\n", flags)

	flags = int64(0)
	Set(&flags, Enable_5)
	fmt.Printf("flags: %d\n", flags)

	flags = int64(0)
	Set(&flags, Enable_6)
	fmt.Printf("flags: %d\n", flags)

	flags = int64(0)
	Set(&flags, Enable_7)
	fmt.Printf("flags: %d\n", flags)


	flags = int64(0)
	for i := uint(0); i < 64; i++ {
		Set(&flags, i)
		assert.True(t, IsSet(flags, i))
	}

	for i := uint(0); i < 64; i++  {
		Clear(&flags, i)
		assert.False(t, IsSet(flags, i))
	}

	flags = int64(0)
	for i := uint(0); i < 64; i++ {
		Set(&flags, i)
		assert.True(t, IsSet(flags, i))
	}

	for i := uint(0); i < 64; i++  {
		Clear2(&flags, i)
		assert.False(t, IsSet(flags, i))
	}
}

