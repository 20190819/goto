package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB
var err error
// https://810.workarea7.live/view_video.php?viewkey=db1e35d0ba414e7a2fff&page=1&viewtype=basic&category=top
//  https://810.workarea7.live/view_video.php?viewkey=415c13e7f2f8940b47f5&page=1&viewtype=basic&category=mr
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

func Conn() {
	configdb=configDB{
		Host:           viper.GetString("MYSQL_HOST"),
		Port:           viper.GetUint32("MYSQL_PORT"),
		Username:       viper.GetString("MYSQL_USERNAME"),
		Password:       viper.GetString("MYSQL_PASSWORD"),
		Database:       viper.GetString("DATABASE"),
		Charset:        viper.GetString("CHARSET"),
		MaxConnect:     viper.GetInt("MAX_CONNECT"),
		MaxIdleConnect: viper.GetInt("MAX_IDLE_CONNECT"),
		MaxLifeSeconds: viper.GetInt("MAX_LIFE_SECONDS"),
	}
	dns = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		configdb.Username, configdb.Password, configdb.Host, configdb.Port, configdb.Database, configdb.Charset, true, "Local")
	fmt.Println("mysql conf>>>",dns)
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		// Logger: //,
	})

	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(configdb.MaxConnect)
	sqlDB.SetConnMaxIdleTime(time.Duration(configdb.MaxIdleConnect))
	sqlDB.SetConnMaxLifetime(time.Duration(configdb.MaxLifeSeconds))
}
