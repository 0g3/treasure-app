package db

import (
	"github.com/0g3/treasure-app/app/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var Conn *gorm.DB

func newConn() (*gorm.DB, error) {
	return gorm.Open("mysql",
		config.Env().DBUser+":"+config.Env().DBPassword+
			"@tcp("+config.Env().DBAddress+":"+config.Env().DBPort+")/"+
			config.Env().DBName+"?charset=utf8mb4&parseTime=True&loc=Local")
}

func init() {
	var err error
	Conn, err = newConn()
	if err != nil {
		log.Fatal("mysqlのオープンに失敗", err)
	}
}
