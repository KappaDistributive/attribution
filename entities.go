package attribution

import (
	"math/big"
	"strings"
)

// A Touchpoint represents a contributing entity in a ContributionSet.
type Touchpoint struct {
	Name string // name of the touchpoint
}

// Touchpoints represents a list of touchpoints.
// It implements the sort.Interface interface.
type Touchpoints []Touchpoint

func (touchpoints Touchpoints) Len() int {
	return len(touchpoints)
}

func (touchpoints Touchpoints) Less(i, j int) bool {
	return strings.Compare(touchpoints[i].Name, touchpoints[j].Name) == -1
}

func (touchpoints Touchpoints) Swap(i, j int) {
	touchpoint := touchpoints[i]
	touchpoints[i] = touchpoints[j]
	touchpoints[j] = touchpoint
}

// A Contribution consists of an ordered list of touchpoints together with their combined value.
type Contribution struct {
	Touchpoints []Touchpoint
	Value       big.Float
}

func (contribution Contribution) Set() ContributionSet {
	touchpoints := make(map[Touchpoint]struct{})

	for _, touchpoint := range contribution.Touchpoints {
		touchpoints[touchpoint] = struct{}{}
	}

	return ContributionSet{
		Touchpoints: touchpoints,
		Value:       contribution.Value,
	}
}

// A ContributionSet consists of an unordered set of touchpoints together with their combined value.
type ContributionSet struct {
	Touchpoints map[Touchpoint]struct{}
	Value       big.Float
}
