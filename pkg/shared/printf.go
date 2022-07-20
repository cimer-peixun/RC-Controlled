package shared

import (
	"fmt"
)

func Print(info string) {
	if BUILDVARIABLE == "debug" {
		fmt.Print(info)
	}
}
