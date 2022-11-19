package server

import "os"

const DATE_FORMAT = "2006-01-02T15:04:05-0700"

func GetConfigLogs() bool {
	if os.Getenv("LOG_FULL") == "false" {
		return false
	}
	return true
}
