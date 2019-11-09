package dao

import (
	"database/sql"
	log "github.com/Luncert/slog"
	_ "github.com/go-sql-driver/mysql"
)

const MysqlInfo = "root:root@tcp(127.0.0.1:3306)/ma_portal?charset=utf8"

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", MysqlInfo)
	if err != nil {
		log.Fatal(err)
	}
}
