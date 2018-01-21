package log

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/zhanglindeng/goweb/config"
)

type Level int

const (
	// 严重程度从低到高
	LevelDebug     Level = iota
	LevelInfo
	LevelNotice
	LevelWarning
	LevelError
	LevelCritical
	LevelAlert
	LevelEmergency
)

const (
	logFileMode os.FileMode = 0666
	logDirMode  os.FileMode = 0777

	// 1024*1024 == 1<<20 // 1 MB
	maxSize = 100 << 20 // 100 MB
)

var mutex sync.Mutex

func Debug(s string) error {
	return log(s, LevelDebug)
}

func Info(s string) error {
	return log(s, LevelInfo)
}

func Notice(s string) error {
	return log(s, LevelNotice)
}

func Warning(s string) error {
	return log(s, LevelWarning)
}

func Error(s string) error {
	return log(s, LevelError)
}

func Critical(s string) error {
	return log(s, LevelCritical)
}

func Alert(s string) error {
	return log(s, LevelAlert)
}

func Emergency(s string) error {
	return log(s, LevelEmergency)
}

func Log(s string, level Level) error {
	return log(s, level)
}

func log(s string, level Level) error {
	now := time.Now()

	mutex.Lock()
	defer mutex.Unlock()

	logDir := fmt.Sprintf(config.LogPath+"%d", now.Year())
	if err := os.MkdirAll(logDir, logDirMode); err != nil {
		return err
	}
	logFile := logDir + "/" + fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day()) + ".log"
	fd, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, logFileMode)
	if err != nil {
		return err
	}
	defer fd.Close()

	// check max size
	info, err := fd.Stat()
	if err != nil {

		return err
	}

	if info.Size() > maxSize {
		if err := copyLog(fd, fmt.Sprintf("%s.%d.log", logFile, now.Unix())); err != nil {
			return err
		}
		//if err := fd.Truncate(0); err != nil {
		//	return err
		//}
		if err := os.Truncate(logFile, 0); err != nil {
			return err
		}
	}

	// now.Format("2006-01-02 15:04:05")
	content := fmt.Sprintf("[%s][%s:%d]%s\n", now.Format(time.RFC3339Nano), levelName(level), level, s)
	if _, err := fd.Write([]byte(content)); err != nil {
		return err
	}

	return nil
}

func copyLog(src *os.File, name string) error {
	dst, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, logFileMode)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func levelName(level Level) string {
	switch level {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelNotice:
		return "notice"
	case LevelWarning:
		return "warning"
	case LevelError:
		return "error"
	case LevelCritical:
		return "critical"
	case LevelAlert:
		return "alert"
	case LevelEmergency:
		return "emergency"
	default:
		return "unknown"
	}
}
