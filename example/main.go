package main

import (
	"errors"
	"os"

	"github.com/Gurpartap/logrus-stack"
	"github.com/sirupsen/logrus"
)

type Worker struct {
	JobID string
}

func (w Worker) Perform() {
	logrus.WithField("jod_id", w.JobID).Infoln("Now working")

	err := errors.New("I don't know what to do yet")
	if err != nil {
		logrus.Errorln(err)
		return
	}

	// ...
}

func main() {
	// Setup logrus.
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stderr)

	// Add the stack hook.
	logrus.AddHook(logrus_stack.StandardHook())

	// Let's try it.
	Worker{"123"}.Perform()
}
