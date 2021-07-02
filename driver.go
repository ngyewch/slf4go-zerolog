package slf4go_zerolog

import (
	"fmt"
	slog "github.com/go-eden/slf4go"
	"github.com/rs/zerolog"
)

type ZerologDriver struct {
	logger *zerolog.Logger
}

func NewZerologDriver(logger *zerolog.Logger) *ZerologDriver {
	return &ZerologDriver{
		logger: logger,
	}
}

func (d *ZerologDriver) Name() string {
	return "slf4go-zerolog"
}

func (d *ZerologDriver) Print(l *slog.Log) {
	var event *zerolog.Event
	switch l.Level {
	case slog.TraceLevel:
		event = d.logger.Trace()
		break
	case slog.DebugLevel:
		event = d.logger.Debug()
		break
	case slog.InfoLevel:
		event = d.logger.Info()
		break
	case slog.WarnLevel:
		event = d.logger.Warn()
		break
	case slog.ErrorLevel:
		event = d.logger.Error()
		break
	case slog.PanicLevel:
		event = d.logger.Panic().Stack()
		break
	case slog.FatalLevel:
		event = d.logger.Fatal().Stack()
		break
	default:
		event = d.logger.WithLevel(zerolog.GlobalLevel())
		break
	}

	if l.Fields != nil {
		for k, v := range l.Fields {
			event = event.Interface(k, v)
		}
	}

	if l.Format == nil {
		for _, msg := range l.Args {
			event.Msgf("%s %v", l.Logger, msg)
			break
		}
	} else {
		event.Msgf("%s %s", l.Logger, fmt.Sprintf(*l.Format, l.Args...))
	}
}

func (d *ZerologDriver) GetLevel(logger string) (sl slog.Level) {
	l := zerolog.GlobalLevel()

	switch l {
	case zerolog.DebugLevel:
		sl = slog.DebugLevel
	case zerolog.InfoLevel:
		sl = slog.InfoLevel
	case zerolog.WarnLevel:
		sl = slog.WarnLevel
	case zerolog.ErrorLevel:
		sl = slog.ErrorLevel
	case zerolog.FatalLevel:
		sl = slog.FatalLevel
	case zerolog.PanicLevel:
		sl = slog.PanicLevel
	default:
		sl = slog.TraceLevel
	}
	return
}
