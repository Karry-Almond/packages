package favoriteDB

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	sqlDB, err := sql.Open("mysql", "root:root@tcp(47.100.224.26:3306)/tiktok8062?parseTime=True")
	Db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}
