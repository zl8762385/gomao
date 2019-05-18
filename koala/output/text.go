package output

import (
	"fmt"
	"net/http"
)

type Text struct {
	Data interface{}
}

var plainType = []string{"text/plain; charset=utf-8"}

func (s *Text) OutputContent(w http.ResponseWriter) (err error) {
	//设置内容类型，然后输出
	writeContentType(w, plainType)

	fmt.Fprintf(w, "%s", s.Data)
	return
}
