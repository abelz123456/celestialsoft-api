package log

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type Level string

type option struct {
	logger *logrus.Logger
}

func NewLog() Log {
	log.SetOutput(os.Stdout)

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&prefixed.TextFormatter{})

	return &option{
		logger: logger,
	}
}

func (o *option) Info(message string, attr interface{}) {
	file, _ := Trace(3)

	dir, err := os.Getwd()
	if err == nil {
		file = strings.ReplaceAll(file, dir+"/", "")
	}

	o.logger.WithFields(logrus.Fields{
		"_file":      fmt.Sprintf("\"%s\"", file),
		"attributes": attr,
	}).Info(message)
}

func (o *option) Warning(message string, _err error, attr interface{}) {
	file, _ := Trace(3)

	dir, err := os.Getwd()
	if err == nil {
		file = strings.ReplaceAll(file, dir+"/", "")
	}

	if _err != nil {
		if message != "" {
			message += " | Error: "
		} else {
			message += "Error: "
		}
	}

	o.logger.WithFields(logrus.Fields{
		"_file":      fmt.Sprintf("\"%s\"", file),
		"attributes": attr,
	}).Warning(message, _err)
}

func (o *option) Error(_err error, message string, attr interface{}) {
	file, _ := Trace(3)

	dir, err := os.Getwd()
	if err == nil {
		file = strings.ReplaceAll(file, dir+"/", "")
	}

	if _err != nil {
		if message != "" {
			message += " | Error: "
		} else {
			message += "Error: "
		}
	}

	o.logger.WithFields(logrus.Fields{
		"_file":      fmt.Sprintf("\"%s\"", file),
		"attributes": attr,
	}).Error(message, _err)
}

func (o *option) Panic(_err error, message string, attr interface{}) {
	file, _ := Trace(3)

	dir, err := os.Getwd()
	if err == nil {
		file = strings.ReplaceAll(file, dir+"/", "")
	}

	if _err != nil {
		if message != "" {
			message += " | Error: "
		} else {
			message += "Error: "
		}
	}

	o.logger.WithFields(logrus.Fields{
		"_file":      fmt.Sprintf("\"%s\"", file),
		"attributes": attr,
	}).Panic(message, _err)
}

func (o *option) PanicOnError(_err error, message string, attr interface{}) {
	if _err == nil {
		return
	}

	o.Panic(_err, message, attr)
}
