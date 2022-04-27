package httpserver

import (
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"sendMsg/libs/logger"
	"sync"
)

type HttpServer struct {
	listener net.Listener
	closing  bool
	Router   *mux.Router
	wg       sync.WaitGroup
}

func NewHttpServer() *HttpServer {
	server := &HttpServer{Router: mux.NewRouter()}
	server.Router.StrictSlash(true)
	return server
}

func (h *HttpServer) Start(netAddr string) error {
	l, err := net.Listen("tcp", netAddr)
	if err != nil {
		return err
	}
	h.listener = l

	h.wg.Add(1)
	go func() {
		h.serve()
		h.wg.Done()
	}()
	return nil
}

func (h *HttpServer) Stop() {
	h.closing = true
	if h.listener != nil {
		h.listener.Close()
	}
	h.wg.Wait()
}

func (h *HttpServer) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	h.Router.HandleFunc(pattern, handler)
}

func (h *HttpServer) Handle(pattern string, handler http.Handler) {
	h.Router.Handle(pattern, handler)
}

func (h *HttpServer) serve() {
	err := http.Serve(h.listener, h.Router)
	if !h.closing && err != nil {
		logger.Info("httpServe|server error: " + err.Error())
	}
}

var DefaultHttpServer = NewHttpServer()

func Start(netAddr string) error {
	return DefaultHttpServer.Start(netAddr)
}

func Stop() {
	DefaultHttpServer.Stop()
}

func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	DefaultHttpServer.HandleFunc(pattern, handler)
}
func Handle(pattern string, handler http.Handler) {
	DefaultHttpServer.Handle(pattern, handler)
}

func Router() *mux.Router {
	return DefaultHttpServer.Router
}
