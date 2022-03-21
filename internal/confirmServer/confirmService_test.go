package confirm

import (
	"io"
	"sync"
	"testing"

	"github.com/AIYEE/go-demo/internal/chain"
	contract "github.com/AIYEE/go-demo/internal/chain/contract/confirm"
	"github.com/AIYEE/go-demo/pkg/logging"
	"github.com/ethereum/go-ethereum/common"
)

func NewServiceMock(t *testing.T) *service {
	chain, err := chain.New("https://data-seed-prebsc-1-s1.binance.org:8545")
	if err != nil {
		t.Error(err)
	}

	confirContract, err := contract.NewConfirm(common.HexToAddress("0xf4768352a50ccfe26490501cd8dc6716af80c843"), chain.Client)
	if err != nil {
		t.Error(err)
	}

	logger := logging.New(io.Discard, 0)
	quit := make(chan struct{}, 1)
	wg := sync.WaitGroup{}
	privateKey := "668db522795b07ba80e178fbb1972f4487d39ebaaadc4bd50416ef3022e61e84"

	s, err := New(confirContract, chain, logger, "./TestMinerNodes.db", quit, &wg, privateKey, true)
	if err != nil {
		t.Error(err)
	}
	return s
}

func TestGetContractState(t *testing.T) {
	s := NewServiceMock(t)
	n := NodeInfo{560, "0x6bDc5fbb2d72b9FE208f7937025F635F851b648c", "775a4f10d676023b603bb36f726cd40437970dc1ab57fa543dc5cbecce97aadd", 0, ""}
	state, err := s.getContractState(n)
	if err != nil {
		t.Error(err)
	}
	t.Log(state)
}

func TestUpdateStateHandler(t *testing.T) {
	s := NewServiceMock(t)

	c, err := s.updateStateHandler()
	if err != nil {
		t.Error(err)
	}

	if c != true {
		t.Log("not complate")
	} else {
		t.Log("complate")
	}
}

func TestConfirmAction(t *testing.T) {
	s := NewServiceMock(t)

	node := NodeInfo{561, "0x8D4Bcaf60Df84FbFb6F68d6EA5D0508d5c7980a1", "775a4f10d676023b603bb36f726cd40437970dc1ab57fa543dc5cbecce97aadd", 0, ""}
	trxHash, err := s.confirmAction(node)
	if err != nil {
		t.Error(err)
	}
	t.Log(trxHash)
}

func TestSendEth(t *testing.T) {
	s := NewServiceMock(t)

	privateKey := "668db522795b07ba80e178fbb1972f4487d39ebaaadc4bd50416ef3022e61e84"

	trxHash, err := s.sendEth(privateKey, "0x8D4Bcaf60Df84FbFb6F68d6EA5D0508d5c7980a1")
	if err != nil {
		t.Error(err)
	}
	t.Log(trxHash)
}
