package attribution

import (
	"fmt"
	"math/big"
)

// touchpointFixture provides a fixture for a slice of Touchpoint objects.
func touchpointFixture() []Touchpoint {
	var touchpoints []Touchpoint

	for i := 0; i < 10; i++ {
		touchpoints = append(touchpoints, Touchpoint{fmt.Sprintf("Touchpoint %d", i)})
	}

	return touchpoints
}

// contributionFixture provides a fixture for a slice of Contribution objects.
func contributionFixture() []Contribution {
	var contributions []Contribution

	touchpoints := touchpointFixture()

	for i := 0; i < 5; i++ {
		for j := 0; j <= 5; j++ {
			touchpointList := Touchpoints(touchpoints[i : i+j])
			contribution := Contribution{
				Touchpoints: touchpointList,
				Value:       *new(big.Float).SetFloat64(float64(100 * i)),
			}
			contributions = append(contributions, contribution)
		}
	}

	return contributions
}

// contributionSetFixture provides a fixture for a slice of ContributionSet objects.
func contributionSetFixture() []ContributionSet {
	var contributions []ContributionSet

	touchpoints := touchpointFixture()

	for i := 0; i < 5; i++ {
		for j := 0; j <= 5; j++ {
			touchpointMap := make(map[Touchpoint]struct{})
			for _, touchpoint := range touchpoints[i : i+j] {
				touchpointMap[touchpoint] = struct{}{}
			}
			contribution := ContributionSet{
				Touchpoints: touchpointMap,
				Value:       *new(big.Float).SetFloat64(float64(100 * i)),
			}
			contributions = append(contributions, contribution)
		}
	}

	return contributions
}

// coalitionFixture provides a coaltion fixture.
func coalitionFixture() map[Touchpoint]struct{} {
	return map[Touchpoint]struct{}{
		Touchpoint{"Touchpoint 1"}: struct{}{},
		Touchpoint{"Touchpoint 2"}: struct{}{},
	}
}
