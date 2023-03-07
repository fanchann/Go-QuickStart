package utils

import "log"

func LogErrorWithPanic(err error) {
	if err != nil {
		log.Panicf("error : %s", err)
	}
}
