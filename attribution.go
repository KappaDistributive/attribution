// Provides multi-channel attribution techniques.
package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Hello, world!")

	t1 := Touchpoint{"touchpoint 1"}
	t2 := Touchpoint{"touchpoint 2"}
	t3 := Touchpoint{"touchpoint 3"}

	c1 := ContributionSet{
		Touchpoints: map[Touchpoint]struct{}{
			t1: struct{}{},
		},
		Value: *big.NewFloat(100.),
	}
	c2 := ContributionSet{
		Touchpoints: map[Touchpoint]struct{}{
			t1: struct{}{},
			t2: struct{}{},
		},
		Value: *big.NewFloat(200.),
	}
	c3 := ContributionSet{
		Touchpoints: map[Touchpoint]struct{}{
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

	contributions := []ContributionSet{c1, c2, c3}

	for _, touchpoint := range []Touchpoint{t1, t2, t3} {
		coalitionValue := GetShapleyValue(touchpoint, contributions)
		fmt.Println(coalitionValue.Int64())
	}
}
