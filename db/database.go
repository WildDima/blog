package database

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type config struct {
	Host     string
	Port     uint
	User     string
	Password string
	Dbname   string
	Sslmode  string `default:disable`
}

func Connect() (db *gorm.DB, err error) {
	conf := new(config)
	configor.Load(conf, "db/database.yml")
	db, err = gorm.Open(
		"postgres",
		fmt.Sprintf("host=%v port=%v user=%v dbname=%v sslmode=%v password=%v",
			conf.Host,
			conf.Port,
			conf.User,
			conf.Dbname,
			conf.Sslmode,
			conf.Password))
	return
}
