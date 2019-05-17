package koala

import (
	"fmt"
	"net/http"
	"sync"
)

type HandlerRouter func (http.ResponseWriter, *http.Request, Params)

// 声明 koala 引擎结构
type Engine struct {
	trees map[string]*node
	pool sync.Pool
}

func New() *Engine {

	engine := &Engine{
	}

	return engine
}

func (e *Engine) Test() string {
	return "123"
}

// 监听端口
func (e *Engine) RUN(addr string) {
	fmt.Printf("listen on %s \n", addr)
	http.ListenAndServe(addr, e)
}

// http中间件实现
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request ) {
	//fmt.Printf("%v == %v", w, req)
	path := req.URL.Path

	if root := e.trees[req.Method]; root != nil {

		if handler, ps, _ := root.getValue(path); handler != nil {
			//fmt.Printf("getValue %+v \n", tsr)
			handler(w, req, ps)
		}


		//fmt.Printf("trees %v \n", e.trees)
		//fmt.Printf("httpmethod %+V \n", req.Method)
		//fmt.Printf("path %v \n", path)
		//fmt.Printf("root %+v \n", root)
		test(w)

	}

}

// 处理router相关节点
func (e *Engine) HandlerRouter(httpMethod, path string, handler HandlerRouter) {
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
