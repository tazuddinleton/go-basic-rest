package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tazuddinleton/go/basicrest/conf"
	"github.com/tazuddinleton/go/basicrest/logger"
)

type RestDb struct {
	db *sql.DB
}

func New(dbConf conf.Db) *RestDb {
	s := dbConf.ConnString()
	newDb, err := sql.Open("mysql", s)
	if err != nil {
		logger.Log.
			Error().
			Msg("Database connection failed " + err.Error())
	}

	err = newDb.Ping()
	if err != nil {
		logger.Log.
			Error().
			Msg("Database connection failed " + err.Error())
	}

	return &RestDb{db: newDb}
}

func (d RestDb) Close() {
	d.db.Close()
}
