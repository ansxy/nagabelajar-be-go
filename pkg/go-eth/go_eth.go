package goeth

import (
	"context"
	"log"

	"github.com/ansxy/nagabelajar-be-go/config"
	goeth "github.com/ansxy/nagabelajar-be-go/pkg/go-eth/artifact"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

	privateKey, err := crypto.HexToECDSA("768be597fb36eb507098eea93246e0e47ad268f8eeefa46e16117d5da803153e")
	if err != nil {
		return nil, err

	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}

	addr := common.HexToAddress(cnf.SmartContractAddress)

	instance, err := goeth.NewCertificate(addr, client)
	if err != nil {
		return nil, err
	}

	log.Println("Auth address: ", auth.From.Hex())

	return &GoethClient{
		Client:   client,
		Auth:     auth,
		Instance: instance,
	}, nil
}
