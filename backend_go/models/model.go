package models

import (
	"fmt"
	"github.com/iltyty/db_connect/backend_go/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	// mysql dsn (data source name) format:
	// { username }:{ password }@tcp({ host }:{ port })/{ dbname }?charset=utf8&parseTime=True&loc=Local
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		"qiu", "123456", "localhost", 3306, "db_connect",
	)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	utils.CheckError(err)

	err = db.AutoMigrate(&Test{})
	utils.CheckError(err)
}
