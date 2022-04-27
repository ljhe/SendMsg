package httpserver

import (
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"sendMsg/logger"
	"strconv"
)

const (
	successCode = 0
	failCode    = 1
	inputErr    = -1
)

type HttpService struct {
	handler HTTPHandlerFunc
}

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type HTTPHandlerFunc func(p *HttpParam) (interface{}, error)
type HttpParam struct {
	url.Values
}

var HttpServiceObj = getHttpService()
var registerHandler = make(map[string]*HttpService)

func getHttpService() *HttpService {
	return &HttpService{}
}

func (h *HttpService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var res = Result{}
	var content *HttpParam
	r.ParseForm()
	content = &HttpParam{
		r.Form,
	}
	data, err := h.handler(content)
	if err != nil {
		res.Code = failCode
		res.Msg = err.Error()
		logger.Err("httpService|ServeHTTP h.handler err:%v", err)
	}
	res.Data = data
	bytes, err := Marshal(res)
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
	// TODO 接口待配置
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

func (h *HttpParam) GetInt(v string) int {
	if v == "" {
		return inputErr
	}
	i, err := strconv.Atoi(h.Get(v))
	if err != nil {
		return inputErr
	}
	return i
}
