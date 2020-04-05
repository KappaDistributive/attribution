package attribution

import (
	"fmt"
	"math/big"
	"testing"
)

func touchpointFixtures() []Touchpoint {
	var touchpoints []Touchpoint

	for i := 0; i < 10; i++ {
		touchpoints = append(touchpoints, Touchpoint{fmt.Sprintf("Touchpoint %d", i)})
	}

	return touchpoints
}

func contributionSetFixtures() []ContributionSet {
	var contributions []ContributionSet

	touchpoints := touchpointFixtures()

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

func ExampleGetAllTouchpoints() {
	contributions := []ContributionSet{
		ContributionSet{
			Touchpoints: map[Touchpoint]struct{}{
				Touchpoint{"Touchpoint 1"}: struct{}{},
				Touchpoint{"Touchpoint 2"}: struct{}{},
			},
			Value: *new(big.Float).SetFloat64(100.),
		},
		ContributionSet{
			Touchpoints: map[Touchpoint]struct{}{
				Touchpoint{"Touchpoint 1"}: struct{}{},
				Touchpoint{"Touchpoint 3"}: struct{}{},
			},
			Value: *new(big.Float).SetFloat64(200.),
		},
	}

	allTouchpoints := GetAllTouchpoints(contributions)
	fmt.Println(allTouchpoints)
	// Output: [{Touchpoint 1} {Touchpoint 2} {Touchpoint 3}]
}

func TestGetAllTouchpoints(t *testing.T) {
	contributions := contributionSetFixtures()

	allTouchpoints := GetAllTouchpoints(contributions)
	for _, contribution := range contributions {
		for touchpoint, _ := range contribution.Touchpoints {
			touchpointFound := false
			for _, candidate := range allTouchpoints {
				if touchpoint == candidate {
					touchpointFound = true
					break
				}
			}
			if !touchpointFound {
				t.Errorf("Couldn't find touchpoint %s", touchpoint)
			}
		}
	}

}
