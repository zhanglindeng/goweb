package config

import "os"

var (
	AppName string
	AppPort string

	LogPath   = "./storage/logs/"
	PanicPath = "./storage/panic/"

	CachePrefix = "goweb"
)

func app() error {
	AppName = os.Getenv("APP_NAME")
	AppPort = ":" + os.Getenv("APP_PORT")

	return nil
}
