package koala

import (
	"fmt"
	"gomao/koala/output"
	"net/http"
)

// 上下文结构体
type Context struct {
	writer 		http.ResponseWriter
	request     *http.Request
	params      Params
	engine      *Engine
}

//////////////////////////////////////////////////////////////////////////////
// 获取
//////////////////////////////////////////////////////////////////////////////

// 设置：koa.Router("GET", "/rank/Index/:name/:title", Index2)
// 获取：fmt.Println(ctx.Param("name"), ctx.Param("title"))
//	koa.Router("GET", "/rank/infos/:id", func(ctx *koala.Context) {
//		id := ctx.Param("id")
//		fmt.Println(id)
// })
func (c *Context) Param(key string) string {
	fmt.Println("getquery")
	return c.params.ByName(key)
}

// 获取URL请求
// rank/index?id=51&name=xiaoliang
// c.GetQuery("id") (51)
// @param key string 需要获取参数 id=51
// @return string bool
func (c *Context) GetQuery(key string) (string, bool) {
	if val, ok := c.GetQueryArray(key); ok {
		return val[0], ok
	}

	return "", false
}

// 封装获取URL http request请求, 并检查字符串是否存在，返回对应数据
// @param key string 需要获取参数 id=51
// @return []string, bool
func (c *Context) GetQueryArray(key string) ([]string, bool) {
	if val, ok := c.request.URL.Query()[key]; ok && len(val) > 0 {
		return val, true
	}

	return []string{}, false
}

//////////////////////////////////////////////////////////////////////////////
// 输出到页面
//////////////////////////////////////////////////////////////////////////////
// 输出字符串
func (c *Context) String(httpCode int, data interface{}) {
	//c.writer.WriteHeader(200)
	c.Output(httpCode, &output.Text{
		Data:data,
	})
}

// 输出JSON
func (c *Context) JSON(httpCode int, data interface{}) {
	c.Output(httpCode, &output.Json{
		Data:data,
	})
}

// 设置 http response code
func (c *Context) StatusCode(code int) {
	c.writer.WriteHeader(code)
}

// 上下文 输出引擎
func (c *Context) Output (code int, o output.Output) {
	// set http response code
	c.StatusCode(code)

	// 输出内容
	o.OutputContent(c.writer)
}