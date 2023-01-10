package utils

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func LogError(logger *log.Logger, err error) {
	level.Error(*logger).Log("err", err.Error())
}
