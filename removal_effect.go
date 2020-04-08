package attribution

import (
	"math/big"
)

func GetTransitionMatrix(allContributions []Contribution) TransitionMatrix {
	// get touchpoints
	contributionSets := make([]ContributionSet, len(allContributions))
	for index, contribution := range allContributions {
		contributionSets[index] = contribution.Set()
	}
	touchpoints := GetAllTouchpoints(contributionSets)

	// create TouchpointIndex
	TouchpointIndex := make(map[Touchpoint]int)
	for index, touchpoint := range touchpoints {
		TouchpointIndex[touchpoint] = index
	}

	// count transitions between touchpoints
	size := touchpoints.Len()
	TransitionProbabilities := make([][]big.Float, size)
	for i := 0; i < size; i++ {
		TransitionProbabilities[i] = make([]big.Float, size)
	}
	for _, contribution := range allContributions {
		length := contribution.Touchpoints.Len()

		for i := 0; i < length-1; i++ {
			for j := i + 1; j < length; j++ {
				startNode := contribution.Touchpoints[i]
				endNode := contribution.Touchpoints[j]
				// add one to the transition startNode -> endNode
				TransitionProbabilities[TouchpointIndex[startNode]][TouchpointIndex[endNode]].Add(
					&TransitionProbabilities[TouchpointIndex[startNode]][TouchpointIndex[endNode]],
					new(big.Float).SetFloat64(1.))
			}
		}
	}

	// transform transition counts into probabilities
	for i := 0; i < size; i++ {
		rowSum := new(big.Float)
		for j := 0; j < size; j++ {
			rowSum.Add(rowSum, &TransitionProbabilities[i][j])
		}
		// check if rowSum > 0
		if rowSum.Sign() == 1 {
			for j := 0; j < size; j++ {
				TransitionProbabilities[i][j].Quo(&TransitionProbabilities[i][j], rowSum)
			}
		}
	}
	return TransitionMatrix{
		TouchpointIndex:         TouchpointIndex,
		TransitionProbabilities: TransitionProbabilities,
	}
}
