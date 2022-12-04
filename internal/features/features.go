package features

import (
	"k8s.io/component-base/featuregate"
)

const (
// TestGate featuregate.Feature = "TestGate"
)

var defaultFeatureGates = map[featuregate.Feature]featuregate.FeatureSpec{
	//TestGate: {Default: false, PreRelease: featuregate.Alpha},
}

func init() {
	if err := DefaultMutableFeatureGate.Add(defaultFeatureGates); err != nil {
		panic(err)
	}
}

func SetDefaultFeatureGates() {
}
