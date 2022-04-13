package db

import (
	"github.com/go-gorp/gorp"
	"sendMsg/logger"
)

type modelMapStruct struct {
	model    Model
	initFunc func(dbMap *gorp.DbMap)
}

type Model interface {
	SetDbMap(dbMap *gorp.DbMap)
	GetDbMap() *gorp.DbMap
}

type CommonModel struct {
	dbMap *gorp.DbMap
}

var modelMap = make(map[string][]modelMapStruct)

func Register(dbKey string, model Model, initFunc func(dbMap *gorp.DbMap)) {
	m, ok := modelMap[dbKey]
	if ok {
		modelMap[dbKey] = append(m, modelMapStruct{model: model, initFunc: initFunc})
		return
	}
	modelMap[dbKey] = make([]modelMapStruct, 0)
	modelMap[dbKey] = append(modelMap[dbKey], modelMapStruct{model: model, initFunc: initFunc})
}

func InitDbModel() error {
	for _, conf := range dbConfig {
		dataSourceName := getDataSourceName(conf.dataBaseName, conf.userName, conf.passWord, conf.host, conf.port)
		db, err := initDb(dataSourceName)
		if err != nil {
			logger.Err("dbModel|InitDbModel initDb is err:%v dataSourceName:%v", err, dataSourceName)
			return err
		}
		logger.Info("db:%v init success", conf.dataBaseName)

		dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
		//dbMap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
		err = dbMap.CreateTablesIfNotExists()
		if err != nil {
			logger.Err("dbModel|InitDbModel createTablesIfNotExists is err:%v", err)
			return err
		}

		if maps, ok := modelMap[conf.dataBaseName]; ok {
			for _, m := range maps {
				m.model.SetDbMap(dbMap)
				m.initFunc(dbMap)
			}
		}
	}
	return nil
}

func (c *CommonModel) SetDbMap(dbMap *gorp.DbMap) {
	c.dbMap = dbMap
}

func (c *CommonModel) GetDbMap() *gorp.DbMap {
	return c.dbMap
}
