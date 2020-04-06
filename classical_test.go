package attribution

import (
	"fmt"
	"math/big"
	"testing"
)

func ExampleGetFirstTouchpointValue() {
	contributions := []Contribution{
		Contribution{
			Touchpoints: []Touchpoint{
				Touchpoint{"Touchpoint 1"},
			},
			Value: *new(big.Float).SetFloat64(100.),
		},
		Contribution{
			Touchpoints: []Touchpoint{
				Touchpoint{"Touchpoint 1"},
				Touchpoint{"Touchpoint 2"},
			},
			Value: *new(big.Float).SetFloat64(200.),
		},
		Contribution{
			Touchpoints: []Touchpoint{
				Touchpoint{"Touchpoint 1"},
				Touchpoint{"Touchpoint 3"},
				Touchpoint{"Touchpoint 1"},
			},
			Value: *new(big.Float).SetFloat64(300.),
		},
	}
	touchpoint := Touchpoint{"Touchpoint 1"}
	firstTouchpointValue := GetFirstTouchpointValue(touchpoint, contributions)

	fmt.Println(firstTouchpointValue.String())
	// Output: 600
}

func TestGetFirstTouchpointValue(t *testing.T) {
	contributions := contributionFixture()
	touchpoint := touchpointFixture()[2]
	firstTouchpointValue := GetFirstTouchpointValue(touchpoint, contributions)
	expectedValue := *(new(big.Float).SetFloat64(1000.))

	got, _ := firstTouchpointValue.Float64()
	want, _ := expectedValue.Float64()

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func ExampleGetLastTouchpointValue() {
	contributions := []Contribution{
		Contribution{
			Touchpoints: []Touchpoint{
				Touchpoint{"Touchpoint 1"},
			},
			Value: *new(big.Float).SetFloat64(100.),
		},
		Contribution{
			Touchpoints: []Touchpoint{
				Touchpoint{"Touchpoint 1"},
				Touchpoint{"Touchpoint 2"},
			},
			Value: *new(big.Float).SetFloat64(200.),
		},
		Contribution{
			Touchpoints: []Touchpoint{
				Touchpoint{"Touchpoint 1"},
				Touchpoint{"Touchpoint 3"},
				Touchpoint{"Touchpoint 1"},
			},
			Value: *new(big.Float).SetFloat64(300.),
		},
	}
	touchpoint := Touchpoint{"Touchpoint 1"}
	lastTouchpointValue := GetLastTouchpointValue(touchpoint, contributions)

	fmt.Println(lastTouchpointValue.String())
	// Output: 400
}

func TestGetLastTouchpointValue(t *testing.T) {
	contributions := contributionFixture()
	touchpoint := touchpointFixture()[2]
	lastTouchpointValue := GetLastTouchpointValue(touchpoint, contributions)
	expectedValue := *(new(big.Float).SetFloat64(300.))

	got, _ := lastTouchpointValue.Float64()
	want, _ := expectedValue.Float64()

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func ExampleGetLinearValue() {
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
	linearValue := GetLinearValue(touchpoint, contributions)

	fmt.Println(linearValue.String())
	// Output: 350
}

func TestGetLinearValue(t *testing.T) {
	contributions := contributionSetFixture()
	touchpoint := touchpointFixture()[2]
	linearValue := GetLinearValue(touchpoint, contributions)
	expectedValue := *(new(big.Float).SetFloat64(585.))

	got, _ := linearValue.Float64()
	want, _ := expectedValue.Float64()

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func ExampleGetRepeatedLinearValue() {
	contributions := []Contribution{
		Contribution{
			Touchpoints: []Touchpoint{
				Touchpoint{"Touchpoint 1"},
			},
			Value: *new(big.Float).SetFloat64(100.),
		},
		Contribution{
			Touchpoints: []Touchpoint{
				Touchpoint{"Touchpoint 1"},
				Touchpoint{"Touchpoint 2"},
			},
			Value: *new(big.Float).SetFloat64(200.),
		},
		Contribution{
			Touchpoints: []Touchpoint{
				Touchpoint{"Touchpoint 1"},
				Touchpoint{"Touchpoint 3"},
				Touchpoint{"Touchpoint 1"},
			},
			Value: *new(big.Float).SetFloat64(300.),
		},
	}
	touchpoint := Touchpoint{"Touchpoint 1"}
	linearValue := GetRepeatedLinearValue(touchpoint, contributions)

	fmt.Println(linearValue.String())
	// Output: 400
}

func TestRepeatedGetLinearValue(t *testing.T) {
	contributions := contributionFixture()
	touchpoint := touchpointFixture()[2]
	linearValue := GetRepeatedLinearValue(touchpoint, contributions)
	expectedValue := *(new(big.Float).SetFloat64(585.))

	got, _ := linearValue.Float64()
	want, _ := expectedValue.Float64()

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
