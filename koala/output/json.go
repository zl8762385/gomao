package output

import (
	"encoding/json"
	"net/http"
)

type Json struct {
	Data interface{}
}

var jsonType = []string{"application/json; charset=utf-8"}

func (j *Json) OutputContent(w http.ResponseWriter) (err error) {
	// 设置JSON content-type
	writeContentType(w, jsonType)

	jsonBytes, err := json.Marshal(j.Data)
	if err != nil {
		return err
	}

	// 写入字节组
	w.Write(jsonBytes)
	//fmt.Printf("%+v == %+v", string(jsonBytes), j.Data)
	//fmt.Fprintf(w, "%s", string(jsonBytes))

	return
}