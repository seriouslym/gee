package main

import (
	"fmt"
	"gee"
	"net/http"
)

/**
Engine is the uni handler for all requests
go http的handler 是一个接口 需要实现一个ServeHTTP的方法
*/

type Engine struct {
	gee.Engine
}

func (engine Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s", req.URL)
	}

}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "URL.Path = %s\n", req.URL.Path)
}
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
func t(values ...interface{}) {
	fmt.Println(values)
	fmt.Println(values...)
}
func main() {
	//r := gee.New()
	//r.GET("/", func(w http.ResponseWriter, req *http.Request) {
	//	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	//})
	//r.Run(":9991")
	r := gee.New()
	r.GET("/", func(context gee.Context) {
		context.HTML(http.StatusOK, "<h1 style='color: red'>Hello Gee</h1>")
	})
	r.Run("localhost:9999")

}
