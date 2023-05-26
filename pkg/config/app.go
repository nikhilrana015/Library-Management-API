package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {

	dialect := EnvConfigs.DbUsername + ":" + EnvConfigs.DbPassword + "@(localhost:" + EnvConfigs.DbServerPort + ")/" + EnvConfigs.DbName
	d, err := gorm.Open("mysql", dialect+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Error: ", err)
		panic("Failed to connect with the database")
	}

	fmt.Println("Connected with database successfully")
	db = d
}

func GetDB() *gorm.DB {
	return db
}
