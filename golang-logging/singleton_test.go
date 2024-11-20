package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("This is a InfoLevel")
	logrus.Warn("This is a WarningLevel")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("This is a InfoLevel")
	logrus.Warn("This is a WarningLevel")
}
