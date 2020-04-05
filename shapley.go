package main

import (
	"log"
	"math/big"
)

// GetTotalValue returns the summed value over all contributions.
func GetTotalValue(contributions []ContributionSet) big.Float {
	value := new(big.Float)

	for _, contribution := range contributions {
		value.Add(value, &contribution.Value)
	}

	return *value
}

// GetAllTouchpoints returns a list (without repetition) all touchpoints encountered in contributions.
func GetAllTouchpoints(contributions []ContributionSet) []Touchpoint {
	seen := make(map[Touchpoint]struct{})
	var touchpoints []Touchpoint

	for _, contribution := range contributions {
		for touchpoint, _ := range contribution.Touchpoints {
			if _, found := seen[touchpoint]; !found {
				seen[touchpoint] = struct{}{}
				touchpoints = append(touchpoints, touchpoint)
			}
		}
	}

	return touchpoints
}

// GetCoalitionValue returns the total value a given coalition achieved over a list of contributions.
func GetCoalitionValue(coalition map[Touchpoint]struct{}, allContributions []ContributionSet) big.Float {
	coalitionValue := new(big.Float)

	var coalitionContributed bool
	for _, contribution := range allContributions {
		coalitionContributed = true
		for touchpoint, _ := range contribution.Touchpoints {
			if _, ok := coalition[touchpoint]; !ok {
				coalitionContributed = false
				break
			}
		}
		if coalitionContributed {
			coalitionValue.Add(coalitionValue, &contribution.Value)
		}
	}

	return *coalitionValue
}

// findTouchpoint attempts to find a given touchpoint in a slice of touchpoints.
// If the search is successful, return the first indice where the touchpoint occured in the first coordinate and true in the second coordinate.
// Otherwise, return (-1, false)
func findTouchpoint(touchpoint Touchpoint, slice []Touchpoint) (int, bool) {
	for index, element := range slice {
		if touchpoint == element {
			return index, true
		}
	}
	return -1, false
}

// GetShapleyValue returns the (unordered) Shapley value of a given touchpoint over all provided contributions.
// For a concise introduction to Shapley values, see https://christophm.github.io/interpretable-ml-book/shapley.html
func GetShapleyValue(touchpoint Touchpoint, allContributions []ContributionSet) big.Float {
	shapleyValue := new(big.Float)
	allTouchpoints := GetAllTouchpoints(allContributions)
	touchpointIndex, found := findTouchpoint(touchpoint, allTouchpoints)
	if !found {
		log.Fatal("Illegal touchpoint!")
	}
	allTouchpoints[touchpointIndex] = allTouchpoints[len(allTouchpoints)-1]
	allTouchpoints[len(allTouchpoints)-1] = touchpoint
	powerset := getPowerSetIndices(uint(len(allTouchpoints) - 1))

	for _, subset := range powerset {
		coalition := make(map[Touchpoint]struct{}, len(subset))
		coalitionSize := int64(0)
		for _, index := range subset {
			coalition[allTouchpoints[index]] = struct{}{}
			coalitionSize++
		}

		if _, ok := coalition[touchpoint]; ok {
			log.Fatal("This should never happen!")
		}
		coalitionValue := GetCoalitionValue(coalition, allContributions)

		coalition[touchpoint] = struct{}{}
		addedCoalitionValue := GetCoalitionValue(coalition, allContributions)
		addedCoalitionValue.Sub(&addedCoalitionValue, &coalitionValue)
		nominator := new(big.Int).MulRange(1, coalitionSize)
		nominator.Mul(nominator, new(big.Int).MulRange(1, int64(len(allTouchpoints))-coalitionSize-1))
		denominator := new(big.Int).MulRange(1, int64(len(allTouchpoints)))
		scalingFactor := new(big.Float)
		scalingFactor.Quo(new(big.Float).SetInt(nominator), new(big.Float).SetInt(denominator))

		addedShapleyValue := new(big.Float)
		addedShapleyValue.Mul(scalingFactor, &addedCoalitionValue)
		shapleyValue.Add(shapleyValue, addedShapleyValue)

	}
	return *shapleyValue
}

// getPowerSetIndices provides the powerset of {0, 1, .., size - 1}.
// This can be used to iterate over arbitary powersets by using this result as an index.
func getPowerSetIndices(size uint) [][]uint {
	if size < 1 {
		return [][]uint{[]uint{}}
	}
	powerSetSize := 2 << (size - 1)
	powerset := make([][]uint, 0, powerSetSize)

	index := 0
	for index < powerSetSize {
		var subSet []uint
		for i := uint(0); i < size; i++ {
			if index&(1<<i) > 0 {
				subSet = append(subSet, i)
			}
		}
		powerset = append(powerset, subSet)
		index++
	}
	return powerset
}
