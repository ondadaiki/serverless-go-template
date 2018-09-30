package database

type Scanner interface {
	Scan(src ...interface{}) error
}
