package snlogic

import "log"

var IsLogEnabledFlag = false

func SnLog(message string) {
	if IsLogEnabledFlag {
		log.Print(message)
	}

}
