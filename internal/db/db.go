package db

type Db struct{}

func NewDb() Db {
	return Db{}
}

func (db *Db) Connect() error {
	return nil
}
