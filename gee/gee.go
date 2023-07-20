package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc 用于处理请求的统一定义
type HandlerFunc func(Context)

type Engine struct {
	router *router
	*RouterGroup
	groups []*RouterGroup
}
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	engine      *Engine
}

// New 类似于构造函数
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}
func (group RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	fmt.Print(engine.groups[0].engine == engine)
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRouter(method, pattern, handler)
}

func (group RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (engine Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(*c)
}
