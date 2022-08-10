package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBconn       *gorm.DB
	DATABASE_URI string = "mSql:mSql2021@tcp(127.0.0.1:3306)/m_db?charset=utf8mb4&parseTime=True&loc=Local"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	DBconn = db

	return db
}
