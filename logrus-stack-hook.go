package logrus_stack

import (
	"github.com/Sirupsen/logrus"
	"github.com/facebookgo/stack"
)

// NewHook is the initializer for logrusStackHook{} (implementing logrus.Hook).
// Set levels to callerLevels for which "caller" value may be set, providing a
// single frame of stack. Set levels to stackLevels for which "stack" value may
// be set, providing the full stack (minus logrus).
func NewHook(callerLevels []logrus.Level, stackLevels []logrus.Level) logrusStackHook {
	return logrusStackHook{
		callerLevels: callerLevels,
		stackLevels: stackLevels,
	}
}

// StandardHook is a convenience initializer for logrusStackHook{} with
// default args.
func StandardHook() logrusStackHook {
	return logrusStackHook{
		callerLevels: logrus.AllLevels,
		stackLevels: []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel},
	}
}

// logrusStackHook is an implementation of logrus.Hook interface.
type logrusStackHook struct {
	// Set levels to callerLevels for which "caller" value may be set,
	// providing a single frame of stack.
	callerLevels []logrus.Level

	// Set levels to stackLevels for which "stack" value may be set,
	// providing the full stack (minus logrus).
	stackLevels  []logrus.Level
}

// Levels provides the levels to filter.
func (hook logrusStackHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire is called by logrus when something is logged.
func (hook logrusStackHook) Fire(entry *logrus.Entry) error {
	var skipFrames int
	if len(entry.Data) == 0 {
		// When WithField(s) is not used, we have 8 logrus frames to skip.
		skipFrames = 8
	} else {
		// When WithField(s) is used, we have 6 logrus frames to skip.
		skipFrames = 6
	}

	// Get the complete stack track past skipFrames count.
	frames := stack.Callers(skipFrames)

	if len(frames) > 0 {
		// If we have a frame, we set it to "caller" field for assigned levels.
		for _, level := range hook.callerLevels {
			if entry.Level == level {
				entry.Data["caller"] = frames[0]
				break
			}
		}

		// Set the available frames to "stack" field.
		for _, level := range hook.stackLevels {
			if entry.Level == level {
				entry.Data["stack"] = frames
				break
			}
		}
	}

	return nil
}
