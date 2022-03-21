package db

type DbOperator interface {
	Open() (err error)
	Create(sqlStr string) error
	Insert(sqlStr string, opt ...interface{}) error
	Query(sqlStr string, opt ...interface{}) (interface{}, error)
	Delete(sqlStr string, opt ...interface{}) (bool, error)
	Update(sqlStr string, opt ...interface{}) (bool, error)
}
