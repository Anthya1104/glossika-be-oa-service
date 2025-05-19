package log

import "github.com/sirupsen/logrus"

type addTraceIDHook struct{}

func (t addTraceIDHook) Levels() []logrus.Level { return logrus.AllLevels }

func (t addTraceIDHook) Fire(entry *logrus.Entry) error {
	ctx := entry.Context
	if ctx == nil {
		return nil
	}

	traceID, ok := ctx.Value("TraceId").(string)
	if ok {
		entry.Data["traceId"] = traceID
	}

	return nil
}
