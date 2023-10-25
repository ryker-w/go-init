package log

import (
	"fmt"
	delegate "github.com/lishimeng/log4go"
)

type Level int

const (
	FINEST Level = iota
	FINE
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	CRITICAL
)

var levels = map[string]Level{
	"FINEST":   FINEST,
	"FINE":     FINE,
	"DEBUG":    DEBUG,
	"TRACE":    TRACE,
	"INFO":     INFO,
	"WARNING":  WARNING,
	"ERROR":    ERROR,
	"CRITICAL": CRITICAL,
}

func init() {
	delegate.CallerSkip = 3 // wrapper一层skip+1
}

func FormatLevel(lvl string) (level Level, err error) {
	if l, ok := levels[lvl]; ok {
		level = l
	} else {
		err = fmt.Errorf("UNKNOWN LEVEL:%s", lvl)
	}
	return
}

func SetLevelAll(lvl Level) {
	for _, value := range delegate.Global {
		value.Level = delegate.Level(lvl)
	}
}

// 动态修改log级别
func SetLevel(name string, lvl Level) {
	if filter, ok := delegate.Global[name]; ok {
		filter.Level = delegate.Level(lvl)
	}
}

func AddFileLog(name string, lvl Level, file string, rotate bool, daily bool, format string, maxlines int, maxsize int, categorys ...string) {
	writer := delegate.NewFileLogWriter(file, rotate, daily)
	writer.SetFormat(format)
	writer.SetRotateLines(maxlines)
	writer.SetRotateSize(maxsize)
	delegate.Global.AddFilter(name, delegate.Level(lvl), writer, categorys...)
}

func Remove(name string) {
	defer func() {
		_ = recover()
	}()
	if item, ok := delegate.Global[name]; ok {
		delete(delegate.Global, name)
		if item != nil {
			item.Close()
			item = nil
		}
	}
}

func Close() {
	delegate.Close()
}

func Log(lvl Level, source, message string) {
	delegate.Log(delegate.Level(lvl), source, message)
}

func Fine(arg0 interface{}, args ...interface{}) {
	delegate.Fine(arg0, args...)
}

func Debug(arg0 interface{}, args ...interface{}) {
	delegate.Debug(arg0, args...)
}

func Trace(arg0 interface{}, args ...interface{}) {
	delegate.Trace(arg0, args...)
}
func Info(arg0 interface{}, args ...interface{}) {
	delegate.Info(arg0, args...)
}

func Warn(arg0 interface{}, args ...interface{}) error {
	return delegate.Warn(arg0, args...)
}
func Error(arg0 interface{}, args ...interface{}) error {
	return delegate.Error(arg0, args...)
}

func Critical(arg0 interface{}, args ...interface{}) error {
	return delegate.Critical(arg0, args...)
}

func Finest(arg0 interface{}, args ...interface{}) {
	delegate.Finest(arg0, args...)
}
