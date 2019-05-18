package main

import (
	"gomao/koala"
	"fmt"
	"net/http"
)

//func Index (w http.ResponseWriter, r *http.Request ,_ koala.Params) {
//	fmt.Fprintf(w, "Index")
//}

// 输出 ===========================  start  string例子
func Index2 (ctx *koala.Context) {
	//str, _ := ctx.GetQuery("id")
	//fmt.Println(ctx.Param("name"), ctx.Param("title"))
	//fmt.Printf("%+V", ctx)

	ctx.String(404, "我想要输出数据123.")
}

// 输出 ===========================  start  json例子
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
// 输出 ===========================  start  end 例子

func main() {
	fmt.Println("start")

	koa := koala.New()

	koa.ServeFiles("/static/*filepath", http.Dir("static"))

	koa.Router("GET", "/rank/Index/:name/:title", Index2)
	koa.Router("GET", "/rank/json", Json)

	// param 使用例子 /rank/infos/50
	koa.Router("GET", "/rank/infos/:id", func(ctx *koala.Context) {
		id := ctx.Param("id")
		fmt.Println(id)
	})

	koa.RUN(":8080")

	//fmt.Println( koa.Test() )
}