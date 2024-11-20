package golanglogging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampleHook struct{}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample Hook", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})

	logger.Info("This is a InfoLevel")
	logger.Warn("This is a WarnLevel")
	logger.Error("This is a ErrorLevel")
}
