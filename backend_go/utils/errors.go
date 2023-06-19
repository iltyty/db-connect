package utils

import "log"

func CheckError(err error) {
	if err == nil {
		return
	}
	log.Fatal(err)
}
