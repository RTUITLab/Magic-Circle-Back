package db

import (
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"

	"github.com/0B1t322/Magic-Circle/ent"
	_ "github.com/go-sql-driver/mysql"
)

func CreateClient(
	dbURI	string,
) (*ent.Client, error) {
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	drv := entsql.OpenDB("mysql", db)

	return ent.NewClient(ent.Driver(drv)), nil
}