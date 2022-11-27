package main

import (
	
	"fmt"
	//"log"
	//"net/http"
	"sync"
	"github.com/hxlpub/httpService/httpserver"

)
var (
	version = ""
	goVersion = ""
	buildTime = ""

)
//打印版本信息
func printVersion()  {
	fmt.Println("REST version" + Version )
}
func main() {
	//打印版本信息
	printVersion()
	//初始化日志
	
	//启动http server v1.0 服务
	var wg sync.WaitGroup
	wg.Add(1)
	go httpserver.httpServer(&wg)
	wg.Wait()

}
