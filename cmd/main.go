package main

import (
	"RemoteControl/pkg/remote"
	"RemoteControl/pkg/runtime"
	"RemoteControl/pkg/shared"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Whether the server is being connected
var GIsLinkingWithServer = false

//Post("http://xxxx","application/json;charset=utf-8",[]byte("{'aaa':'bbb'}"))
func Post(url string, contentType string, body []byte) (string, error) {
	res, err := http.Post(url, contentType, strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	remote.CaptureFullScreen()
	shared.Print("检查运行环境：")
	osid, osName := runtime.GetSystem()
	shared.Print(fmt.Sprintf("    操作系统：%s\n", osName))
	GIsLinkingWithServer = remote.TestServerLink()
	if GIsLinkingWithServer {
		shared.Print(fmt.Sprintf("    确认服务器连接情况：%s\n", "在线"))
	} else {
		shared.Print(fmt.Sprintf("    确认服务器连接情况：%s\n", "离线"))
	}
	GIsLinkingWithServer = true
	if GIsLinkingWithServer {
		// send os infomation
		remote.SendMsg(fmt.Sprintf("操作系统：%s\n", osName))
		// send the system infomation to the server
		remote.SendMsg(runtime.GetSystemInfo(osid))
	}
}
