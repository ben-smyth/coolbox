package logging

import (
	"context"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
}

func LogWithContext(ctx context.Context, logger *logrus.Logger) *logrus.Entry {
	fields := logrus.Fields{"service": "not_provided"}

	if ctx == nil {
		return logger.WithFields(fields)
	}

	if service, ok := ctx.Value("service").(string); ok {
		fields["service"] = service
	}

	if requestId, ok := ctx.Value("requestId").(string); ok {
		fields["requestId"] = requestId
	}

	return logger.WithFields(fields)
}
