package model

import (
	"github.com/go-gorp/gorp"
	"sendMsg/db"
	"sendMsg/logger"
)

type Test1 struct {
	Id   int    `db:"id, primarykey, autoincrement"`
	Name string `db:"name"`
}

type TestModel struct {
	db.CommonModel
}

var testModel = &TestModel{}

func init() {
	db.Register(db.DataBaseName[0], testModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(Test1{}, "test1").SetKeys(true, "id")
	})
}

func GetTestModel() *TestModel {
	return testModel
}

func (t *TestModel) Insert() error {
	data := &Test1{Name: "test2"}
	err := t.GetDbMap().Insert(data)
	if err != nil {
		logger.Err("db_test1|Insert err:%v", err)
		return err
	}
	return nil
}
