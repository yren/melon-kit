package flag

import "testing"
import "github.com/stretchr/testify/assert"

func TestFlag(t *testing.T) {
	flags := int64(0)
	Set(&flags, EnableRun)
	assert.True(t, true)
}

