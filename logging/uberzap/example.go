package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"

	"github.com/op/go-logging"
	"github.com/sirupsen/logrus"
)

const (
	prefix = ""
)

var (
	logger zap.Logger
	//stdFlags  = log.LUTC | log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile
	stdFlags  = log.LUTC | log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	stdLogger = log.New(os.Stderr, "[std] ", stdFlags)
)

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var goLoggingFormat = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	t0 := time.Now()
	logger.Info("first log", zap.String("field1", "field1"), zap.Int("field2", 002))
	fmt.Printf("printed in %s\n", time.Since(t0))
	t0 = time.Now()
	stdLogger.Printf("first log: field1:%v, field2:%v", "field1", 002)
	fmt.Printf("printed in %s\n", time.Since(t0))
	t0 = time.Now()
	type data struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}
	b, _ := json.Marshal(data{Field1: "field1", Field2: 002})
	stdLogger.Printf("%s", string(b))
	fmt.Printf("printed in %s\n", time.Since(t0))
	t0 = time.Now()
	stdLogger.Printf("first log: %+v", map[string]interface{}{"field1": "field1", "field2": 002})
	fmt.Printf("printed in %s\n", time.Since(t0))
	t0 = time.Now()
	logrus.Printf("first log: %+v", map[string]interface{}{"field1": "field1", "field2": 002})
	fmt.Printf("printed in %s\n", time.Since(t0))
}
