package main

import (
	"os"
	"log"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
const(
	MYSQL_DRIVER_NAME = "mysql"
)


func FailOnError(l *log.Logger, err error, msg string) {
	if err != nil {
		l.Panicf("%s: %s", msg, err)
	}
}
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
