package main

import "log"

type LogLevel int

const (
	DEBUG int = 0
	INFO      = 1

	WARN = 2
	CRITICAL
)

func ReportError(level LogLevel, err error) {
	if level != CRITICAL {
		log.Default().Println(err)
	} else {
		log.Fatalln(err)
	}

}
