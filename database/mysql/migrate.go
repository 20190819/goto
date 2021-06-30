package mysql

import "github.com/sirupsen/logrus"

func AutoMigrate(list map[string]interface{}) {
	for _, migration := range list {
		if err := DB.AutoMigrate(migration); err != nil {
			logrus.Error(err)
		}
	}
}
