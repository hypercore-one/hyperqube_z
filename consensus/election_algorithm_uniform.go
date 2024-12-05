package consensus

import (
	"math/rand"
	"sort"

	"github.com/zenon-network/go-zenon/common/types"
)

type uniformElectionAlgorithm struct {
	group *Context
}

func NewUniformElectionAlgorithm(group *Context) *uniformElectionAlgorithm {
	return &uniformElectionAlgorithm{
		group: group,
	}
}

func (ea *uniformElectionAlgorithm) findSeed(context *AlgorithmConfig) int64 {
	return int64(context.hashH.Height)
}

func (ea *uniformElectionAlgorithm) SelectProducers(context *AlgorithmConfig) []*types.PillarDelegation {
	return ea.uniformRandom(context.delegations, context)
}

func (ea *uniformElectionAlgorithm) uniformRandom(groupA []*types.PillarDelegation, context *AlgorithmConfig) []*types.PillarDelegation {
	var result []*types.PillarDelegation
	total := int(ea.group.NodeCount)
	sort.Sort(types.SortPDByWeight(groupA))

	seed := ea.findSeed(context)
	for len(result) < total {
		random1 := rand.New(rand.NewSource(seed))
		arr := random1.Perm(len(groupA))
		for _, index := range arr {
			result = append(result, groupA[index])
		}
	}
	return result[:total]
}
