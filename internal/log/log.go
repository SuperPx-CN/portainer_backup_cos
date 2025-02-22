package log

import "log"

func SetUp() {
	log.SetFlags(log.LstdFlags)
}
