package utils

import (
	"fmt"
	"log"
	"wow_blog/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() (err error) {
	// 连接到mysql数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DbConfig.DBUser,
		config.DbConfig.DBPassword,
		config.DbConfig.DBHost,
		config.DbConfig.DBPort,
		config.DbConfig.DBName)

	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return
	}

	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	return
}

func Close() {
	_ = DB.Close()
}
