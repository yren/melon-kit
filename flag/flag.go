package flag

const (
	// EnableRun ...
	EnableRun = uint(0)

	// EnableSwim ...
	EnableSwim = uint(1)

	// EnableFly ...
	EnableFly = uint(2)

	// Enable3 ...
	Enable_3 = uint(3)

	// Enable_4 ...
	Enable_4 = uint(4)

	// Enable_5 ...
	Enable_5 = uint(5)

	// Enable_6 ...
	Enable_6 = uint(6)

	// Enable_7 ...
	Enable_7 = uint(7)

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

// Clear2 clear the feature flag on given flags
var Clear2 = func(flags *int64, feature uint) {
	*flags = *flags &^ (1 << feature)
}