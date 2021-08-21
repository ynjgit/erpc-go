package log

import "go.uber.org/zap"

type zapLogger struct {
	l *zap.Logger
}

// Debug ...
func (z *zapLogger) Debug(args ...interface{}) {
	z.l.Sugar().Debug(args...)
}

// Debugf ...
func (z *zapLogger) Debugf(format string, args ...interface{}) {
	z.l.Sugar().Debugf(format, args...)
}

// Info ...
func (z *zapLogger) Info(args ...interface{}) {
	z.l.Sugar().Info(args...)
}

// Infof ...
func (z *zapLogger) Infof(format string, args ...interface{}) {
	z.l.Sugar().Infof(format, args...)
}

// Warn ...
func (z *zapLogger) Warn(args ...interface{}) {
	z.l.Sugar().Warn(args...)
}

// Warnf ...
func (z *zapLogger) Warnf(format string, args ...interface{}) {
	z.l.Sugar().Warnf(format, args...)

}

// Error ...
func (z *zapLogger) Error(args ...interface{}) {
	z.l.Sugar().Error(args...)
}

// Errorf ...
func (z *zapLogger) Errorf(format string, args ...interface{}) {
	z.l.Sugar().Errorf(format, args...)

}

// Fatal ...
func (z *zapLogger) Fatal(args ...interface{}) {
	z.l.Sugar().Fatal(args...)
}

// Fatalf ...
func (z *zapLogger) Fatalf(format string, args ...interface{}) {
	z.l.Sugar().Fatalf(format, args...)
}

// WithFields ...
func (z *zapLogger) WithFields(fields ...string) Logger {
	zapFields := make([]zap.Field, len(fields)/2)
	for index := range zapFields {
		zapFields[index] = zap.String(fields[2*index], fields[2*index+1])
	}

	return &zapLogger{l: z.l.With(zapFields...)}
}

// Sync ...
func (z *zapLogger) Sync() error {
	return nil
}

// type zapLoggerWrapper struct {
// 	z *zapLogger
// }

// // Debug ...
// func (w *zapLoggerWrapper) Debug(args ...interface{}) {
// 	w.z.Debug(args...)
// }

// // Debugf ...
// func (w *zapLoggerWrapper) Debugf(format string, args ...interface{}) {
// 	w.z.Debugf(format, args...)
// }

// // Info ...
// func (w *zapLoggerWrapper) Info(args ...interface{}) {
// 	w.z.Info(args...)
// }

// // Infof ...
// func (w *zapLoggerWrapper) Infof(format string, args ...interface{}) {
// 	w.z.Infof(format, args...)
// }

// // Warn ...
// func (w *zapLoggerWrapper) Warn(args ...interface{}) {
// 	w.z.Warn(args...)
// }

// // Warnf ...
// func (w *zapLoggerWrapper) Warnf(format string, args ...interface{}) {
// 	w.z.Warnf(format, args...)

// }

// // Error ...
// func (w *zapLoggerWrapper) Error(args ...interface{}) {
// 	w.z.Error(args...)
// }

// // Errorf ...
// func (w *zapLoggerWrapper) Errorf(format string, args ...interface{}) {
// 	w.z.Errorf(format, args...)

// }

// // Fatal ...
// func (w *zapLoggerWrapper) Fatal(args ...interface{}) {
// 	w.z.Fatal(args...)
// }

// // Fatalf ...
// func (w *zapLoggerWrapper) Fatalf(format string, args ...interface{}) {
// 	w.z.Fatalf(format, args...)
// }

// // WithFields ...
// func (w *zapLoggerWrapper) WithFields(fields ...string) Logger {
// 	return w.z.WithFields(fields...)
// }

// // Sync ...
// func (w *zapLoggerWrapper) Sync() error {
// 	return w.z.Sync()
// }
