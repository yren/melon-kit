package feature

import "testing"

func TestGetFeature(t *testing.T) {
	type args struct {
		features map[string]bool
		key Feature
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty map",
			args: args{
				features: map[string]bool{},
				key:      SupportFeatureRun,
			},
			want: false,
		},
		{
			name: "support fly",
			args: args{
				features: map[string]bool{
					string(SupportFeatureFly): true,
				},
				key:      SupportFeatureFly,
			},
			want: true,
		},
		{
			name: "nil map",
			args: args{
				features: nil,
				key: SupportFeatureFly,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFeature(tt.args.features, tt.args.key); got != tt.want {
				t.Errorf("GetFeature() = %v, want %v", got, tt.want)
			}
		})
	}
}
