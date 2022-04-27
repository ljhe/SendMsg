package model

import (
	"github.com/go-gorp/gorp"
	"sendMsg/libs/db"
	"sendMsg/libs/logger"
)

type Test1 struct {
	Id   int    `db:"id, primarykey, autoincrement"`
	Name string `db:"name"`
	Age  int    `db:"age"`
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

func (t *TestModel) Insert(id int, name string) error {
	data := &Test1{Name: "test2", Age: 18}
	err := t.GetDbMap().Insert(data)
	if err != nil {
		logger.Err("db_test1|Insert err:%v", err)
		return err
	}
	return nil
}

func (t *TestModel) Query() error {
	data := &Test1{}
	res, err := t.GetDbMap().Get(data, 1)
	if err != nil {
		logger.Err("db_test1|Query err:%v", err)
		return err
	}
	logger.Info("db_test1|Query res:%v", res)
	return nil
}
