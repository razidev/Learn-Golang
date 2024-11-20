package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("This is a TraceLevel")
	logger.Debug("This is a DebugLevel")
	logger.Info("This is a InfoLevel")
	logger.Warn("This is a WarnLevel")
	logger.Error("This is a ErrorLevel")
}

func TestSetLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("This is a TraceLevel")
	logger.Debug("This is a DebugLevel")
	logger.Info("This is a InfoLevel")
	logger.Warn("This is a WarnLevel")
	logger.Error("This is a ErrorLevel")
}
