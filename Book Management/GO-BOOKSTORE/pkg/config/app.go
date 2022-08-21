package config

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	ConectionString := "sqlserver://:@127.0.0.1:1433?database=BookManagement"
	database,_:= gorm.Open(sqlserver.Open(ConectionString), &gorm.Config{})
	db = database
}
func GetDB() *gorm.DB {
	
	return db
}
