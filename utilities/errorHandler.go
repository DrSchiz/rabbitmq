package utilities

import "log"

func ErrorHandler(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
