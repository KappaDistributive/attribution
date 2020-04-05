package shapley

import (
	"log"
	"math/big"
)

type Touchpoint struct {
	Name string
}

type ContributionSet struct {
	Touchpoints map[Touchpoint]struct{}
	Value       big.Float
}

func GetTotalValue(contributions []ContributionSet) big.Float {
	value := new(big.Float)

	for _, contribution := range contributions {
		value.Add(value, &contribution.Value)
	}

	return *value
}

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

func findTouchpoint(touchpoint Touchpoint, slice []Touchpoint) (int, bool) {
	for index, element := range slice {
		if touchpoint == element {
			return index, true
		}
	}
	return -1, false
}

func GetShapleyValue(touchpoint Touchpoint, allContributions []ContributionSet) big.Float {
	shapleyValue := new(big.Float)
	allTouchpoints := GetAllTouchpoints(allContributions)
	touchpointIndex, found := findTouchpoint(touchpoint, allTouchpoints)
	if !found {
		log.Fatal("Illegal touchpoint: %s", touchpoint)
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
