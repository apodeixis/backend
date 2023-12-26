package config

import (
	"crypto/ecdsa"
	"fmt"
	"net/url"

	"github.com/apodeixis/backend/internal/types"
	"github.com/apodeixis/backend/pkg/posts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var figuredConfigs map[string]*EvmChainConfig

type EvmChainConfig struct {
	Client           *ethclient.Client
	Contract         *posts.Posts
	SenderPrivateKey *ecdsa.PrivateKey
}

type evmChainConfig struct {
	ClientURL        *url.URL `figure:"client_url,required"`
	ContractAddress  string   `figure:"contract_address,required"`
	SenderPrivateKey string   `figure:"sender_private_key,required"`
}

func (c *config) EvmChainConfig(chain types.EVMChain) *EvmChainConfig {
	config, figured := figuredConfigs[string(chain)]
	if !figured {
		auxConfig := new(evmChainConfig)
		err := figure.
			Out(auxConfig).
			From(kv.MustGetStringMap(c.getter, string(chain))).
			Please()
		if err != nil {
			panic(errors.Wrap(err, fmt.Sprintf("failed to figure out %s config", chain)))
		}
		client, err := ethclient.Dial(auxConfig.ClientURL.String())
		if err != nil {
			panic(errors.Wrap(err, fmt.Sprintf("failed to dial: %s", auxConfig.ClientURL.String())))
		}
		if !common.IsHexAddress(auxConfig.ContractAddress) {
			panic(errors.Wrap(err, fmt.Sprintf("contract address is not a valid hex address: %s", auxConfig.ContractAddress)))
		}
		contractAddress := common.HexToAddress(auxConfig.ContractAddress)
		postsContract, err := posts.NewPosts(contractAddress, client)
		if err != nil {
			panic(errors.Wrap(err, "failed to create an instance contract"))
		}
		senderPrivateKey, err := crypto.HexToECDSA(auxConfig.SenderPrivateKey)
		if err != nil {
			panic(errors.Wrap(err, fmt.Sprintf("failed to parse sender private key from hex to ecdsa in %s", chain)))
		}
		config = &EvmChainConfig{
			Client:           client,
			Contract:         postsContract,
			SenderPrivateKey: senderPrivateKey,
		}
	}
	return config
}
