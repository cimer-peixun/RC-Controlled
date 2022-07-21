package remote

import (
	"RemoteControl/pkg/shared"
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

var GConn net.Conn

// send a heartbeat packet
// check if the connection to the server is successful
func TestServerLink() bool {
	var err error
	GConn, err = net.Dial("tcp", "172.16.2.55:5678")
	if err != nil {
		shared.Panic(err)
		return false
	}

	send_str := "hello"
	_, err1 := GConn.Write([]byte(send_str))
	if err1 != nil {
		shared.Panic(err1)
		return false
	}

	buf := make([]byte, 1024)
	cnt, err2 := GConn.Read(buf)
	if err2 != nil {
		shared.Panic(err2)
		return false
	}
	if string(buf[:cnt]) == "OK" {
		return true
	} else {
		return false
	}
}

func SendMsg(msg string) bool {
	cnt, err := GConn.Write([]byte(msg))
	if err != nil {
		return false
	}

	shared.Print(fmt.Sprintf("send: %v\n", string(msg[:cnt])))

	return true
}

func SendFile(fname string) {
	buffer, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("发送文件-读文件-错误")
	}
	GConn.Write(buffer)
}

// receive data from server
func postFile(url, fname string) error {
	//这是一个Post 参数会被返回的地址
	// strinUrl := "http://localhost:8080/aaa"
	byteary, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	resopne, err := http.Post(url, "multipart/form-data", bytes.NewReader(byteary))
	if err != nil {
		return err
	}
	defer func() {
		resopne.Body.Close()
	}()
	body, err := ioutil.ReadAll(resopne.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
