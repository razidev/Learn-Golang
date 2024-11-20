package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "razisyahputro").Info("This is a InfoLevel")

	logger.
		WithField("username", "razisyahputro").
		WithField("name", "razi").
		Info("This is a InfoLevel")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "razisyahputro",
		"name":     "razi",
	}).Info("This is a InfoLevel")
}
