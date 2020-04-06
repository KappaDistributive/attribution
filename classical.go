package attribution

import (
	"math/big"
)

// GetFirstTouchpointValue returns summed value of all contributions where the given touchpoints happened to be
// first in its list of contributors.
func GetFirstTouchpointValue(touchpoint Touchpoint, allContributions []Contribution) big.Float {
	firstTouchpointValue := new(big.Float)

	for _, contribution := range allContributions {
		length := len(contribution.Touchpoints)
		if length > 0 && touchpoint == contribution.Touchpoints[0] {
			firstTouchpointValue.Add(firstTouchpointValue, &contribution.Value)
		}
	}

	return *firstTouchpointValue
}

// GetLastTouchpointValue returns summed value of all contributions where the given touchpoints happened to be
// last in its list of contributors.
func GetLastTouchpointValue(touchpoint Touchpoint, allContributions []Contribution) big.Float {
	lastTouchpointValue := new(big.Float)

	for _, contribution := range allContributions {

		length := len(contribution.Touchpoints)

		if length > 0 && touchpoint == contribution.Touchpoints[length-1] {
			lastTouchpointValue.Add(lastTouchpointValue, &contribution.Value)
		}
	}

	return *lastTouchpointValue
}
