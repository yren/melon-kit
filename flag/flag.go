package flag

const (
	// EnableRun ...
	EnableRun = uint(0)

	// EnableSwim ...
	EnableSwim = uint(1)

	// EnableFly
	EnableFly = uint(2)

)


// Set set the feature flag on given flags
var Set = func(flags *int64, feature uint) {
	*flags |= 1 << feature
}

// IsSet check whether given feature is set on flags
var IsSet = func(flags int64, feature uint) bool {
	return flags & (1 << feature) != 0
}

// Clear clear the feature flag on given flags
var Clear = func(flags *int64, feature uint) {
	*flags &= ^(1 << feature)
}