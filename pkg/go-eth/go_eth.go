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

	//768be597fb36eb507098eea93246e0e47ad268f8eeefa46e16117d5da803153e
	// turn that string to a private key
	privateKey, err := crypto.HexToECDSA("768be597fb36eb507098eea93246e0e47ad268f8eeefa46e16117d5da803153e")
	if err != nil {
		return nil, err

	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}

	addr := common.HexToAddress(cnf.SmartContractAddress)

	// auth, err := bind.NewKeyedTransactorWithChainID(strings.NewReader(cnf.Key), "", chainId)
	// if err != nil {
	// 	return nil, err
	// }

	instance, err := goeth.NewCertificate(addr, client)
	if err != nil {
		return nil, err
	}

	//iwanna get the address of transactior
	log.Println("Auth address: ", auth.From.Hex())

	return &GoethClient{
		Client:   client,
		Auth:     auth,
		Instance: instance,
	}, nil
}
