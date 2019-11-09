package dao

import (
	"database/sql"
	log "github.com/Luncert/slog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ma-portal/core"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", core.MysqlInfo)
	if err != nil {
		log.Fatal(err)
	}
}
