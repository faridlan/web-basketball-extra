package app

import (
	"database/sql"
	"time"

	"github.com/faridlan/web-basketball-extra/helper"
)

func NewDatabase() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/basketball_extra?parseTime=true")
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(60)

	return db
}
