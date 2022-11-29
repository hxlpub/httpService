package server

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
        "sync"
	"time"
	"github.com/julienschmidt/httprouter"
	"github.com/tidwall/gjson"
)

func ResponseMsg(code, message, dialog string) string {
	resMsg := `{"code":"` + code + `","message":"` + message + `","dialog":"` + dialog + `"}`
	return resMsg
}

func SayHello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bodyBuf := new(bytes.Buffer)
	bodyBuf.ReadFrom(r.Body)
	bodyStr := bodyBuf.String()
	if bodyStr == "" {
		resMsg := ResponseMsg("400", "request body is null", "")
		fmt.Fprint(w, resMsg)
		return
	}

	// 解 析 Body{"dialog":"hello"}
	dialog := gjson.Get(bodyStr, "dialog")
	if !dialog.Exists() {
		resMsg := ResponseMsg("400", "dialog not exist", "")
		fmt.Fprint(w, resMsg)
		return
	}
	dialogValue := dialog.String()
	if dialogValue == "" {
		resMsg := ResponseMsg("400", "dialogValue is null", "")
		fmt.Fprint(w, resMsg)
		return
	}
	var body string =  "world" 
	resMsg := ResponseMsg("200","success",body)
	fmt.Fprint(w, resMsg)
}


func HttpServer(wg *sync.WaitGroup) {
	defer wg.Done()

	router := httprouter.New()

	// say hello service
	router.POST("/hello", SayHello)
	// 启 动 监 听 服 务
	srv := &http.Server{
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      180 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       10 * time.Second,
		Addr:              "0.0.0.0:12345",
		Handler:           router,
	}
	log.Fatal(srv.ListenAndServe())
}
