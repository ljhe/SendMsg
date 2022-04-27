package handler

import (
	"sendMsg/libs/httpserver"
	"sendMsg/model"
)

func init() {
	httpserver.RegisterHandler("/", test)
}

type Test struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func test(p *httpserver.HttpParam) (interface{}, error) {
	for i := 0; i < 1000; i++ {
		name := p.Get("name")
		model.GetTestModel().Insert(i, name)
	}
	return nil, nil
}
