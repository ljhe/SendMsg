package db

import "sendMsg/log"

func InitDbModel() error {
	for _, conf := range dbConfig {
		dataSourceName := getDataSourceName(conf.dataBaseName, conf.userName, conf.passWord, conf.host, conf.port)
		err := initDb(dataSourceName)
		if err != nil {
			log.Err("dbModel|InitDbModel initDb is err:%v dataSourceName:%v", err, dataSourceName)
			return err
		}
		log.Info("db:%v init success", conf.dataBaseName)
	}
	return nil
}
