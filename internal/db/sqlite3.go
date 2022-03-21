package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type NodeInfo struct {
	Id         int
	PublicKey  string
	PrivateKey string
}

const DRIVER = "sqlite3"

type Sqlite3 struct {
	driver string
	name   string
	db     *sql.DB
}

func NewSqlite3(name string) (*Sqlite3, error) {
	db, err := sql.Open(DRIVER, name)
	if err != nil {
		return nil, fmt.Errorf("open failed. Error: %v", err)
	}
	return &Sqlite3{
		driver: DRIVER,
		name:   name,
		db:     db,
	}, nil
}

func (n *Sqlite3) Open() (err error) {
	db, err := sql.Open(n.driver, n.name)
	if err != nil {
		return fmt.Errorf("open failed. Error: %v", err)
	}
	n.db = db
	return nil
}

func (n *Sqlite3) Create(sqlStr string) error {
	if n.db == nil {
		if err := n.Open(); err != nil {
			return err
		}
	}

	_, err := n.db.Exec(sqlStr)
	return err
}

func (n *Sqlite3) exec(sqlStr string, opt ...interface{}) (sql.Result, error) {
	if n.db == nil {
		return nil, fmt.Errorf("no db")
	}

	stmt, err := n.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(opt...)
	return result, err
}

func (n *Sqlite3) Insert(sqlStr string, opt ...interface{}) error {
	_, err := n.exec(sqlStr, opt...)
	return err
}

func (n *Sqlite3) query(sqlStr string, opt ...interface{}) (*sql.Rows, error) {
	if n.db == nil {
		return nil, fmt.Errorf("no db")
	}

	stmt, err := n.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(opt...)
	if err != nil {
		return nil, err
	}
	// defer rows.Close()

	return rows, nil
}

func (n *Sqlite3) Query(sqlStr string, opt ...interface{}) (interface{}, error) {
	return n.query(sqlStr, opt...)
}

func (n *Sqlite3) Delete(sqlStr string, opt ...interface{}) (bool, error) {
	result, err := n.exec(sqlStr, opt...)
	if err != nil {
		return false, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (n *Sqlite3) Update(sqlStr string, opt ...interface{}) (bool, error) {
	result, err := n.exec(sqlStr, opt...)
	if err != nil {
		return false, err
	}
	nums, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if nums == 0 {
		return false, nil
	}

	return true, nil
}
