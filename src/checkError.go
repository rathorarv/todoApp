package src

import "log"

func CheckError(err error) {
	if err != nil {
		log.Fatal("error in %v", err)
	}
}
