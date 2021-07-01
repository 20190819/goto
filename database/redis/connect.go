package redis

import (
	"fmt"
	"time"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

var DB *redigo.Pool

type configRedis struct {
	Host        string
	Port        uint32
	Password    string
	MaxIdle     int
	IdleTimeout time.Duration
	MaxActive   int
}

func Conn() {
	conf := configRedis{
		Host:        viper.GetString("redis.host"),
		Port:        viper.GetUint32("redis.port"),
		Password:    viper.GetString("redis.password"),
		MaxIdle:     viper.GetInt("redis.max_idle"),
		IdleTimeout: viper.GetDuration("redis.idle_timeout"),
		MaxActive:   viper.GetInt("redis.max_active"),
	}
	fmt.Println("redis conf >>>", conf)
	DB = &redigo.Pool{
		MaxIdle:     conf.MaxIdle, //空闲数
		IdleTimeout: conf.IdleTimeout * time.Second,
		MaxActive:   conf.MaxActive, //最大数
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", conf.Host)
			if err != nil {
				return nil, err
			}
			if conf.Password != "" {
				if _, err := c.Do("AUTH", conf.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
