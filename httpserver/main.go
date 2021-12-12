package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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


	// 使用环境变量到达配置与代码分离
	Port := "8080"
	if os.Getenv("Port") != "" {
		Port = os.Getenv("Port")
	}

	server := &http.Server{Addr: fmt.Sprintf("0.0.0.0:%s", Port), Handler: nil}


	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	// 监听信号优雅关闭
	go func(server *http.Server) {
		log.Println("Waiting signal")
		sig := <-sigs
		log.Printf("Get Signal: %s, call server.Shutdown ", sig.String())
		server.Shutdown(context.Background())
	}(server)


	log.Println("start listen and serve")
	if err := server.ListenAndServe(); err != nil{
		log.Printf("serve error: %s", err)
	}






	//if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
	//	panic("restart error")
	//}








}
