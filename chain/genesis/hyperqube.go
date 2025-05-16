package genesis

import (
	"strconv"
	"strings"

	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/vm/constants"
)

func applyHyperQubeConfig(config *GenesisConfig) {

	// HYPERQUBE Z UNIFORM 60
	args := strings.Split(config.ExtraData, " ")
	if args[0] == "HYPERQUBE" {
		types.ZnnTokenStandard = config.TokenConfig.Tokens[0].TokenStandard
		types.QsrTokenStandard = config.TokenConfig.Tokens[1].TokenStandard
		constants.ConsensusConfig.CountingZTS = types.ZnnTokenStandard
		// constants.Decimals is a const, hyperqubes should set their z and q tokens to use same Decimals as mainnet for now

		if args[2] == "UNIFORM" {
			constants.ConsensusConfig.Algorithm = constants.UNIFORM
		}

		// constants.ConsensusConfig.BlockTime and constants.MomentumsPerHour are set separately
		// TODO error handling
		duration, _ := strconv.Atoi(args[3])
		constants.ConsensusConfig.BlockTime = int64(duration)
		updateVmEmbeddedConstants(int64(duration))
	}

}

func updateVmEmbeddedConstants(duration int64) {
	constants.ConsensusConfig.BlockTime = duration

	constants.MomentumsPerHour = 3600 / duration

	constants.MomentumsPerEpoch = constants.MomentumsPerHour * 24
	constants.UpdateMinNumMomentums = uint64(constants.MomentumsPerHour * 5 / 6)
	constants.FuseExpiration = uint64(constants.MomentumsPerHour * 10) // for testnet, 10 hours

	// TODO bridge constants when it is integrated with hyperqube_z

	constants.NetworkZnnRewardConfig = []int64{
		10 * constants.MomentumsPerEpoch / 6 * constants.Decimals,
		6 * constants.MomentumsPerEpoch / 6 * constants.Decimals,
		5 * constants.MomentumsPerEpoch / 6 * constants.Decimals,
		7 * constants.MomentumsPerEpoch / 6 * constants.Decimals,
		5 * constants.MomentumsPerEpoch / 6 * constants.Decimals,
		4 * constants.MomentumsPerEpoch / 6 * constants.Decimals,
		7 * constants.MomentumsPerEpoch / 6 * constants.Decimals,
		4 * constants.MomentumsPerEpoch / 6 * constants.Decimals,
		3 * constants.MomentumsPerEpoch / 6 * constants.Decimals,
		7 * constants.MomentumsPerEpoch / 6 * constants.Decimals,
		3 * constants.MomentumsPerEpoch / 6 * constants.Decimals,
	}

}
