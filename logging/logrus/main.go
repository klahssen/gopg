package main

import (
	formatter "github.com/klahssen/gopg/logging/logrus/formatter"
	"github.com/klahssen/gopg/logging/logrus/hello"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&formatter.Formatter{})
	logrus.Info("log from main")
	hello.Do()
}
