// Package logger represents a generic logging interface

package logger

import (
	"github.com/parag08/go-errors/pkg/errors"

	"go.uber.org/zap"
)

// Log is a package level variable, every program should access logging function through "Log"
var Log Logger

func init() {
	if Log == nil {
		logger, _ := zap.NewProduction()
		defer logger.Sync() // flushes buffer, if any
		sugar := logger.Sugar()
		Log = sugar
	}
}

// Logger represent common interface for logging function
type Logger interface {
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Warn(args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	With(args ...interface{}) *zap.SugaredLogger
}

// SetLogger is the setter for log variable, it should be the only way to assign value to log
func SetLogger() {
}

func SystemErr(err error) {
	sysErr, ok := err.(*errors.Error)
	if !ok {
		Log.Error(err)
		return
	}

	entry := Log.With(
		"operations", errors.Ops(*sysErr), "kind", errors.GetKind(*sysErr),
	)

	switch errors.Level(*sysErr) {
	case errors.WarnLevel:
		entry.Warnf("%s : %v", sysErr.Op, err)
	case errors.InfoLevel:
		entry.Infof("%s : %v", sysErr.Op, err)
	case errors.DebugLevel:
		entry.Debugf("%s : %v", sysErr.Op, err)
	default:
		entry.Errorf("%s : %v", sysErr.Op, err)
	}

}
