package main

import (
	"fmt"
	"go-mod-sandbox/internal/libs/gzip"
	"net/http"
)

const (
	//typeApplicationJson = "application/json"
	typeTextPlain = "text/plain"
	encodingGzip  = "gzip"
)

// gzipFunc　は関数定義を返す関数
func gzipFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		v := r.FormValue("k")
		header := map[string]string{
			//"Content-Type": typeApplicationJson,
			"Content-Type":     typeTextPlain,
			"Content-Encoding": encodingGzip,
		}
		for k, v := range header {
			w.Header().Add(k, v)
		}
		w.WriteHeader(403)
		ret, err := gzip.Write(v)
		if err != nil {
			fmt.Println("error1")
			fmt.Println(err.Error())
		}
		unzipped, err := gzip.Read(ret)
		if err != nil {
			fmt.Println("error2")
			fmt.Println(err.Error())
		}
		fmt.Println(unzipped)
		_, _ = w.Write([]byte(ret))
	}
}
