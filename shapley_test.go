package attribution

import (
	"fmt"
	"math/big"
	"testing"
)

func touchpointFixture() []Touchpoint {
	var touchpoints []Touchpoint

	for i := 0; i < 10; i++ {
		touchpoints = append(touchpoints, Touchpoint{fmt.Sprintf("Touchpoint %d", i)})
	}

	return touchpoints
}

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

func coalitionFixture() map[Touchpoint]struct{} {
	return map[Touchpoint]struct{}{
		Touchpoint{"Touchpoint 1"}: struct{}{},
		Touchpoint{"Touchpoint 2"}: struct{}{},
	}
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
	contributions := contributionSetFixture()

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

func ExampleGetTotalValue() {
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

	totalValue := GetTotalValue(contributions)

	fmt.Println(totalValue.String())
	// Output: 300
}

func TestGetTotalValue(t *testing.T) {
	contributions := contributionSetFixture()
	totalValue := GetTotalValue(contributions)

	realValue := new(big.Float)
	for _, contribution := range contributions {
		realValue.Add(realValue, &contribution.Value)
	}

	if (*realValue).String() != totalValue.String() {
		t.Errorf("Miscalculated total value.\nExpected: %s\nGot:%s", (*realValue).String(), totalValue.String())
	}
}

func ExampleGetCoalitionValue_singleton() {
	contributions := []ContributionSet{
		ContributionSet{
			Touchpoints: map[Touchpoint]struct{}{
				Touchpoint{"Touchpoint 1"}: struct{}{},
			},
			Value: *new(big.Float).SetFloat64(100.),
		},
		ContributionSet{
			Touchpoints: map[Touchpoint]struct{}{
				Touchpoint{"Touchpoint 1"}: struct{}{},
				Touchpoint{"Touchpoint 2"}: struct{}{},
			},
			Value: *new(big.Float).SetFloat64(200.),
		},
		ContributionSet{
			Touchpoints: map[Touchpoint]struct{}{
				Touchpoint{"Touchpoint 2"}: struct{}{},
				Touchpoint{"Touchpoint 3"}: struct{}{},
			},
			Value: *new(big.Float).SetFloat64(300.),
		},
	}

	coalition := map[Touchpoint]struct{}{
		Touchpoint{"Touchpoint 1"}: struct{}{},
	}

	coalitionValue := GetCoalitionValue(coalition, contributions)
	fmt.Println(coalitionValue.String())
	// Output: 100
}

func ExampleGetCoalitionValue_multiple() {
	contributions := []ContributionSet{
		ContributionSet{
			Touchpoints: map[Touchpoint]struct{}{
				Touchpoint{"Touchpoint 1"}: struct{}{},
			},
			Value: *new(big.Float).SetFloat64(100.),
		},
		ContributionSet{
			Touchpoints: map[Touchpoint]struct{}{
				Touchpoint{"Touchpoint 1"}: struct{}{},
				Touchpoint{"Touchpoint 2"}: struct{}{},
			},
			Value: *new(big.Float).SetFloat64(200.),
		},
		ContributionSet{
			Touchpoints: map[Touchpoint]struct{}{
				Touchpoint{"Touchpoint 1"}: struct{}{},
				Touchpoint{"Touchpoint 3"}: struct{}{},
			},
			Value: *new(big.Float).SetFloat64(300.),
		},
	}

	coalition := map[Touchpoint]struct{}{
		Touchpoint{"Touchpoint 1"}: struct{}{},
		Touchpoint{"Touchpoint 3"}: struct{}{},
	}

	coalitionValue := GetCoalitionValue(coalition, contributions)
	fmt.Println(coalitionValue.String())
	// Output: 400
}

func TestGetCoalitionValue(t *testing.T) {
	contributions := contributionSetFixture()
	coalition := coalitionFixture()

	coalitionValue := GetCoalitionValue(coalition, contributions)
	expectedValue := *new(big.Float).SetFloat64(1400.)

	if coalitionValue.String() != expectedValue.String() {
		t.Errorf("Miscalculated coalition value.\nExpected: %s\nGot: %s", expectedValue.String(), coalitionValue.String())
	}
}

func ExampleGetShapleyValue() {
	contributions := []ContributionSet{
		ContributionSet{
			Touchpoints: map[Touchpoint]struct{}{
				Touchpoint{"Touchpoint 1"}: struct{}{},
			},
			Value: *new(big.Float).SetFloat64(100.),
		},
		ContributionSet{
			Touchpoints: map[Touchpoint]struct{}{
				Touchpoint{"Touchpoint 1"}: struct{}{},
				Touchpoint{"Touchpoint 2"}: struct{}{},
			},
			Value: *new(big.Float).SetFloat64(200.),
		},
		ContributionSet{
			Touchpoints: map[Touchpoint]struct{}{
				Touchpoint{"Touchpoint 1"}: struct{}{},
				Touchpoint{"Touchpoint 3"}: struct{}{},
			},
			Value: *new(big.Float).SetFloat64(300.),
		},
	}
	touchpoint := Touchpoint{"Touchpoint 1"}
	shapleyValue := GetShapleyValue(touchpoint, contributions)

	fmt.Println(shapleyValue.String())
	// Output: 350
}

func TestGetShapleyValue(t *testing.T) {
	contributions := contributionSetFixture()
	touchpoint := touchpointFixture()[2]
	shapleyValue := GetShapleyValue(touchpoint, contributions)
	expectedValue := *(new(big.Float).SetFloat64(585.))

	got, _ := shapleyValue.Float64()
	want, _ := expectedValue.Float64()

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func ExampleSet() {
	contribution := Contribution{
		Touchpoints: Touchpoints([]Touchpoint{
			Touchpoint{"Touchpoint 2"},
			Touchpoint{"Touchpoint 1"},
			Touchpoint{"Touchpoint 3"},
		}),
		Value: *(new(big.Float).SetFloat64(100.)),
	}

	fmt.Println(contribution.Set())
	// Output: {map[{Touchpoint 1}:{} {Touchpoint 2}:{} {Touchpoint 3}:{}] 100}
}

func TestSet(t *testing.T) {
	touchpoints := touchpointFixture()
	contribution := Contribution{
		Touchpoints: touchpoints,
		Value:       *(new(big.Float).SetFloat64(100.)),
	}

	got := contribution.Set()

	touchpointSet := map[Touchpoint]struct{}{}
	for _, touchpoint := range touchpoints {
		touchpointSet[touchpoint] = struct{}{}
	}
	want := ContributionSet{
		Touchpoints: touchpointSet,
		Value:       *(new(big.Float).SetFloat64(100.)),
	}

	if got.Value.String() != want.Value.String() {
		t.Errorf("got %s want %s", got.Value.String(), want.Value.String())
	}
	for touchpoint, _ := range got.Touchpoints {
		if _, ok := want.Touchpoints[touchpoint]; !ok {
			t.Errorf("want is missing %s", touchpoint)
		}
	}
	for touchpoint, _ := range want.Touchpoints {
		if _, ok := got.Touchpoints[touchpoint]; !ok {
			t.Errorf("got is missing %s", touchpoint)
		}
	}
}
