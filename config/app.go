package config

import "os"

var (
	AppName string
	AppPort string
	AppSecret []byte

	LogPath   = "./storage/logs/"
	PanicPath = "./storage/panic/"

	CachePrefix = "goweb"
)

func app() error {
	AppName = os.Getenv("APP_NAME")
	AppSecret = []byte(os.Getenv("APP_SECRET"))
	AppPort = ":" + os.Getenv("APP_PORT")

	return nil
}
