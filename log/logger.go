package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	funcCallDepth = 3
	errLabel      = "[E]"
	debugLabel    = "[D]"
	InfoLabel     = "[I]"
)

func Err(str string, param ...interface{}) {
	Write(errLabel, str, param...)
}

func Debug(str string, param ...interface{}) {
	Write(debugLabel, str, param...)
}

func Info(str string, param ...interface{}) {
	Write(InfoLabel, str, param...)
}

// Write 日志写入文件
func Write(label, str string, param ...interface{}) {
	// 判断路径是否存在
	if _, e := os.Stat(filePath); os.IsNotExist(e) {
		e := os.MkdirAll(filePath, os.ModePerm)
		if e != nil {
			log.Printf("logger|Write makedir e:%v", e)
			return
		}
	}
	fileName := fmt.Sprintf("%v/%v.log", filePath, time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Printf("logger|Write openFile err:%v", err)
		return
	}
	logger := log.New(file, label, log.LstdFlags|log.Ldate|log.Ltime|log.Lmsgprefix)
	logger.Printf(fmt.Sprintf(getFileInfo()+str, param...))
}

func getFileInfo() string {
	// 方法调用深度
	_, s, line, _ := runtime.Caller(funcCallDepth)
	i := strings.LastIndex(s, "/")
	fileName := s[i+1:]
	return fmt.Sprintf(" [%v:%d] ", fileName, line)
}
