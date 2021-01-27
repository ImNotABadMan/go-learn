package mqLog

import "log"

func Log(message string, err error) {
	log.Fatalf("%s, %s", message, err)
}
