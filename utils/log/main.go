package log

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
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

type Config struct {
	Level         Level
	EnableConsole bool
	EnableFile    bool
	FilePath      string
	MaxSize       int64
}

var cfg *Config

var file fileColl

type fileColl struct {
	info  *os.File
	warn  *os.File
	error *os.File
}

func init() {
	var defaultCfg = Config{
		Level:         0,
		EnableConsole: true,
		EnableFile:    false,
		FilePath:      "",
		MaxSize:       0,
	}
	bytes, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		Print(ERROR, err.Error())
		Print(ERROR, "use default log config")
		cfg = &defaultCfg
		return
	}
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		Print(ERROR, err.Error())
		Print(ERROR, "use default log config")
		cfg = &defaultCfg
		return
	}
	fileInit()
	Print(INFO, "success to load log system")
}

func (f fileColl) Info(msg string) {
	if f.info == nil {
		Print(ERROR, "can't open file:%s/info.log", cfg.FilePath)
		return
	}
	_, err := f.info.WriteString(msg)
	if err != nil {
		Print(ERROR, "%v", err)
		return
	}

	// check(f.info)
}

func (f fileColl) Warn(msg string) {
	if f.warn == nil {
		Print(ERROR, "can't open file:%s/warn.log", cfg.FilePath)
		return
	}
	_, err := f.warn.WriteString(msg)
	if err != nil {
		Print(ERROR, "%v", err)
		return
	}
	// check(f.warn)
}

func (f fileColl) Error(msg string) {
	if f.error == nil {
		Print(ERROR, "can't open file:%s/error.log", cfg.FilePath)
		return
	}
	_, err := f.error.WriteString(msg)
	if err != nil {
		Print(ERROR, "%v", err)
		return
	}
	// check(f.error)
}

func check(f *os.File) {
	stat, err := f.Stat()
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	if stat.Size() < cfg.MaxSize {
		return
	}
	sp := strings.Split(f.Name(), "_")
	itoa, err := strconv.Atoi(sp[1])
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	itoa++
	switch sp[0] {
	case "info":
		openFile, err := os.OpenFile(cfg.FilePath+"/"+sp[0]+strconv.Itoa(itoa), os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("error:%v\n", err)
			return
		}
		file.info = openFile
	case "warn":
		openFile, err := os.OpenFile(cfg.FilePath+"/"+sp[0]+strconv.Itoa(itoa), os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("error:%v\n", err)
			return
		}
		file.warn = openFile
	case "error":
		openFile, err := os.OpenFile(cfg.FilePath+"/"+sp[0]+strconv.Itoa(itoa), os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("error:%v\n", err)
			return
		}
		file.error = openFile
	default:
		fmt.Println("no matched pattern")
		return
	}
	_ = f.Close()
}

func fileInit() {
	if !cfg.EnableFile {
		Print(INFO, "no log file loaded")
		return
	}
	info, err := os.OpenFile(cfg.FilePath+"/info.log", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		Print(FATAL, "%v", err)
	}
	warn, err := os.OpenFile(cfg.FilePath+"/warn.log", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		Print(FATAL, "%v", err)
	}
	erro, err := os.OpenFile(cfg.FilePath+"/error.log", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		Print(FATAL, "%v", err)
	}
	file.info = info
	file.warn = warn
	file.error = erro
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

	// if level < ERROR {
	// 	return fmt.Sprintf("%s [%s] %s [%s %s %d]\n", time.Now().Format("2006-01-02 15:04:05"), lv, msg, fileName, funcName, line)
	// } else {
	// 	return fmt.Sprintf("%s [%s] [%s %s %d] %s \n", time.Now().Format("2006-01-02 15:04:05"), lv, fileName, funcName, line, msg)
	// }
}

func Log(level Level, format string, a ...interface{}) {
	if cfg.Level > level {
		return
	}
	s := formatString(level, format, a...)
	if cfg.EnableConsole {
		if level < ERROR {
			fmt.Printf(s)
		} else {
			print(s)
		}
	}
	if cfg.EnableFile {
		switch level {
		case INFO:
			file.Info(s)
		case WARN:
			file.Warn(s)
		case ERROR:
			file.Error(s)
		case FATAL:
			openFile, err := os.OpenFile("./logs/fatal.log", os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				_ = fmt.Errorf("error:%v\n", err)
				break
			}
			_, err = openFile.WriteString(s)
			if err != nil {
				_ = fmt.Errorf("error:%v\n", err)
				break
			}
		}
	}
	if level == FATAL {
		os.Exit(1)
	}
}

func Print(level Level, format string, a ...interface{}) {
	s := formatString(level, format, a...)
	if level < ERROR {
		fmt.Printf(s)
	} else {
		print(s)
	}
	if level == FATAL {
		os.Exit(1)
	}
}

func Debug(format string, a ...interface{}) {
	Log(DEBUG, format, a...)
}
func Info(format string, a ...interface{}) {
	Log(INFO, format, a...)
}
func Warn(format string, a ...interface{}) {
	Log(WARN, format, a...)
}
func Error(format string, a ...interface{}) {
	Log(ERROR, format, a...)
}
func Fatal(format string, a ...interface{}) {
	Log(FATAL, format, a...)
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
