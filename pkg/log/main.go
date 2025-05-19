package log

import (
	"context"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	// C is an alias for WithContext.
	C = WithContext

	// L is an alias for the standard logger.
	L = logrus.NewEntry(logrus.StandardLogger())
)

func Setup(level string) error {
	if err := setLevel(level); err != nil {
		return err
	}
	logrus.AddHook(addTraceIDHook{})
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File)
			functionName := path.Base(frame.Function)
			return functionName, fileName
		},
	})
	return nil
}

func setLevel(level string) error {
	lvl, err := logrus.ParseLevel(strings.ToLower(level))
	if err != nil {
		return err
	}

	logrus.SetLevel(lvl)
	return nil
}

func WithContext(ctx context.Context) *logrus.Entry {
	return L.WithContext(ctx)
}
