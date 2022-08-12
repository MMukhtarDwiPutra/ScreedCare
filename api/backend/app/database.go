package app

import(
	"ScreedCare/backend/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB{
	dsn := "root:@tcp(127.0.0.1:3306)/screedcare?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/screedcare")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	return db
}