package mysql

func AutoMigrate(list map[string]interface{}) {
	for _, migration := range list {
		DB.AutoMigrate(migration)
	}
}
