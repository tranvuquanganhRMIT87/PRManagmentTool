package components

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.SetLevel(logrus.InfoLevel)

	logrus.SetOutput(os.Stdout)
}
