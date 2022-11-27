package httpserver

import (

	"log"
	"net/http"
	"os"
	"sync"

  client "github.com/hxlpub/httpService/httpclient"
)
func server(wg *sync.WaitGroup) {
	//服务（线程）的最后执行，WaitGroup计数器的值减1
	defer wg.Done()	
	//向DefaultServeMux添加处理器
	http.HandlerFunc("/helloworld", client.client)

	//自定义server
	
	//使用指定监听地址和处理器启动一个http服务器
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))



}
