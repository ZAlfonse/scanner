package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

//Log levels
const (
	TRACE  = iota
	DEBUG  = iota
	INFO   = iota
	QUIET  = iota
	SILENT = iota
)

//Log Simple log struct to wrap for colors
type Log struct {
	*log.Logger
	colorCode string
}

//Println mimic of logger.println code so that we can wrap output in colors
func (l *Log) Println(v ...interface{}) {
	s := fmt.Sprintf("\033[1;%v%v\033[0m\n", l.colorCode, fmt.Sprint(v...))
	l.Logger.Output(2, s)
}

//NewLog create a new log with the correct color and outstream
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

func initLogger(loglevel int) {

	var (
		traceHandle   io.Writer
		debugHandle   io.Writer
		successHandle io.Writer
		infoHandle    io.Writer
		warningHandle io.Writer
		errorHandle   io.Writer
	)

	switch loglevel {
	case TRACE:
		traceHandle = os.Stdout
		debugHandle = os.Stdout
		successHandle = os.Stdout
		infoHandle = os.Stdout
		warningHandle = os.Stdout
		errorHandle = os.Stderr
	case DEBUG:
		traceHandle = ioutil.Discard
		debugHandle = os.Stdout
		successHandle = os.Stdout
		infoHandle = os.Stdout
		warningHandle = os.Stdout
		errorHandle = os.Stderr
	case INFO:
		traceHandle = ioutil.Discard
		debugHandle = ioutil.Discard
		successHandle = os.Stdout
		infoHandle = os.Stdout
		warningHandle = os.Stdout
		errorHandle = os.Stderr
	case QUIET:
		traceHandle = ioutil.Discard
		debugHandle = ioutil.Discard
		successHandle = os.Stdout
		infoHandle = os.Stdout
		warningHandle = ioutil.Discard
		errorHandle = os.Stderr
	case SILENT:
		traceHandle = ioutil.Discard
		debugHandle = ioutil.Discard
		successHandle = ioutil.Discard
		infoHandle = ioutil.Discard
		warningHandle = ioutil.Discard
		errorHandle = ioutil.Discard
	}

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
