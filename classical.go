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

// GetLinearValue returns the linear value (ignoring repetition) of a given touchpoint summed over all contributions.
// The linear value without repititions for Contribution objecs can best be calculated by first transformating them
// to ContributionSet objects with the Set() method and then applying this function.
func GetLinearValue(touchpoint Touchpoint, allContributions []ContributionSet) big.Float {
	linearValue := new(big.Float)

	for _, contribution := range allContributions {
		// check if touchpoint was part of this contribution
		for candidate, _ := range contribution.Touchpoints {
			if touchpoint == candidate {
				numberTouchpoints := float64(len(contribution.Touchpoints))
				addedValue := contribution.Value
				// distribute value equally among all contributors
				addedValue.Quo(&addedValue, new(big.Float).SetFloat64(numberTouchpoints))
				linearValue.Add(linearValue, &addedValue)
				break
			}
		}
	}

	return *linearValue
}

// GetRepeatedLinearValue returns the linear value (with repition) of a given touchpoint summed over all contributions.
func GetRepeatedLinearValue(touchpoint Touchpoint, allContributions []Contribution) big.Float {
	linearValue := new(big.Float)

	for _, contribution := range allContributions {
		touchpointContributions := 0
		// check if touchpoint was part of this contribution
		for _, candidate := range contribution.Touchpoints {
			if touchpoint == candidate {
				touchpointContributions++
			}
		}
		if touchpointContributions > 0 {
			numberTouchpoints := float64(len(contribution.Touchpoints))
			addedValue := contribution.Value
			// distribute value equally among all contributors according to their number of contributions
			addedValue.Mul(&addedValue, new(big.Float).SetFloat64(float64(touchpointContributions)))
			addedValue.Quo(&addedValue, new(big.Float).SetFloat64(numberTouchpoints))
			linearValue.Add(linearValue, &addedValue)
		}

	}

	return *linearValue
}
