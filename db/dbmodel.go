package db

func InitDbModel() error {
	dataSourceName := getDataSourceName(dataBaseName)
	err := initDb(dataSourceName)
	return err
}
