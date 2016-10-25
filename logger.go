package main

import (
	"fmt"
	"io"
	"log"
)

//Log aaaa
type Log struct {
	*log.Logger
	colorCode string
}

//Println aaaa
func (l *Log) Println(v ...interface{}) {
	s := fmt.Sprintf("\033[1;%v%v\033[0m\n", l.colorCode, fmt.Sprint(v...))
	l.Logger.Output(2, s)
}

//NewLog aaa
func NewLog(out io.Writer, prefix string, flag int, colorCode string) *Log {
	return &Log{
		log.New(out, prefix, flag),
		colorCode,
	}
}

var (
	traceLog   *Log
	debugLog   *Log
	successLog *Log
	infoLog    *Log
	warningLog *Log
	errorLog   *Log
)

func initLogger(
	traceHandle io.Writer,
	debugHandle io.Writer,
	successHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	traceLog = NewLog(traceHandle,
		"TRACE: ",
		log.Ltime|log.Lshortfile,
		"32m")

	debugLog = NewLog(debugHandle,
		"DEBUG: ",
		log.Ltime,
		"35m")

	successLog = NewLog(successHandle,
		"SUCCESS: ",
		log.Ltime,
		"32m")

	infoLog = NewLog(infoHandle,
		"INFO: ",
		log.Ltime,
		"36m")

	warningLog = NewLog(warningHandle,
		"WARNING: ",
		log.Ltime,
		"31m")

	errorLog = NewLog(errorHandle,
		"ERROR: ",
		log.Ltime|log.Lshortfile,
		"31m")
}
