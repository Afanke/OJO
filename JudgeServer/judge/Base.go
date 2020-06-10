package judge

import (
	"context"
	"errors"
	"fmt"
	"github.com/afanke/OJO/JudgeServer/dto"
	"github.com/afanke/OJO/utils/log"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Base struct {
	rt RtJug
	sp SpJug
}

const (
	INPUT_SUFFIX           = "_input.txt"
	EXPECTED_OUTPUT_SUFFIX = "_expectedOutput.txt"
	INPUT_SUFFIX           = "_input.txt"
)

type RtJug interface {
	needCompile() bool
	getSuffix() string
	getCmpCmd(form *dto.JudgeForm, ts *dto.TempStorage) string
	getRunCmd(form *dto.JudgeForm, ts *dto.TempStorage) string
	writeCode(form *dto.JudgeForm, ts *dto.TempStorage) error
	writeInput(form *dto.JudgeForm, i int, ts *dto.TempStorage) error
}

type SpJug interface {
	needCompile() bool
	getSuffix() string
	getSPJCmpCmd(form *dto.JudgeForm, ts *dto.TempStorage) string
	getSPJRunCmd(form *dto.JudgeForm, ts *dto.TempStorage) string
	writeSPJCode(form *dto.JudgeForm, ts *dto.TempStorage) error
	writeSPJInput(form *dto.JudgeForm, i int, ts *dto.TempStorage) error
}

func NewJudge(rt RtJug, sp SpJug) Base {
	return Base{rt: rt, sp: sp}
}

func (b Base) Judge(form *dto.JudgeForm) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("%v", err)
			n := 0

			if s, ok := err.(string); ok {
				form.ErrorMsg = "Internal Server Error:" + s + "\n"
			} else {
				form.ErrorMsg = "Internal Server Error:\n"
			}
			for {
				pc, fileName, line, ok := runtime.Caller(n)
				if !ok {
					form.Flag = "ISE"
					form.TestCase = nil
					return
				}
				funcName := runtime.FuncForPC(pc).Name()
				fileName = path.Base(fileName)
				fmt.Println(fileName, funcName, line)
				form.ErrorMsg = fileName + " " + funcName + " " + strconv.Itoa(line) + "\n"
				n++
			}
		}
	}()
	ts := &dto.TempStorage{}
	form.Flag = "JUG"
	err := b.rt.writeCode(form, ts)
	defer b.cleanCodeFile(ts)
	if err != nil {
		form.Flag = "ISE"
		form.ErrorMsg = "Internal Server Error: " + err.Error() + "\n"
		form.TestCase = nil
		log.Error("%v", err)
		return
	}
	if b.rt.needCompile() {
		err = b.compile(form, ts)
		if err != nil {
			form.ErrorMsg = err.Error()
			form.TestCase = nil
			log.Error("%v", err)
			return
		}
	}
	if form.UseSPJ {
		err := b.sp.writeSPJCode(form, ts)
		defer b.cleanSPJFile(ts)
		if err != nil {
			form.Flag = "ISE"
			form.ErrorMsg = "Internal Server Error: " + err.Error() + "\n"
			form.TestCase = nil
			log.Error("%v", err)
			return
		}
		if b.sp.needCompile() {
			err = b.compileSPJ(form, ts)
			if err != nil {
				form.ErrorMsg = "SPJ " + err.Error()
				form.TestCase = nil
				log.Error("%v", err)
				return
			}
		}
	}
	for i := 0; i < len(form.TestCase); i++ {
		form.TestCase[i].Flag = "JUG"
		b.judgeTestCase(form, i, ts)
		log.Debug("i=%v", i)
	}
	b.summary(form)
}

func (Base) summary(form *dto.JudgeForm) {
	l := len(form.TestCase)
	if l == 0 {
		return
	}
	form.TotalScore = 0
	for i := 0; i < l; i++ {
		form.TotalScore += form.TestCase[i].Score
	}
	hasAC := false
	hasWA := false
	for i := 0; i < len(form.TestCase); i++ {
		flag := form.TestCase[i].Flag
		if flag == "AC" {
			hasAC = true
			continue
		}
		switch flag {
		case "WA":
			hasWA = true
			continue
		default:
			form.Flag = flag
			return
		}
	}
	if hasAC && !hasWA {
		form.Flag = "AC"
		return
	}
	if hasAC && hasWA {
		form.Flag = "PA"
		return
	}
	form.Flag = "WA"
}

func (b Base) judgeTestCase(form *dto.JudgeForm, i int, ts *dto.TempStorage) {
	err := b.rt.writeInput(form, i, ts)
	defer b.cleanInputFile(ts)
	if err != nil {
		form.TestCase[i].Flag = "ISE"
		form.TestCase[i].Score = 0
		log.Error("%v", err)
		return
	}
	err = b.run(form, i, ts)
	if err != nil {
		form.TestCase[i].Flag = "ISE"
		form.TestCase[i].Score = 0
		log.Error("%v", err)
		return
	}
	err = b.concludeFlag(form, i)
	if err != nil {
		form.TestCase[i].Flag = "ISE"
		form.TestCase[i].Score = 0
		log.Error("%v", err)
		return
	}
	ts.UseSPJ = form.UseSPJ && form.TestCase[i].Flag == "JUG"
	if !ts.UseSPJ {
		return
	}
	err = b.sp.writeSPJInput(form, i, ts)
	if err != nil {
		form.TestCase[i].Flag = "ISE"
		form.TestCase[i].Score = 0
		log.Error("%v", err)
		return
	}
	err = b.spj(form, i, ts)
	if err != nil {
		form.TestCase[i].Flag = "ISE"
		form.TestCase[i].Score = 0
		log.Error("%v", err)
		return
	}

}

func (b Base) compile(form *dto.JudgeForm, ts *dto.TempStorage) error {
	cmd := exec.Command("/bin/bash")
	stdoutPipe, err := cmd.StdoutPipe()
	stderrPipe, err := cmd.StderrPipe()
	stdinPipe, err := cmd.StdinPipe()
	err = cmd.Start()
	if err != nil {
		log.Error("%v", err)
		form.Flag = "ISE"
		return errors.New("Internal Server Error: " + err.Error() + "\n")
	}
	// _, err = stdinPipe.Write([]byte("su judge\n"))
	// if err != nil {
	// 	log.Error("%v", err)
	// 	form.Flag = "ISE"
	// 	return errors.New("Internal Server Error: " + err.Error() + "\n")
	// }
	_, err = stdinPipe.Write([]byte("./CPSBOX " + b.rt.getCmpCmd(form, ts) + "\n"))
	if err != nil {
		log.Error("%v", err)
		form.Flag = "ISE"
		return errors.New("Internal Server Error: " + err.Error() + "\n")
	}
	err = stdinPipe.Close()
	if err != nil {
		log.Error("%v", err)
		form.Flag = "ISE"
		return errors.New("Internal Server Error: " + err.Error() + "\n")
	}
	stderr, err := ioutil.ReadAll(stderrPipe)
	if err != nil {
		log.Error("%v", err)
		form.Flag = "ISE"
		return errors.New("Internal Server Error: " + err.Error() + "\n")
	}
	stdout, err := ioutil.ReadAll(stdoutPipe)
	if err != nil {
		log.Error("%v", err)
		form.Flag = "ISE"
		return errors.New("Internal Server Error: " + err.Error() + "\n")
	}
	res := string(stdout)
	esr := string(stderr)
	err = cmd.Wait()
	log.Debug("compile esr:%v", esr)
	if !strings.HasPrefix(esr, "^") {
		form.Flag = "ISE"
		log.Error("no prefix ^")
		msg := "Internal Server Error 1: Can't get runtime message\n"
		return errors.New(msg)
	}
	r, err := strconv.Atoi(esr[strings.IndexByte(esr, 'c')+1 : strings.IndexByte(esr, 'r')])
	if err != nil {
		log.Error("%v", err)
		form.Flag = "ISE"
		return errors.New("Internal Server Error: " + err.Error() + "\n")
	}
	actualCpuTime := r
	r, err = strconv.Atoi(esr[strings.IndexByte(esr, 'r')+1 : strings.IndexByte(esr, '$')])
	if err != nil {
		log.Error("%v", err)
		form.Flag = "ISE"
		return errors.New("Internal Server Error: " + err.Error() + "\n")
	}
	actualRealTime := r
	r, err = strconv.Atoi(esr[strings.IndexByte(esr, 'm')+1 : strings.IndexByte(esr, 'c')])
	if err != nil {
		log.Error("%v", err)
		form.Flag = "ISE"
		return errors.New("Internal Server Error: " + err.Error() + "\n")
	}
	realMemory := r
	errorOutput := esr[strings.IndexByte(esr, '$')+1:]
	lr := len(res)
	if lr >= 64000 {
		log.Error("output length:%v", len(res))
		form.Flag = "CE"
		msg := "Compile Error : Out of output limit" + "\n" +
			"Output Length:" + strconv.Itoa(lr) + "\n"
		return errors.New(msg)
	}
	if esr[1:2] == "s" {
		sig := esr[2:strings.IndexByte(esr, 'm')]
		form.Flag = "CE"
		msg := ""
		switch sig {
		case "14":
			fallthrough
		case "24":
			log.Debug("sig 24 tle")
			msg = "Compile Error: Out of time limit\n" +
				"CPU Time:" + strconv.Itoa(actualCpuTime) + "\n" +
				"Real Time:" + strconv.Itoa(actualRealTime) + "\n"
		case "11":
			log.Debug("sig 11 mle")
			msg = "Compile Error: Out of memory limit\n" +
				"Memory Used:" + strconv.Itoa(realMemory) + "\n"
		default:
			log.Debug("sig " + sig + " mle")
			msg = "Compile Error: Interrupted by system signal when compiling program\n" +
				"Signal Received:" + sig + "\n"
		}
		return errors.New(msg)
	}
	// tc.ExpectOutput = strings.ReplaceAll(tc.ExpectOutput, "\r\n", "\n")
	if actualCpuTime > form.MaxCpuTime || actualRealTime > form.MaxRealTime {
		form.Flag = "CE"
		msg := "Compile Error: Out of time limit\n" +
			"CPU Time:" + strconv.Itoa(actualCpuTime) + "\n" +
			"Real Time:" + strconv.Itoa(actualRealTime) + "\n"
		return errors.New(msg)
	}
	if realMemory > form.MaxMemory {
		form.Flag = "CE"
		msg := "Compile Error: Out of memory limit\n" +
			"Memory Used:" + strconv.Itoa(realMemory) + "\n"
		return errors.New(msg)
	}
	if errorOutput != "" {
		form.Flag = "CE"
		msg := "Compile Error:\n" + errorOutput + "\n"
		return errors.New(msg)
	}
	return nil
}

func (b Base) run(form *dto.JudgeForm, i int, ts *dto.TempStorage) error {
	cmd := exec.Command("/bin/bash")
	stdoutPipe, err := cmd.StdoutPipe()
	stderrPipe, err := cmd.StderrPipe()
	stdinPipe, err := cmd.StdinPipe()
	err = cmd.Start()
	if err != nil {
		log.Error("%v", err)
		return err
	}
	// _, err = stdinPipe.Write([]byte("su judge\n"))
	// if err != nil {
	// 	log.Error("%v", err)
	// 	return err
	// }
	_, err = stdinPipe.Write([]byte("./RTSBOX " + b.rt.getRunCmd(form, ts) + "\n"))
	if err != nil {
		log.Error("%v", err)
		return err
	}
	err = stdinPipe.Close()
	if err != nil {
		log.Error("%v", err)
		return err
	}
	stderr, err := ioutil.ReadAll(stderrPipe)
	if err != nil {
		log.Error("%v", err)
		return err
	}
	stdout, err := ioutil.ReadAll(stdoutPipe)
	if err != nil {
		log.Error("%v", err)
		return err
	}
	res := string(stdout)
	esr := string(stderr)
	err = cmd.Wait()
	tc := &form.TestCase[i]
	log.Debug("res:%v", res)
	log.Debug("esr:%v", esr)
	if !strings.HasPrefix(esr, "^") {
		tc.Flag = "ISE"
		tc.RealOutput = res
		tc.ErrorOutput = "ISE"
		tc.ActualRealTime = 0
		tc.ActualCpuTime = 0
		tc.RealMemory = 0
		tc.Score = 0
		log.Error("no prefix ^")
		log.Error("%v", esr)
		return nil
	}
	tc.ErrorOutput = esr
	tc.RealOutput = res
	return nil
}

func (Base) concludeFlag(form *dto.JudgeForm, i int) error {
	tc := &form.TestCase[i]
	if tc.Flag != "JUG" {
		tc.Score = 0
		return nil
	}
	esr := tc.ErrorOutput
	log.Debug("%+v", esr)
	res, err := strconv.Atoi(esr[strings.IndexByte(esr, 'c')+1 : strings.IndexByte(esr, 'r')])
	if err != nil {
		log.Error("%v", err)
		return err
	}
	tc.ActualCpuTime = res
	res, err = strconv.Atoi(esr[strings.IndexByte(esr, 'r')+1 : strings.IndexByte(esr, '$')])
	if err != nil {
		log.Error("%v", err)
		return err
	}
	tc.ActualRealTime = res
	res, err = strconv.Atoi(esr[strings.IndexByte(esr, 'm')+1 : strings.IndexByte(esr, 'c')])
	if err != nil {
		log.Error("%v", err)
		return err
	}
	tc.RealMemory = res
	tc.ErrorOutput = esr[strings.IndexByte(esr, '$')+1:]
	if len(tc.RealOutput) >= 64000 {
		log.Debug("output length:%v", len(tc.RealOutput))
		tc.Flag = "OLE"
		tc.Score = 0
		return nil
	}
	if esr[1:2] == "s" {
		sig := esr[2:strings.IndexByte(esr, 'm')]
		tc.Score = 0
		switch sig {
		case "24":
			tc.Flag = "TLE"
			log.Debug("sig 24 tle")
			if tc.ActualCpuTime < form.MaxCpuTime {
				tc.ActualCpuTime = form.MaxCpuTime
			}
		case "14":
			log.Debug("sig 14 tle")
			tc.Flag = "TLE"
		case "11":
			log.Debug("sig 11 mle")
			tc.Flag = "MLE"
			tc.RealMemory = form.MaxMemory
		case "31":
			log.Debug("sig 31 mle")
			tc.Flag = "MLE"
			// tc.RealMemory = form.MaxMemory
		default:
			tc.Flag = "RE"
		}
		return nil
	}
	// tc.ExpectOutput = strings.ReplaceAll(tc.ExpectOutput, "\r\n", "\n")
	if tc.ActualCpuTime > form.MaxCpuTime {
		tc.Flag = "TLE"
		tc.Score = 0
		return nil
	}
	if tc.ActualRealTime > form.MaxRealTime {
		tc.Flag = "TLE"
		tc.Score = 0
		return nil
	}
	if tc.RealMemory > form.MaxMemory {
		tc.Flag = "MLE"
		tc.Score = 0
		return nil
	}
	if tc.ErrorOutput != "" {
		tc.Score = 0
		tc.Flag = "RE"
		return nil
	}
	if form.UseSPJ {
		return nil
	}
	if tc.ExpectOutput == tc.RealOutput {
		tc.Flag = "AC"
		return nil
	}
	tc.Flag = "WA"
	tc.Score = 0
	return nil
}

func (b Base) compileSPJ(form *dto.JudgeForm, ts *dto.TempStorage) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(form.MaxRealTime)*5)
	defer cancel()
	cmd := exec.CommandContext(ctx, "/bin/bash")
	// stdoutPipe, err := cmd.StdoutPipe()
	stderrPipe, err := cmd.StderrPipe()
	stdinPipe, err := cmd.StdinPipe()
	err = cmd.Start()
	if err != nil {
		log.Error("%v", err)
		form.Flag = "ISE"
		return errors.New("Internal Server Error: " + err.Error() + "\n")
	}
	// _, err = stdinPipe.Write([]byte("su judge\n"))
	// if err != nil {
	// 	log.Error("%v", err)
	// 	form.Flag = "ISE"
	// 	return errors.New("Internal Server Error: " + err.Error() + "\n")
	// }
	_, err = stdinPipe.Write([]byte(b.sp.getSPJCmpCmd(form, ts) + "\n"))
	if err != nil {
		log.Error("%v", err)
		form.Flag = "ISE"
		return errors.New("Internal Server Error: " + err.Error() + "\n")
	}
	err = stdinPipe.Close()
	if err != nil {
		log.Error("%v", err)
		form.Flag = "ISE"
		return errors.New("Internal Server Error: " + err.Error() + "\n")
	}
	stderr, err := ioutil.ReadAll(stderrPipe)
	if err != nil {
		log.Error("%v", err)
		form.Flag = "ISE"
		return errors.New("Internal Server Error: " + err.Error() + "\n")
	}
	// stdout, err := ioutil.ReadAll(stdoutPipe)
	// if err != nil {
	// 	log.Error("%v", err)
	// 	form.Flag = "ISE"
	// 	return errors.New("Internal Server Error: " + err.Error() + "\n")
	// }
	// res := string(stdout)
	esr := string(stderr)
	err = cmd.Wait()
	if err != nil {
		log.Error("%v", err)
		form.Flag = "ISE"
		return errors.New("Internal Server Error: " + err.Error() + "\n")
	}
	if esr != "" {
		form.Flag = "CE"
		return errors.New("SPJ Compile Error:\n" + esr + "\n")
	}
	return nil
}

func (b Base) spj(form *dto.JudgeForm, i int, ts *dto.TempStorage) error {
	tc := &form.TestCase[i]
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(form.MaxRealTime)*5)
	defer cancel()
	cmd := exec.CommandContext(ctx, "/bin/bash")
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Error("%v", err)
		return err
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		log.Error("%v", err)
		return err
	}
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		log.Error("%v", err)
		return err
	}
	err = cmd.Start()
	if err != nil {
		log.Error("%v", err)
		return err
	}
	_, err = stdinPipe.Write([]byte(b.sp.getSPJRunCmd(form, ts)))
	if err != nil {
		log.Error("%v", err)
		return err
	}
	err = stdinPipe.Close()
	if err != nil {
		log.Error("%v", err)
		return err
	}
	stdout, err := ioutil.ReadAll(stdoutPipe)
	if err != nil {
		log.Error("%v", err)
		return err
	}
	stderr, err := ioutil.ReadAll(stderrPipe)
	if err != nil {
		log.Error("%v", err)
		return err
	}
	tc.SPJOutput = string(stdout)
	tc.SPJErrorOutput = string(stderr)
	err = cmd.Wait()
	if err != nil {
		log.Error("%v", err)
		return err
	}
	switch tc.SPJOutput {
	case "AC":
		tc.Flag = "AC"
	default:
		tc.Flag = "WA"
		tc.Score = 0
	}
	return nil
}

func (b Base) cleanInputFile(ts *dto.TempStorage) {
	_ = os.Remove(ts.FilePath + "_input.txt")
	if ts.UseSPJ && ts.SPJPath != "" {
		_ = os.Remove(ts.SPJPath + "_input.txt")
		_ = os.Remove(ts.SPJPath + "_expectOutput.txt")
		_ = os.Remove(ts.SPJPath + "_realOutput.txt")
	}
}

func (b Base) cleanCodeFile(ts *dto.TempStorage) {
	if b.rt.needCompile() {
		_ = os.Remove(ts.FilePath)
	}
	_ = os.Remove(ts.FilePath + b.rt.getSuffix())
}

func (b Base) cleanSPJFile(ts *dto.TempStorage) {
	if b.sp.needCompile() {
		_ = os.Remove(ts.SPJPath)
	}
	_ = os.Remove(ts.SPJPath + b.sp.getSuffix())
}

func getLmtStr(form *dto.JudgeForm) string {
	return strconv.Itoa(form.MaxCpuTime) + " " + (strconv.Itoa(form.MaxRealTime)) + " " + strconv.Itoa(form.MaxMemory)
}
