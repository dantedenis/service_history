package contract

import "database/sql"

type IProvider interface {
	Open() error
	GetConn() *sql.DB
	Close() error
}
