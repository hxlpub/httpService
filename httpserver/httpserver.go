package httpserver

import (

	"log"
	"net/http"
	"os"
	"sync"

	"httpclient"
)
func httpServer(wg *sync.WaitGroup) {
	//服务（线程）的最后执行，WaitGroup计数器的值减1
	defer wg.Done()	
	//向DefaultServeMux添加处理器
	http.HandlerFunc("/hello-world", httpclient.helloWorld)

	//自定义server
	
	//使用指定监听地址和处理器启动一个http服务器
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))



}
