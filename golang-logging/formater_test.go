package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("This is a InfoLevel")
	logger.Warn("This is a WarnLevel")
	logger.Error("This is a ErrorLevel")
}
