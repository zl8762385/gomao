package koala

import (
	"net/http"
)

type Router struct {
	ctx *Context
}

// 处理静态文件
func (r *Router) handlerServerFiles(path string, root http.FileSystem, e *Engine) {
	if len(path) < 10 || path[len(path)-10:] != "/*filepath" {
		panic("path must end with /*filepath in path '" + path + "'")
	}

	fileServer := http.FileServer(root)

	e.Router("GET", path, func(ctx *Context) {
		ctx.request.URL.Path = ctx.params.ByName("filepath")
		fileServer.ServeHTTP(ctx.writer, ctx.request)
	})

}

// 处理server http 中间件
func (r *Router) handlerServeHTTP() {
	// 获取浏览器操作行为 GET POST PUT等
	method := r.ctx.request.Method
	// 获取路径
	path := r.ctx.request.URL.Path

	if root := r.ctx.engine.trees[method]; root != nil {

		// 在树上找到对应请求方法，然后执行
		if handler, ps, _ := root.getValue(path); handler != nil {

			r.ctx.params = ps
			// 实体方法
			handler(r.ctx)
			//handler(w, req, ps)
			return
		}

		// 都没有找到404操作
		http.NotFound(r.ctx.writer, r.ctx.request)
		//http.Error(w, http.StatusText(404), 404)
	}
}