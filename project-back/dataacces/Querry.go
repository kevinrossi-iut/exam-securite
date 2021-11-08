package dataacces

import "database/sql"

type Query interface {
	GetResult() *sql.Rows
}
