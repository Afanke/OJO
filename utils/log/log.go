package log

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"time"
)

type Level uint8

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

type LConfig struct {
	Log Config `json:"Log"`
}

type Config struct {
	ConsoleLevel  Level
	FileLevel     Level
	EnableConsole bool
	EnableFile    bool
	FilePath      string
	FileChanSize  int
	MaxSize       int64
}

var cfg Config

var file FileLogger

var l Logger

type Logger struct {
}

func (l Logger) Printf(format string, v ...interface{}) {
	log.Debug(format, v)
}

type FileLogger struct {
	debug   *os.File
	info    *os.File
	warn    *os.File
	error   *os.File
	fatal   *os.File
	logChan chan *FileLog
}

type FileLog struct {
	Level Level
	Msg   string
}

func init() {
	var defaultCfg = Config{
		ConsoleLevel:  0,
		FileLevel:     0,
		EnableConsole: true,
		EnableFile:    false,
		FilePath:      "",
		MaxSize:       0,
		FileChanSize:  10000,
	}
	var lcfg = LConfig{}
	bytes, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		l.Print(ERROR, err.Error())
		l.Print(ERROR, "use default log config")
		cfg = defaultCfg
		return
	}
	err = json.Unmarshal(bytes, &lcfg)
	if err != nil {
		l.Print(ERROR, err.Error())
		l.Print(ERROR, "use default log config")
		cfg = defaultCfg
		return
	}
	cfg = lcfg.Log
	file.Init()
	// Print(INFO, "%v", cfg)
	l.Print(INFO, "success to load log system")
}

func GetLogger() Logger {
	return l
}

func (f *FileLogger) LogBG() {
	i := 0
	for {
		i++
		select {
		case log := <-f.logChan:
			switch log.Level {
			case DEBUG:
				if f.debug == nil {
					l.Print(ERROR, "can't open file:%s/debug.log", cfg.FilePath)
					break
				}
				_, err := f.debug.WriteString(log.Msg)
				if err != nil {
					l.Print(ERROR, "%v", err)
					break
				}
				stat, err := f.debug.Stat()
				if err != nil {
					l.Print(ERROR, "check file stat error:%v", err)
					break
				}
				if stat.Size() < cfg.MaxSize {
					break
				}
				logName := f.debug.Name()
				err = f.debug.Close()
				if err != nil {
					l.Print(ERROR, "file close error:%v", err)
					break
				}
				newLogName := logName + ".bak" + time.Now().Format("2006-01-02_15_04_05")
				err = os.Rename(logName, newLogName)
				if err != nil {
					l.Print(ERROR, "file rename error:%v", err)
					break
				}
				logFile, err := os.OpenFile(logName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
				if err != nil {
					l.Print(ERROR, "file open error:%v", err)
					return
				}
				f.debug = logFile
			case INFO:
				if f.info == nil {
					l.Print(ERROR, "can't open file:%s/info.log", cfg.FilePath)
					break
				}
				_, err := f.info.WriteString(log.Msg)
				if err != nil {
					l.Print(ERROR, "%v", err)
					break
				}
				stat, err := f.info.Stat()
				if err != nil {
					l.Print(ERROR, "check file stat error:%v", err)
					break
				}
				if stat.Size() < cfg.MaxSize {
					break
				}
				logName := f.info.Name()
				err = f.info.Close()
				if err != nil {
					l.Print(ERROR, "file close error:%v", err)
					break
				}
				newLogName := logName + ".bak" + time.Now().Format("2006-01-02_15_04_05")
				err = os.Rename(logName, newLogName)
				if err != nil {
					l.Print(ERROR, "file rename error:%v", err)
					break
				}
				logFile, err := os.OpenFile(logName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
				if err != nil {
					l.Print(ERROR, "file open error:%v", err)
					return
				}
				f.info = logFile
			case WARN:
				if f.warn == nil {
					l.Print(ERROR, "can't open file:%s/warn.log", cfg.FilePath)
					break
				}
				_, err := f.warn.WriteString(log.Msg)
				if err != nil {
					l.Print(ERROR, "%v", err)
					break
				}
				stat, err := f.warn.Stat()
				if err != nil {
					l.Print(ERROR, "check file stat error:%v", err)
					break
				}
				if stat.Size() < cfg.MaxSize {
					break
				}
				logName := f.warn.Name()
				err = f.warn.Close()
				if err != nil {
					l.Print(ERROR, "file close error:%v", err)
					break
				}
				newLogName := logName + ".bak" + time.Now().Format("2006-01-02_15_04_05")
				err = os.Rename(logName, newLogName)
				if err != nil {
					l.Print(ERROR, "file rename error:%v", err)
					break
				}
				logFile, err := os.OpenFile(logName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
				if err != nil {
					l.Print(ERROR, "file open error:%v", err)
					return
				}
				f.warn = logFile
			case ERROR:
				if f.error == nil {
					l.Print(ERROR, "can't open file:%s/error.log", cfg.FilePath)
					break
				}
				_, err := f.error.WriteString(log.Msg)
				if err != nil {
					l.Print(ERROR, "%v", err)
					break
				}
				stat, err := f.error.Stat()
				if err != nil {
					l.Print(ERROR, "check file stat error:%v", err)
					break
				}
				if stat.Size() < cfg.MaxSize {
					break
				}
				logName := f.error.Name()
				err = f.error.Close()
				if err != nil {
					l.Print(ERROR, "file close error:%v", err)
					break
				}
				newLogName := logName + ".bak" + time.Now().Format("2006-01-02_15_04_05")
				err = os.Rename(logName, newLogName)
				if err != nil {
					l.Print(ERROR, "file rename error:%v", err)
					break
				}
				logFile, err := os.OpenFile(logName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
				if err != nil {
					l.Print(ERROR, "file open error:%v", err)
					return
				}
				f.error = logFile
			case FATAL:
				err := ioutil.WriteFile(cfg.FilePath+"/fatal.log", []byte(log.Msg), 0666)
				if err != nil {
					l.Print(ERROR, "%v\n", err)
					return
				}
				os.Exit(1)
			}
		}
	}
}

func (f *FileLogger) Init() {
	if !cfg.EnableFile {
		l.Print(INFO, "no log file loaded")
		return
	}
	switch cfg.FileLevel {
	case DEBUG:
		debug, err := os.OpenFile(cfg.FilePath+"/debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			l.Print(ERROR, "%v", err)
		}
		f.debug = debug
		fallthrough
	case INFO:
		info, err := os.OpenFile(cfg.FilePath+"/info.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			l.Print(ERROR, "%v", err)
		}
		f.info = info
		fallthrough
	case WARN:
		warn, err := os.OpenFile(cfg.FilePath+"/warn.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			l.Print(ERROR, "%v", err)
		}
		f.warn = warn
		fallthrough
	case ERROR:
		erro, err := os.OpenFile(cfg.FilePath+"/error.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			l.Print(ERROR, "%v", err)
		}
		f.error = erro
	}
	f.logChan = make(chan *FileLog, 1000000)
	go f.LogBG()
}

func (l Logger) Log(level Level, format string, a ...interface{}) {
	s := ""
	if cfg.EnableConsole && level >= cfg.ConsoleLevel {
		s = formatString(level, format, a...)
		fmt.Print(s)
	}
	if cfg.EnableFile && level >= cfg.FileLevel {
		if s == "" {
			s = formatString(level, format, a...)
		}
		v := &FileLog{
			Level: level,
			Msg:   s,
		}
		select {
		case file.logChan <- v:
		default:
			l.Print(ERROR, "logChan full out")
		}
	}
}

func (l Logger) Print(level Level, format string, a ...interface{}) {
	if level < cfg.ConsoleLevel {
		return
	}
	s := formatString(level, format, a...)
	fmt.Printf(s)
}

func Debug(format string, a ...interface{}) {
	l.Log(DEBUG, format, a...)
}

func Info(format string, a ...interface{}) {
	l.Log(INFO, format, a...)
}

func Warn(format string, a ...interface{}) {
	l.Log(WARN, format, a...)
}

func Error(format string, a ...interface{}) {
	l.Log(ERROR, format, a...)
}

func Fatal(format string, a ...interface{}) {
	l.Log(FATAL, format, a...)
}

func getFuncInfo(skip int) (fileName, funcName string, line int) {
	pc, fileName, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(fileName)
	return
}

func formatString(level Level, format string, a ...interface{}) (s string) {
	var lv string
	switch level {
	case DEBUG:
		lv = "DEBUG"
	case INFO:
		lv = "INFO"
	case WARN:
		lv = "WARN"
	case ERROR:
		lv = "ERROR"
	case FATAL:
		lv = "FATAL"
	default:
		lv = "DEBUG"
	}
	fileName, funcName, line := getFuncInfo(4)
	msg := fmt.Sprintf(format, a...)
	return fmt.Sprintf("%s [%s] %s [%s %s %d]\n", time.Now().Format("2006-01-02 15:04:05"), lv, msg, fileName, funcName, line)
}
