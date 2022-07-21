package runtime

import (
	"bytes"
	"fmt"
	"os/exec"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func GetSystemInfo(osid int) (sysinfo string) {
	switch osid {
	case 1:
	case 2:
		// windows
		cmd := exec.Command("systeminfo")
		var outBuf bytes.Buffer
		cmd.Stdout = &outBuf
		err := cmd.Run()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			sysinfo = ConvertByte2String(outBuf.Bytes(), GB18030)
		}
	case 3:
		// mac os
		cmd := exec.Command("system_profiler", "SPHardwareDataType")
		var outBuf bytes.Buffer
		cmd.Stdout = &outBuf
		err := cmd.Run()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			sysinfo = outBuf.String()
		}
	}
	fmt.Println(sysinfo)
	return sysinfo
}

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}
