package goeth

import (
	"context"
	"strings"

	"github.com/ansxy/nagabelajar-be-go/config"
	goeth "github.com/ansxy/nagabelajar-be-go/pkg/go-eth/artifact"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type GoethClient struct {
	Client   *ethclient.Client
	Auth     *bind.TransactOpts
	Instance *goeth.Certificate
}

func NewGoethClient(cnf config.SmartContractConfig) (*GoethClient, error) {
	client, err := ethclient.Dial(cnf.Dial)
	if err != nil {
		return nil, err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	addr := common.HexToAddress(cnf.SmartContractAddress)

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(cnf.Key), "", chainId)
	if err != nil {
		return nil, err
	}

	instance, err := goeth.NewCertificate(addr, client)
	if err != nil {
		return nil, err
	}

	return &GoethClient{
		Client:   client,
		Auth:     auth,
		Instance: instance,
	}, nil
}
