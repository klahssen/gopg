package hello

import (
	"github.com/klahssen/gopg/logging/logrus/formatter"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&formatter.Formatter{NoColors: true})
}

//Do is just logging something, reimporting logrus
func Do() {
	logrus.Info("hello: done")
}
