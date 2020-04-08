package attribution

import (
	"fmt"
	"math/big"
	"strings"
)

// A TransitionMatrix captures the transition probabilities between states in an underlying Markov process.
type TransitionMatrix struct {
	TouchpointIndex         map[Touchpoint]int
	TransitionProbabilities [][]big.Float
}

// GetTransitionProbability returns the transition probability of startNode -> endNode in the underlying Markov process.
func (matrix TransitionMatrix) GetTransitionProbability(startNode Touchpoint, endNode Touchpoint) big.Float {
	return matrix.TransitionProbabilities[matrix.TouchpointIndex[startNode]][matrix.TouchpointIndex[endNode]]
}

// String provides a string representation of TransitionMatrix.
func (matrix TransitionMatrix) String() string {
	represenation := "{\n  TouchpointIndex: {"
	for touchpoint, index := range matrix.TouchpointIndex {
		represenation += fmt.Sprintf("    %s: %d,\n", touchpoint, index)
	}
	represenation += "  },\n  TransitionProbabilities: {"
	for i, row := range matrix.TransitionProbabilities {
		represenation += "\n    "
		for j, _ := range row {
			represenation += matrix.TransitionProbabilities[i][j].Text('f', 2) + ", "
		}
	}
	represenation += "\n  },\n}"
	return represenation
}

// A Touchpoint represents a contributing entity in a ContributionSet.
type Touchpoint struct {
	Name string // name of the touchpoint
}

// Touchpoints represents a list of touchpoints.
// It implements the sort.Interface interface.
type Touchpoints []Touchpoint

// String provides a string represenation of Touchpoints.
func (touchpoints Touchpoints) String() string {
	names := []string{}
	for _, touchpoint := range touchpoints {
		names = append(names, fmt.Sprintf("{%s}", touchpoint.Name))
	}
	return "[" + strings.Join(names, " ") + "]"
}

// Len returns the length of Touchpoints.
func (touchpoints Touchpoints) Len() int {
	return len(touchpoints)
}

// Less provides a strict order on Touchpoints.
func (touchpoints Touchpoints) Less(i, j int) bool {
	return strings.Compare(touchpoints[i].Name, touchpoints[j].Name) == -1
}

// Swap swaps the order of two elements of Touchpoints.
func (touchpoints Touchpoints) Swap(i, j int) {
	touchpoint := touchpoints[i]
	touchpoints[i] = touchpoints[j]
	touchpoints[j] = touchpoint
}

// A Contribution consists of an ordered list of touchpoints together with their combined value.
type Contribution struct {
	Touchpoints Touchpoints
	Value       big.Float
}

func (contribution Contribution) String() string {
	return fmt.Sprintf("{%s %s}", contribution.Touchpoints, contribution.Value.String())
}

func (contribution Contribution) Set() ContributionSet {
	touchpoints := make(map[Touchpoint]struct{})

	for _, touchpoint := range contribution.Touchpoints {
		touchpoints[touchpoint] = struct{}{}
	}

	return ContributionSet{
		Touchpoints: touchpoints,
		Value:       contribution.Value,
	}
}

// A ContributionSet consists of an unordered set of touchpoints together with their combined value.
type ContributionSet struct {
	Touchpoints map[Touchpoint]struct{}
	Value       big.Float
}

func (contribution ContributionSet) String() string {
	return fmt.Sprintf("{%s %s}", contribution.Touchpoints, contribution.Value.String())
}
