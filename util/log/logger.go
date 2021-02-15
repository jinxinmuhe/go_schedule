package log

import (
	"log"
	"os"
)

var InfoLogger *log.Logger
var ErrLogger *log.Logger

func init() {
	infoFile, _ := os.OpenFile("log/schedule.log", os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)
	InfoLogger = log.New(infoFile, "", log.Ldate | log.Ltime | log.Lshortfile)

	errFile, _ := os.OpenFile("log/schedule.log.wf", os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)
	ErrLogger = log.New(errFile, "", log.Ldate | log.Ltime | log.Llongfile)
}