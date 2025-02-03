package database

import (
	"fmt"

	"github.com/Jorozcob/todo-golang/config"
	"github.com/jmoiron/sqlx"
)

func Connect() (*sqlx.DB, error) {
	constr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s sslrootcert=%s",
		config.CFG.DatabaseUser,
		config.CFG.DatabasePassword,
		config.CFG.DatabaseHost,
		config.CFG.DatabasePort,
		config.CFG.DatabaseName,
		config.CFG.DatabaseSslMode,
		config.CFG.Sslrootcert,
	)
	dbConnection, err := sqlx.Connect("postgres", constr)
	if err != nil {
		return nil, err
	}

	err = dbConnection.Ping()
	if err != nil {
		return nil, err
	}
	return dbConnection, err
}
