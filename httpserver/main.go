package main

import (
	"log"
	"net/http"
	"os"
)


func Handler(writer http.ResponseWriter, request *http.Request)  {


	for key, value := range request.Header{
		for _, headerVal := range value{
			writer.Header().Set(key, headerVal)
		}
	}
	// 读取当前系统的环境变量中的VERSION配置，并写入response header
	writer.Header().Set("VERSION", os.Getenv("VERSION"))

	var resonseCode = 200
	writer.WriteHeader(resonseCode)
	writer.Write([]byte("hi"))

	log.Printf("client_address: %s, response_code: %d, uri: %s", request.RemoteAddr, resonseCode, request.URL.Path)

}


func main()  {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/healthz", Handler)

	http.ListenAndServe("0.0.0.0:8080", nil)
}
