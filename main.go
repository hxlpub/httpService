package main

import (
	
	"sync"

        "httpservice/server"
)

func main() {
	/*WaitGroup用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量。每个被等待的线程在结束时应调	   用Done方法。同时，主线程里可以调用Wait方法阻塞至所有线程结束
	*/
	var wg sync.WaitGroup
	wg.Add(1)
	go server.HttpServer(&wg)
	wg.Wait()
}
