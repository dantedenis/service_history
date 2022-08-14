package contract

import "database/sql"

type IRequester interface {
	Start(db *sql.DB) error
}
