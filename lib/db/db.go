package db

import (
	"database/sql"
	"fmt"
	"ikbs/lib/config"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var sqlDb *sql.DB

func Init() {
	mysqlConfig := config.LoadConfig().MySQL
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DB)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err.Error())
	}
	var err1 error
	sqlDb, err1 = db.DB()
	if err1 != nil {
		log.Panic(err1.Error())
	}
	sqlDb.SetMaxIdleConns(mysqlConfig.MaxIdleConns)
	sqlDb.SetMaxOpenConns(mysqlConfig.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Hour)

}

func GetSqlDb() *sql.DB {
	return sqlDb
}
