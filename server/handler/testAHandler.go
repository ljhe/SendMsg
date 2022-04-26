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
	id := p.GetInt("id")
	name := p.Get("name")
	model.GetTestModel().Insert(id, name)
	return nil, nil
}