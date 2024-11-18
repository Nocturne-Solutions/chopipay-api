package exceptions

import (
	"fmt"
	"log"
)

func PanicException(err error, errMsg string) {
	if err != nil {
		errMsg = fmt.Sprintf("%s: %s", errMsg, err.Error())
		log.Panic(errMsg)
	}
}
