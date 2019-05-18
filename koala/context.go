package koala

import (
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