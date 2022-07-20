package runtime

import (
	"fmt"
	"runtime"
)

func GetSystem() (int, string) {
	optSys := runtime.GOOS
	fmt.Println(optSys)
	if optSys == "linux" {
		return 1, optSys
	} else if optSys == "windows" {
		return 2, optSys
	} else if optSys == "darwin" {
		return 3, optSys
	}
	return 0, "unknow"
}
