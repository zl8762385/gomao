package output

import (
	"net/http"
)

type Output interface {
	OutputContent(w http.ResponseWriter) error
}

// 设置内容类型
func writeContentType(w http.ResponseWriter, value []string) {
	// 获取response的header
	header := w.Header()

	//fmt.Printf("%v %v", header, value)
	// 如果没有找到 就去设置
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
