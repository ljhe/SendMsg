package handler

import (
	"sendMsg/libs/httpserver"
)

func init() {
	httpserver.RegisterHandler("/", test)
}

type Test struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func test(p interface{}) (interface{}, error) {
	res := p.(*Test)
	test := &Test{
		res.Id,
		res.Name,
	}
	return test, nil
}
