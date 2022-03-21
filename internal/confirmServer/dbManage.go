package confirm

import (
	"database/sql"
	"fmt"

	"github.com/AIYEE/go-demo/internal/db"
)

type NodeInfo struct {
	id         int64
	publicKey  string
	privateKey string
	state      int16
	trxHash    string
}

const (
	SQLInsertResult     = `insert into Nodes (publickey, state, trxhash) values (?, ?, ?)`
	SQLQueryNodes       = `select * from Nodes`
	SQLQueryByPublicKey = `select * from Nodes where publickey=?`
	SQLQueryByState     = `select * from Nodes where state=?`
	SQLUpdateState      = `UPDATE Nodes SET state=? WHERE publickey=?`
	SQLUpdateTrxhash    = `update Nodes set trxhash=? where publickey=?`
)

type Manager struct {
	dbOpt db.DbOperator
}

func NewManager(dbFile string) (*Manager, error) {
	operator, err := db.NewSqlite3(dbFile)
	if err != nil {
		return nil, err
	}
	// if err = operator.Create(SQLCreateResultTbl); err != nil {
	// 	return nil, err
	// }

	return &Manager{
		dbOpt: operator,
	}, nil
}

func (m *Manager) GetNodes() ([]NodeInfo, error) {
	result, err := m.dbOpt.Query(SQLQueryNodes)
	if err != nil {
		return nil, err
	}
	nodes := make([]NodeInfo, 0)
	rows := result.((*sql.Rows))
	defer rows.Close()
	for rows.Next() {
		var id int64
		var publicKey string
		var privateKey string
		var state int16
		var trxHash string
		rows.Scan(&id, &publicKey, &privateKey, &state, &trxHash)
		nodes = append(nodes, NodeInfo{id, publicKey, privateKey, state, trxHash})
	}

	return nodes, nil
}

// func (m *Manager) CreateResultTbl() error {
// 	return m.dbOpt.Create(SQL)
// }

func (m *Manager) InsertResult(publicKey string, state int16, txHash string) error {
	return m.dbOpt.Insert(SQLInsertResult, publicKey, state, txHash)
}

func (m *Manager) GetNodeByPublicKey(publicKey_ string) (NodeInfo, error) {
	result, err := m.dbOpt.Query(SQLQueryByPublicKey, publicKey_)
	if err != nil {
		return NodeInfo{}, err
	}

	rows := result.((*sql.Rows))
	defer rows.Close()

	var id int64
	var publicKey string
	var privateKey string
	var state int16
	var trxHash string

	rows.Next()
	if err = rows.Scan(&id, &publicKey, &privateKey, &state, &trxHash); err != nil {
		return NodeInfo{}, err
	}

	return NodeInfo{id, publicKey, privateKey, state, trxHash}, nil
}

func (m *Manager) GetNodesByState(state_ int16) ([]NodeInfo, error) {
	ret, err := m.dbOpt.Query(SQLQueryByState, state_)
	if err != nil {
		return nil, err
	}
	nodes := make([]NodeInfo, 0)
	rows := ret.((*sql.Rows))
	defer rows.Close()

	for rows.Next() {
		var id int64
		var publicKey string
		var privateKey string
		var state int16
		var trxHash string
		rows.Scan(&id, &publicKey, &privateKey, &state, &trxHash)
		nodes = append(nodes, NodeInfo{id, publicKey, privateKey, state, trxHash})
	}

	return nodes, nil
}

func (m *Manager) UpdateState(publicKey_ string, state_ int16) error {
	ret, err := m.dbOpt.Update(SQLUpdateState, state_, publicKey_)
	if !ret {
		return fmt.Errorf("update failed")
	}
	return err
}

func (m *Manager) UpdateTrxHash(publicKey_ string, trxHash_ string) error {
	_, err := m.dbOpt.Update(SQLUpdateTrxhash, trxHash_, publicKey_)
	return err
}
