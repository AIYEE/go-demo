package confirm

import "testing"

const DBFILE = "TestMinerNodes.db"

func TestInsertResult(t *testing.T) {
	manager, err := NewManager(DBFILE)
	if err != nil {
		return
	}

	manager.InsertResult("123", 1, "123trx")
}

func TestUpdateResultState(t *testing.T) {
	manager, err := NewManager(DBFILE)
	if err != nil {
		t.Error(err)
	}
	err = manager.UpdateState("123", 2)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateResultTrxHash(t *testing.T) {
	manager, err := NewManager(DBFILE)
	if err != nil {
		t.Error(err)
	}
	err = manager.UpdateTrxHash("123", "testhash")
	if err != nil {
		t.Error(err)
	}
}

func TestGetResult(t *testing.T) {
	manager, err := NewManager(DBFILE)
	if err != nil {
		t.Error(err)
	}

	result, err := manager.GetNodeByPublicKey("123")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestGetResultByState(t *testing.T) {
	manager, err := NewManager(DBFILE)
	if err != nil {
		t.Error(err)
	}

	result, err := manager.GetNodesByState(Init)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
