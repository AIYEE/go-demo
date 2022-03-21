package confirm

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/AIYEE/go-demo/internal/chain"
	contract "github.com/AIYEE/go-demo/internal/chain/contract/confirm"
	"github.com/AIYEE/go-demo/pkg/logging"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

const Default int16 = -1
const (
	Init int16 = iota
	Confirmed
	Settled
	Claimed
	NoProfit
	Bottom
)

type Options struct {
	EthFromAddress string
}

type service struct {
	confirmContract *contract.Confirm
	transactor      *chain.Transactor
	logger          logging.Logger
	dbManager       *Manager
	ethFromAddress  string
	needGas         bool
	wg              *sync.WaitGroup
	quit            chan struct{}
	done            chan struct{}
}

func New(contract *contract.Confirm, actor *chain.Transactor, logger logging.Logger, dbFile string, quit chan struct{}, wg *sync.WaitGroup, ethFromAddress string, needGas bool) (s *service, err error) {
	manager, err := NewManager(dbFile)
	if err != nil {
		return nil, err
	}

	return &service{
		confirmContract: contract,
		transactor:      actor,
		logger:          logger,
		dbManager:       manager,
		ethFromAddress:  ethFromAddress,
		needGas:         needGas,
		wg:              wg,
		quit:            quit,
		done:            make(chan struct{}),
	}, nil
}

func (s *service) Start() {
	defer s.wg.Done()
	timer := time.NewTimer(0)
	defer timer.Stop()

	// transfer gas
	if s.needGas {
		if err := s.sendEthToNees(s.ethFromAddress); err != nil {
			s.logger.Warning(err)
		}
	}

	timer.Reset(1 * time.Second)
	for {
		select {
		case <-s.quit:
			s.logger.Info("quit confirm server.")
			return
		case <-s.done:
			s.logger.Info("finish confirm work.")
			return
		case <-timer.C:
			// update
			updateFinish, err := s.updateStateHandler()
			if err != nil {
				s.logger.Errorf("update state failed. Error: %v", err)
				timer.Reset(10 * time.Minute)
				break
			}

			if !updateFinish {
				timer.Reset(0)
				break
			}

			// handleConfirm
			s.handleConfirm()
			timer.Reset(0)
		}

	}

}

func (s *service) getContractState(n NodeInfo) (int16, error) {
	publicKey := n.publicKey
	// privateKey := n.privateKey

	opts, err := s.transactor.GetCallOpts(common.HexToAddress(publicKey))
	if err != nil {
		return -1, err
	}

	ret, err := s.confirmContract.ProfitInfo(opts, common.HexToAddress(publicKey))
	if err != nil {
		return -1, err
	}

	return int16(ret.State), nil
}

func (s *service) confirmAction(n NodeInfo) (string, error) {
	opts, err := s.transactor.GetTransactOpts(n.privateKey)
	if err != nil {
		return "", err
	}

	ret, err := s.confirmContract.ConfirmProfits(opts)
	if err != nil {
		return "", err
	}

	return ret.Hash().String(), nil
}

func (s *service) handleConfirm() {
	nodes, err := s.dbManager.GetNodesByState(Init)
	if err != nil {
		s.logger.Errorf("get nodes fail.Error: %v", err)
		time.Sleep(3 * time.Second)
	}

	if len(nodes) == 0 {
		s.done <- struct{}{}
		s.logger.Info("finish confirm")
		return
	}

	for _, n := range nodes {
		time.Sleep(100 * time.Millisecond)
		needConfirm, err := s.needConfirm(n.publicKey)
		if err != nil {
			s.logger.Error(err)
			continue
		}

		if needConfirm {
			trxHash, err := s.confirmAction(n)
			if err != nil {
				s.logger.Errorf("confirm action excute failed. From: %v, Error: %v", n.publicKey, err)
				continue
			}
			err = s.dbManager.UpdateTrxHash(n.publicKey, trxHash)
			if err != nil {
				s.logger.Errorf("update confirm hash failed. Error: %v, From: %v, TrxHash: %v", err, n.publicKey, trxHash)
			}
			s.logger.Infof("confirm ID: %v Address: %v success. Trx: %v", n.id, n.publicKey, trxHash)
		}
	}
}

func (s *service) needConfirm(publicKey string) (bool, error) {
	n, err := s.dbManager.GetNodeByPublicKey(publicKey)
	if err != nil {
		return false, fmt.Errorf("get node(%v) from db fail. Error:%v", publicKey, err)
	}

	if n.state <= Init {
		state, err := s.getContractState(n)
		if err != nil {
			return false, fmt.Errorf("get state(%v) from db fail. Error:%v", publicKey, err)
		}

		if state >= Confirmed {
			return false, nil
		}

		return true, nil
	} else {
		return false, nil
	}
}

func (s *service) updateStateHandler() (bool, error) {
	s.logger.Debug("update db")
	complate := true
	nodes, err := s.dbManager.GetNodes()
	if err != nil {
		return false, fmt.Errorf("get nodes from db fail. Error:%v", err)
	}

	for _, n := range nodes {
		if n.state == Claimed && n.trxHash != "" {
			continue
		}

		state, err := s.getContractState(n)
		if err != nil {
			s.logger.Warningf("get state fail.Error: %v. addr: %v", err, n.publicKey)
			complate = false
			continue
		}
		if state != n.state {
			err = s.dbManager.UpdateState(n.publicKey, state)
			if err != nil {
				s.logger.Warningf("update state fail.Error: %v. addr: %v", err, n.publicKey)
				complate = false
			}
			if n.trxHash != "" {
				err = s.dbManager.UpdateTrxHash(n.publicKey, n.trxHash)
				if err != nil {
					s.logger.Warningf("update state fail.Error: %v. addr: %v", err, n.publicKey)
				}
			}
			s.logger.Debugf("update db ID: %v Address: %v State: %v", n.id, n.publicKey, state)
		}
	}

	return complate, nil
}

func (s *service) sendEth(fromPrivateKey string, to string) (string, error) {
	privateKey, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)
	nonce, err := s.transactor.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	value := big.NewInt(10000000000000000) //0.01 eth
	gasLimit := uint64(21000)
	gasPrice, err := s.transactor.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	toAddress := common.HexToAddress(to)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := s.transactor.Client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	err = s.transactor.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}

func (s *service) Shutdown(ctx context.Context) error {
	return nil
}

func (s *service) sendEthToNees(signer string) error {
	nodes, err := s.dbManager.GetNodesByState(Init)
	if err != nil {
		return fmt.Errorf("get nodes from db fail. Error:%v", err)
	}

	for _, n := range nodes {
		need, err := s.needConfirm(n.publicKey)
		if err != nil {
			return err
		}

		if need {
			trx, err := s.sendEth(signer, n.publicKey)
			if err != nil {
				s.logger.Debugf("send BNB to %v failed. Error: %v", n.publicKey, err)
			} else {
				s.logger.Infof("send BNB to %v success. Trx: %v", n.publicKey, trx)
			}
		}
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}
