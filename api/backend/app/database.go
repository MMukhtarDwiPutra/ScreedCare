package app

import(
	"ScreedCare/backend/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB{
	dsn := "screedca_putra:Mukhtar1600@tcp(127.0.0.1:3306)/screedca_screedcare?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/screedcare")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	return db
}