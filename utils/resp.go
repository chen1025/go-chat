package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H struct {
	Code  int
	Msg   string
	Data  interface{}
	Rows  interface{}
	Total interface{}
}

func Resp(w http.ResponseWriter, code int, msg string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code: code,
		Data: data,
		Msg:  msg,
	}
	marshal, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(marshal)
}

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, msg, nil)
}

func RespOk(w http.ResponseWriter, msg string, data interface{}) {
	Resp(w, 0, msg, data)
}

func RespOkList(w http.ResponseWriter, data, total interface{}) {
	RespList(w, 0, data, total)
}

func RespList(w http.ResponseWriter, code int, data, total interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	marshal, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(marshal)
}
