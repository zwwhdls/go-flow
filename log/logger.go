/*
   Copyright 2022 Go-Flow Authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package log

import (
	"fmt"
	glog "log"
	"os"
)

const (
	defaultLogName = ""
	defaultLogFmt  = " [%s] %v\n"
)

type defaultLog struct {
	name string
	*glog.Logger
}

func (l defaultLog) printLogf(level string, format string, v ...interface{}) {
	l.Print(fmt.Sprintf(defaultLogFmt, level, fmt.Sprintf(format, v...)))
}

func (l defaultLog) printLog(level string, v interface{}) {
	l.Print(fmt.Sprintf(defaultLogFmt, level, v))
}

func (l defaultLog) Error(v interface{}) {
	l.printLog("ERROR", v)
}

func (l defaultLog) Errorf(format string, v ...interface{}) {
	l.printLogf("ERROR", format, v...)
}

func (l defaultLog) Warn(v interface{}) {
	l.printLog("WARN", v)
}

func (l defaultLog) Warnf(format string, v ...interface{}) {
	l.printLogf("WARN", format, v...)
}

func (l defaultLog) Info(v interface{}) {
	l.printLog("INFO", v)
}

func (l defaultLog) Infof(format string, v ...interface{}) {
	l.printLogf("INFO", format, v...)
}

func (l defaultLog) Debug(v interface{}) {
	l.printLog("DEBUG", v)
}

func (l defaultLog) Debugf(format string, v ...interface{}) {
	l.printLogf("DEBUG", format, v...)
}

func (l defaultLog) With(name string) Logger {
	if l.name != "" {
		name = fmt.Sprintf("%s.%s", l.name, name)
	}
	return defaultLog{
		name:   name,
		Logger: glog.New(os.Stdout, name+" - ", glog.LstdFlags),
	}
}

func buildDefaultLogger() defaultLog {
	return defaultLog{
		name:   defaultLogName,
		Logger: glog.New(os.Stdout, defaultLogName, glog.LstdFlags),
	}
}

var root Logger = buildDefaultLogger()

func NewLogger(name string) Logger {
	return root.With(name)
}

func SetLogger(logger Logger) {
	root = logger
}
