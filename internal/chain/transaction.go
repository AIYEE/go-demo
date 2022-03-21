package chain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Transactor struct {
	Client *ethclient.Client
}

func New(rawurl string) (actor *Transactor, err error) {
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		return &Transactor{}, err
	}
	return &Transactor{
		Client: client,
	}, nil
}

func (t *Transactor) GetCallOpts(from common.Address) (*bind.CallOpts, error) {
	return &bind.CallOpts{
		Pending: false,
		From:    from,
	}, nil
}

func (t *Transactor) GetTransactOpts(operatorPrivateKey string) (actor *bind.TransactOpts, err error) {
	privateKey, e := crypto.HexToECDSA(operatorPrivateKey)
	if e != nil {
		return nil, e
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, e := t.Client.PendingNonceAt(context.Background(), fromAddress)
	if e != nil {
		return nil, e
	}
	gasPrice, e := t.Client.SuggestGasPrice(context.Background())
	if e != nil {
		return nil, e
	}

	chainId, err := t.Client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	// transactor := bind.NewKeyedTransactor(privateKey)
	actor, err = bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}
	actor.Nonce = big.NewInt(int64(nonce))
	actor.Value = big.NewInt(0)
	actor.GasLimit = 1297270
	actor.GasPrice = gasPrice
	return actor, nil
}
