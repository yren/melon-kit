package feature

// Feature ...
type Feature string

const (
	SupportFeatureRun Feature = "SupportFeatureRun"
	SupportFeatureFly Feature = "SupportFeatureFly"
)

// GetFeature ...
var GetFeature = func(features map[string]bool, key Feature) bool {
	return features[string(key)]
}