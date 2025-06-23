package tools

import (
	"time"

	"github.com/sirupsen/logrus"
)

const (
	LogsFilePath = "logs/app.log"
	LogLevel     = logrus.DebugLevel
	TimeFormat   = time.RFC3339
)
