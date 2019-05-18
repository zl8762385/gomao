package main

import (
	"gomao/koala"
)

//func Index (w http.ResponseWriter, r *http.Request ,_ koala.Params) {
//	fmt.Fprintf(w, "Index")
//}

func Index2 (ctx *koala.Context) {
	ctx.String(404, "我想要输出数据123.")
}

type testJson struct {
	Name string `json:"name"`
	Age int
	Title string
}

func Json(ctx *koala.Context) {
	val := &testJson{
		Name: "xiaoliang",
		Age: 32,
		Title: "ceo",
	}

	ctx.JSON(200, val)
}

func main() {
	//fmt.Println("start")

	koa := koala.New()
	koa.Router("GET", "/rank/Index", Index2)
	koa.Router("GET", "/rank/json", Json)
	//koa.HandlerRouter("GET", "/rank/Index", Index)
	//koa.HandlerRouter("GET", "/rank/Index1", Index1)
	//koa.HandlerRouter("GET", "/member/Index", Index)
	koa.RUN(":8080")

	//fmt.Println( koa.Test() )
}