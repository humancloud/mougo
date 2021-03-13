package model

import (
	"fmt"
	"log"
	"mougo/utils"
	"time"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func InitDb() {
	var err error

	// connect
	Db, err = gorm.Open(utils.Db,
		fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			utils.DbUser,
			utils.DbPassWord,
			utils.DbHost,
			utils.DbPort,
			utils.DbName))

	if err != nil {
		log.Fatalln("Init Db" + err.Error())
	}

	Db.AutoMigrate(&Article{}, &Category{})
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetConnMaxLifetime(10 * time.Second)
}
