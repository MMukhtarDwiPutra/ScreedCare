package app

import(
	"database/sql"
	"ScreedCare/backend/helper"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB{
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/screedcare")
	helper.PanicIfError(err)

	return db
}