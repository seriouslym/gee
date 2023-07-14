package gee

import (
	"net/http"
)

// HandlerFunc 用于处理请求的统一定义
type HandlerFunc func(Context)

type Engine struct {
	router *router
}

// New 类似于构造函数
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRouter(method, pattern, handler)
}

func (engine Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(*c)
}
