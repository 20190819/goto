package mysql

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func AutoMigrate(list map[string]interface{}) {
	for _, migration := range list {
		fmt.Println("migration>>>",migration)
		if err := DB.AutoMigrate(migration); err != nil {
			logrus.Error(err)
		}
	}
}
