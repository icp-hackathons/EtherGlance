package networks

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/network"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/filecoin-project/venus/pkg/config"
	"github.com/filecoin-project/venus/venus-shared/types"
)

func Mainnet() *NetworkConf {
	return &NetworkConf{
		Bootstrap: config.BootstrapConfig{
			Addresses: []string{
				"/dns4/bootstrap-0.mainnet.filops.net/tcp/1347/p2p/12D3KooWCVe8MmsEMes2FzgTpt9fXtmCY7wrq91GRiaC8PHSCCBj",
				"/dns4/bootstrap-1.mainnet.filops.net/tcp/1347/p2p/12D3KooWCwevHg1yLCvktf2nvLu7L9894mcrJR4MsBCcm4syShVc",
				"/dns4/bootstrap-2.mainnet.filops.net/tcp/1347/p2p/12D3KooWEWVwHGn2yR36gKLozmb4YjDJGerotAPGxmdWZx2nxMC4",
				"/dns4/bootstrap-3.mainnet.filops.net/tcp/1347/p2p/12D3KooWKhgq8c7NQ9iGjbyK7v7phXvG6492HQfiDaGHLHLQjk7R",
				"/dns4/bootstrap-4.mainnet.filops.net/tcp/1347/p2p/12D3KooWL6PsFNPhYftrJzGgF5U18hFoaVhfGk7xwzD8yVrHJ3Uc",
				"/dns4/bootstrap-5.mainnet.filops.net/tcp/1347/p2p/12D3KooWLFynvDQiUpXoHroV1YxKHhPJgysQGH2k3ZGwtWzR4dFH",
				"/dns4/bootstrap-6.mainnet.filops.net/tcp/1347/p2p/12D3KooWP5MwCiqdMETF9ub1P3MbCvQCcfconnYHbWg6sUJcDRQQ",
				"/dns4/bootstrap-7.mainnet.filops.net/tcp/1347/p2p/12D3KooWRs3aY1p3juFjPy8gPN95PEQChm2QKGUCAdcDCC4EBMKf",
				"/dns4/bootstrap-8.mainnet.filops.net/tcp/1347/p2p/12D3KooWScFR7385LTyR4zU1bYdzSiiAb5rnNABfVahPvVSzyTkR",
				"/dns4/lotus-bootstrap.forceup.cn/tcp/41778/p2p/12D3KooWFQsv3nRMUevZNWWsY1Wu6NUzUbawnWU5NcRhgKuJA37C",
				"/dns4/bootstrap-0.starpool.in/tcp/12757/p2p/12D3KooWGHpBMeZbestVEWkfdnC9u7p6uFHXL1n7m1ZBqsEmiUzz",
				"/dns4/bootstrap-1.starpool.in/tcp/12757/p2p/12D3KooWQZrGH1PxSNZPum99M1zNvjNFM33d1AAu5DcvdHptuU7u",
				"/dns4/node.glif.io/tcp/1235/p2p/12D3KooWBF8cpp65hp2u9LK5mh19x67ftAam84z9LsfaquTDSBpt",
				"/dns4/bootstrap-0.ipfsmain.cn/tcp/34721/p2p/12D3KooWQnwEGNqcM2nAcPtRR9rAX8Hrg4k9kJLCHoTR5chJfz6d",
				"/dns4/bootstrap-1.ipfsmain.cn/tcp/34723/p2p/12D3KooWMKxMkD5DMpSWsW7dBddKxKT7L2GgbNuckz9otxvkvByP",
				"/dns4/bootstarp-0.1475.io/tcp/61256/p2p/12D3KooWQjaNmbz9b1XmheQB3RWsRjKSzuRLfjeiDZHyX7Y5RcBr",
			},
			MinPeerThreshold: 1,
			Period:           "30s",
		},
		Network: config.NetworkParamsConfig{
			DevNet:                false,
			NetworkType:           types.NetworkMainnet,
			GenesisNetworkVersion: network.Version0,
			ReplaceProofTypes: []abi.RegisteredSealProof{
				abi.RegisteredSealProof_StackedDrg32GiBV1,
				abi.RegisteredSealProof_StackedDrg64GiBV1,
			},
			BlockDelay:              30,
			ConsensusMinerMinPower:  10 << 40,
			MinVerifiedDealSize:     1 << 20,
			PreCommitChallengeDelay: abi.ChainEpoch(150),
			ForkUpgradeParam: &config.ForkUpgradeConfig{
				UpgradeBreezeHeight:   41280,
				UpgradeSmokeHeight:    51000,
				UpgradeIgnitionHeight: 94000,
				UpgradeRefuelHeight:   130800,
				UpgradeAssemblyHeight: 138720,
				UpgradeTapeHeight:     140760,
				UpgradeLiftoffHeight:  148888,
				// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.
				// Miners, clients, developers, custodians all need time to prepare.
				// We still have upgrades and state changes to do, but can happen after signaling timing here.
				UpgradeKumquatHeight:    170000,
				UpgradeCalicoHeight:     265200,
				UpgradePersianHeight:    265200 + (builtin2.EpochsInHour * 60),
				UpgradeOrangeHeight:     336458,
				UpgradeTrustHeight:      550321, // 2021-03-04T00:00:30Z
				UpgradeNorwegianHeight:  665280, // 2021-04-12T22:00:00Z
				UpgradeTurboHeight:      712320, // 2021-04-29T06:00:00Z
				UpgradeHyperdriveHeight: 892800, // 2021-06-30T22:00:00Z
				UpgradeChocolateHeight:  1231620,
				UpgradeOhSnapHeight:     1594680, // 2022-03-01T15:00:00Z
				UpgradeSkyrHeight:       1960320, // 2022-07-06T14:00:00Z

				BreezeGasTampingDuration: 120,
				UpgradeClausHeight:       343200, // 2020-12-22T02:00:00Z
			},
			DrandSchedule:  map[abi.ChainEpoch]config.DrandEnum{0: 5, 51000: 1},
			AddressNetwork: address.Mainnet,
		},
	}
}