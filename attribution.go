package main

import (
	"fmt"
	"github.com/KappaDistributive/attribution/shapley"
	"math/big"
)

func main() {
	fmt.Println("Hello, world!")

	t1 := shapley.Touchpoint{"touchpoint 1"}
	t2 := shapley.Touchpoint{"touchpoint 2"}
	t3 := shapley.Touchpoint{"touchpoint 3"}

	c1 := shapley.ContributionSet{
		Touchpoints: map[shapley.Touchpoint]struct{}{
			t1: struct{}{},
		},
		Value: *big.NewFloat(100.),
	}
	c2 := shapley.ContributionSet{
		Touchpoints: map[shapley.Touchpoint]struct{}{
			t1: struct{}{},
			t2: struct{}{},
		},
		Value: *big.NewFloat(200.),
	}
	c3 := shapley.ContributionSet{
		Touchpoints: map[shapley.Touchpoint]struct{}{
			t1: struct{}{},
			t3: struct{}{},
		},
		Value: *big.NewFloat(300.),
	}
	//c4 := shapley.ContributionSet{
	//	Touchpoints: map[shapley.Touchpoint]struct{}{
	//		t1: struct{}{},
	//		t3: struct{}{},
	//	},
	//	Value: *big.NewFloat(200.),
	//}

	contributions := []shapley.ContributionSet{c1, c2, c3}

	for _, touchpoint := range []shapley.Touchpoint{t1, t2, t3} {
		coalitionValue := shapley.GetShapleyValue(touchpoint, contributions)
		fmt.Println(coalitionValue.Int64())
	}
}
