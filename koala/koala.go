package koala

import (
	"fmt"
	"net/http"
	"sync"
)

// 框架信息
const (
	Version       = "0.0.1"
	FrameworkName = "koala"
	Anthor        = "xiaoliang"
)

//type HandlerRouter func (http.ResponseWriter, *http.Request, Params)
type HandlerRouter func (*Context)

// 声明 koala 引擎结构
type Engine struct {
	version string
	trees   map[string]*node
	pool    sync.Pool
}

func New() *Engine {

	engine := &Engine{
		version:Version,
	}

	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}

	return engine
}

// 建立上下文数据
func (e *Engine) allocateContext() *Context {
	return &Context{engine:e}
}

// 监听端口
func (e *Engine) RUN(addr string) {
	fmt.Printf("listen on %s \n", addr)
	http.ListenAndServe(addr, e)
}

// http中间件实现
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request ) {
	// 获取路径
	path := req.URL.Path

	if root := e.trees[req.Method]; root != nil {

		// 在树上找到对应请求方法，然后执行
		if handler, ps, _ := root.getValue(path); handler != nil {

			// 获取临时对象池
			ctx := e.pool.Get().(*Context)
			ctx.writer = w
			ctx.request = req
			ctx.params = ps
			// 实体方法
			handler(ctx)
			//handler(w, req, ps)
			// 设置
			e.pool.Put(ctx)
			return
		}

		// 都没有找到404操作
		http.NotFound(w, req)
		//http.Error(w, http.StatusText(404), 404)
	}
}

// 注册路由
func (e *Engine) Router(httpMethod, path string, handler HandlerRouter) {
	// 树为空 创建map 分配内存
	if e.trees == nil {
		e.trees = make(map[string]*node)
	}

	// 查找节点是否有该方法
	root := e.trees[httpMethod]
	if root == nil {
		// 没找到method节点 实例化Node
		root = new(node)
		e.trees[httpMethod] = root
	}

	root.addRoute(path, handler)
}

func test(w http.ResponseWriter) {
	fmt.Fprint(w, "shuju")
}
