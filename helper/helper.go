package helper

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/zhanglindeng/goweb/config"
)

type panicLogger struct{}

func SetPanicErrorWriter() io.Writer {
	return panicLogger{}
}

func (panicLogger) Write(p []byte) (n int, err error) {
	// TODO 发送通知给管理员
	now := time.Now()
	logFile := config.PanicPath + fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day()) + ".log"
	fd, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return 0, nil
	}
	defer fd.Close()
	return fd.Write(p)
}
