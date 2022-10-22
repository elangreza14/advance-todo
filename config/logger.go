package config

import (
	"errors"
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	EncodingType string

	LoggerOption struct {
		EncodingType     EncodingType
		NameService      string
		EnableStackTrace bool
		IsDebug          bool
	}

	ILogger interface {
		Info(logStr string, payload ...interface{})
		Error(logStr string, payload ...interface{})
		Debug(logStr string, payload ...interface{})
		Warn(logStr string, payload ...interface{})
	}

	logger struct {
		zapBase *zap.Logger
		IsDebug bool
	}
)

const (
	EncodingTypeJson    EncodingType = "json"
	EncodingTypeConsole EncodingType = "console"
)

func (c LoggerOption) apply(config *Configuration) error {
	configZap := zap.Config{
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		DisableCaller:    true,
		Encoding:         "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			TimeKey:     "time",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		},

		Level: zap.NewAtomicLevelAt(zapcore.DebugLevel),
	}

	if c.EncodingType == EncodingTypeConsole {
		configZap.Encoding = string(EncodingTypeConsole)
		configZap.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if c.EnableStackTrace {
		configZap.EncoderConfig.StacktraceKey = "stacktrace"
	}

	if c.NameService != "" {
		appName := make(map[string]interface{})
		appName["service"] = c.NameService
		configZap.InitialFields = appName
	}

	zapBuilder, err := configZap.Build()
	if err != nil {
		return err
	}

	config.Logger = &logger{
		zapBase: zapBuilder,
		IsDebug: c.IsDebug,
	}

	return nil
}

func WithLogger(c LoggerOption) Option {
	return LoggerOption(c)
}

func setField(params ...interface{}) []zap.Field {
	res := []zap.Field{}

	for i := 0; i < len(params); i++ {
		typePrams := reflect.TypeOf(params[i])

		if typePrams == nil {
			index := fmt.Sprintf("nil-%v", i+1)
			res = append(res, zap.Any(index, params[i]))
			continue
		}

		switch typePrams.Kind().String() {
		case "chan":
			index := fmt.Sprintf("chan-%v", i+1)
			res = append(res, zap.Any(index, "cannot log chan"))
		case "ptr":
			switch typePrams.String() {
			case "*errors.errorString":
				index := fmt.Sprintf("error-%v", i+1)
				res = append(res, zap.Any(index, params[i]))
			default:
				index := fmt.Sprintf("pointer-%v", i+1)
				res = append(res, zap.Any(index, params[i]))
			}
		default:
			index := fmt.Sprintf("data-%v", i+1)
			res = append(res, zap.Any(index, params[i]))
		}
	}

	var skip int
	// if s := opts.GetSkip(); s != nil {
	// 	skip = *s
	// } else if l.skip != nil {
	// 	skip = *l.skip
	// }

	if ci, err := retrieveCallInfo(skip); err == nil {
		res = append(res, []zap.Field{
			zap.String("package", ci.Package),
			zap.String("function", ci.Function),
			zap.String("file", ci.File),
			zap.Int("line", ci.Line),
		}...)
	}

	return res
}

type callInfo struct {
	Package  string
	Function string
	File     string
	Line     int
}

func retrieveCallInfo(skip int) (*callInfo, error) {
	skip = 3 + skip // omit stacks of logger library call
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return nil, errors.New("failed to get call info")
	}
	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	return &callInfo{
		Package:  packageName,
		Function: funcName,
		File:     fileName,
		Line:     line,
	}, nil
}

func (l *logger) Info(logStr string, payload ...interface{}) {
	base := setField(payload...)
	l.zapBase.With(base...).Info(logStr)
}

func (l *logger) Warn(logStr string, payload ...interface{}) {
	base := setField(payload...)
	l.zapBase.With(base...).Warn(logStr)
}

func (l *logger) Error(logStr string, payload ...interface{}) {
	base := setField(payload...)
	l.zapBase.With(base...).Error(logStr)
}

func (l *logger) Debug(logStr string, payload ...interface{}) {
	if !l.IsDebug {
		return
	}
	base := setField(payload...)
	l.zapBase.With(base...).Debug(logStr)
}
