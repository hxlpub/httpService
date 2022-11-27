package httpclient

import (

	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
        "github.com/tidwall/gjson"
)
// ResponseMsg 组装响应
func ResponseMsg(code, message, dialog string) string {
	resMsg := `{"code":"` + code + `","message":"` + message + `","dialog":"` + dialog + `"}`
	return resMsg
}

// hello world handler
func helloWorld(w http.ResponseWriter, r *http.Request) {
	bodyBuf := new(bytes.Buffer)
	bodyBuf.ReadFrom(r.Body)
	bodyStr := bodyBuf.String()
	if bodyStr == "" {
		resMsg := ResponseMsg("400", "request body is null", "")
		fmt.Fprint(w, resMsg)
		return
	}
	//解析body
	//{"dialog": "hello" }
	value := gjson.Get(bodyStr, "dialog")
	if !value.Exists() {
		resMsg := ResponseMsg("400", "dialog does not exsits", "")
		fmt.Fprint(w, resMsg)
		return
	}
	dialog = value.String()	
	if dialog == "" {
		resMeg := ResponseMsg("400", "dialog is null", "")
		fmt.Fprint(w, resMsg)
		return
	}
	resMeg := ResponseMsg("200", "success", "world")
	fmt.Fpint(w.resMeg)
	//

}
