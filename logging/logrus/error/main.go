package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	err := fmt.Errorf("failed")
	logrus.Errorf("err: %v", err)
}
