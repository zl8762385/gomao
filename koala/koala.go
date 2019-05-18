package koala

import (
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

// 初始化路由结构体
var IRooter Router

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
	debugPrint("监听端口%s \n", addr)
	debugPrintf("%s","[KOALA] 全部服务启动成功")
	http.ListenAndServe(addr, e)
}

// 文件Server
func (e *Engine) ServeFiles(path string, root http.FileSystem) {
	IRooter.handlerServerFiles(path, root, e)
	debugPrint("%s","静态文件服务器,启动完成")
}

// http中间件实现
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request ) {
	// 获取临时对象池
	ctx := e.pool.Get().(*Context)
	ctx.writer = w
	ctx.request = req
	ctx.engine = e

	// 执行处理
	IRooter.ctx = ctx
	IRooter.handlerServeHTTP()

	// 重新写入临时对象池
	e.pool.Put(ctx)
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

func init () {
	debugPrintf("%s","[KOALA] 启动中")
}
