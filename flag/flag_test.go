package flag

import "testing"

func TestFlag(t *testing.T) {
	flags := int64(0)

	Set(&flags, EnableRun)
}

