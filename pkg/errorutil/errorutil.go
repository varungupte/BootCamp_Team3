package errorutil

import (
	"fmt"
	"os"
)

func CheckError(err error, msg string)  {
	if err != nil {
		fmt.Println(msg, err.Error())
		os.Exit(1)
	}
}