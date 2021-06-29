package mysql

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type configDB struct {
	Host           string
	Port           uint32
	Username       string
	Password       string
	Database       string
	Charset        string
	MaxConnect     int
	MaxIdleConnect int
	MaxLifeSeconds int
}

var configdb configDB
var dns string

func init() {
	configdb = configDB{
		Host:           viper.GetString("mysql.host"),
		Port:           viper.GetUint32("mysql.port"),
		Username:       viper.GetString("mysql.username"),
		Password:       viper.GetString("mysql.password"),
		Database:       viper.GetString("mysql.database"),
		Charset:        viper.GetString("mysql.charset"),
		MaxConnect:     viper.GetInt("mysql.max_connect"),
		MaxIdleConnect: viper.GetInt("mysql.max_idle_connect"),
		MaxLifeSeconds: viper.GetInt("mysql.max_life_seconds"),
	}
	dns = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		configdb.Username, configdb.Password, configdb.Host, configdb.Port, configdb.Database, configdb.Charset, true, "Local")
}

func Conn() {
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		// Logger: //,
	})

	sqlDB, err := DB.DB()
	sqlDB.SetMaxOpenConns(configdb.MaxConnect)
	sqlDB.SetConnMaxIdleTime(time.Duration(configdb.MaxIdleConnect))
	sqlDB.SetConnMaxLifetime(time.Duration(configdb.MaxLifeSeconds))

	if err != nil {
		panic(err)
	}
}
