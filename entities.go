package attribution

import (
	"math/big"
)

// A Touchpoint represents a contributing entity in a ContributionSet.
type Touchpoint struct {
	Name string // name of the touchpoint
}

// A ContributionSet consists of a set of touchpoints together with their combined value.
type ContributionSet struct {
	Touchpoints map[Touchpoint]struct{}
	Value       big.Float
}
