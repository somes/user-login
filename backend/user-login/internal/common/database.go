package common

import (
	"fmt"
	"server/internal/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func initDBConf() gorm.Dialector {
	var dialector gorm.Dialector

	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	if driverName == "mysql" {
		dialector = mysql.Open(dsn)
	} else if driverName == "sqlite" {
		dialector = sqlite.Open("sqlite.db")
	}
	return dialector
}

func InitDB() *gorm.DB {
	dialector := initDBConf()
	db, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "tb_",
		},
	})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
