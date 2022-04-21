package httpserver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"sendMsg/logger"
)

type HttpService struct {
	handler HTTPHandlerFunc
}

type HTTPHandlerFunc func(p interface{}) (interface{}, error)

var HttpServiceObj = getHttpService()
var registerHandler = make(map[string]*HttpService)

func getHttpService() *HttpService {
	return &HttpService{}
}

func (h *HttpService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var content []byte
	r.ParseForm()
	if r.Method == "GET" {
		content = []byte(r.Form.Get("content"))
	} else if r.Method == "POST" {
		content, _ = ioutil.ReadAll(r.Body)
		defer r.Body.Close()
	}
	logger.Debug("这里是测试content:%v", content)
	data, err := h.handler(content)
	if err != nil {
		logger.Err("httpService|ServeHTTP h.handler err:%v", err)
		return
	}
	bytes, err := Marshal(data)
	if err != nil {
		logger.Err("httpService|ServeHTTP Marshal is err:%v", err)
		return
	}
	w.Write(bytes)
}

func (h *HttpService) Init() error {
	for path, handler := range registerHandler {
		Handle(path, handler)
	}
	return Start(":8080")
}

func (h *HttpService) Stop() {
	Stop()
}

func RegisterHandler(path string, handlerFunc HTTPHandlerFunc) {
	registerHandler[path] = &HttpService{
		handler: handlerFunc,
	}
}

func Marshal(data interface{}) ([]byte, error) {
	writeJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return writeJson, nil
}

func Unmarshal(data []byte, paramType reflect.Type) (interface{}, error) {
	param := reflect.New(paramType).Interface()
	err := json.Unmarshal(data, param)
	if err != nil {
		return nil, err
	}
	return param, nil
}
