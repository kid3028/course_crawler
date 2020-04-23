package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

var DB *gorm.DB

/**
   初始化数据库
 */
func InitDB()  {
	DB = openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.database"))

}

/**
 打开数据库连接
 */
func openDB(username, password, addr, database string) *gorm.DB  {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		database,
		true,
		"Local")
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "database connect fail")
	}
	setupDB(db)
	return db

}

func setupDB(db *gorm.DB)  {
	db.LogMode(viper.GetBool("gormlog"))
}

/**
   关闭数据库连接
 */
func CloseDB()  {
	DB.Close()
}
