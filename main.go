package main

import (
	"fmt"
	"gomao/koala"
	"net/http"
)

func Index (w http.ResponseWriter, r *http.Request ,_ koala.Params) {
	fmt.Fprintf(w, "Index")
}


func Index1 (w http.ResponseWriter, r *http.Request ,_ koala.Params) {
	fmt.Fprintf(w, "Index1")
}

func main() {
	fmt.Println("start")

	koa := koala.New()
	koa.HandlerRouter("GET", "/rank/Index", Index)
	koa.HandlerRouter("GET", "/rank/Index1", Index1)
	koa.HandlerRouter("GET", "/member/Index", Index)
	koa.RUN(":8080")

	//fmt.Println( koa.Test() )
}