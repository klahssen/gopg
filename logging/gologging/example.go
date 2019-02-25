package main

import (
	"log"
	"os"

	"github.com/op/go-logging"
)

const (
	prefix = ""
)

var (
	logger = logging.MustGetLogger(prefix)
	//stdFlags  = log.LUTC | log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile
	stdFlags         = log.LUTC | log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	stdInfoLogger    = log.New(os.Stdout, "[INFO] ", stdFlags)
	stdWarningLogger = log.New(os.Stdout, "[WARNING] ", stdFlags)
)

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var goLoggingFormat = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func main() {
	lbk := logging.NewLogBackend(os.Stdout, prefix, 0)
	logging.SetBackend(logging.NewBackendFormatter(lbk, goLoggingFormat))
	stdInfoLogger.Print("first log via stdLogger")
	stdWarningLogger.Print("second log via stdLogger")
	logger.Info("first log via gologging")
	logger.Info("second log via gologging")
}
