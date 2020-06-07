package judge

import (
	"context"
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

type RtJug interface {
	needCompile() bool
	getSuffix() string
	beforeCompile(form *dto.JudgeForm, i int, ts *dto.TempStorage) error
	beforeRun(form *dto.JudgeForm, i int, ts *dto.TempStorage) error
}

type SpJug interface {
	needCompile() bool
	getSuffix() string
	beforeSPJ(form *dto.JudgeForm, i int, ts *dto.TempStorage) error
}

func NewJudge(rt RtJug, sp SpJug) Base {
	return Base{rt: rt, sp: sp}
}

func (b Base) Mark(form *dto.JudgeForm) {
	b.loopOperate(form)
	b.summary(form)
}

func (b Base) loopOperate(form *dto.JudgeForm) {
	for i := 0; i < len(form.TestCase); i++ {
		form.TestCase[i].Flag = "JUG"
		b.operate(form, i)
		log.Debug("i=%v", i)
	}
}

func (Base) summary(form *dto.JudgeForm) {
	form.TotalScore = 0
	for i := 0; i < len(form.TestCase); i++ {
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

func (b Base) operate(form *dto.JudgeForm, i int) {
	ts := &dto.TempStorage{
		FilePath: "",
		SPJPath:  "",
		CmdLine:  "",
	}
	defer func() {
		if err := recover(); err != nil {
			log.Error("%v", err)
			n := 0
			for {
				pc, fileName, line, ok := runtime.Caller(n)
				if !ok {
					form.TestCase[i].Flag = "ISE"
					form.TestCase[i].Score = 0
					return
				}
				funcName := runtime.FuncForPC(pc).Name()
				fileName = path.Base(fileName)
				fmt.Println(fileName, funcName, line)
				n++
			}
		}
		b.afterRun(ts)
	}()
	if b.rt.needCompile() {
		err := b.rt.beforeCompile(form, i, ts)
		if err != nil {
			form.TestCase[i].Flag = "ISE"
			form.TestCase[i].Score = 0
			log.Error("%v", err)
			return
		}
		err = b.compile(form, i, ts)
		if err != nil {
			form.TestCase[i].Flag = "ISE"
			form.TestCase[i].Score = 0
			log.Error("%v", err)
			return
		}
	}
	if form.TestCase[i].Flag != "JUG" {
		return
	}
	err := b.rt.beforeRun(form, i, ts)
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
	err = b.judge(form, i)
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
	err = b.sp.beforeSPJ(form, i, ts)
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

func (Base) compile(form *dto.JudgeForm, i int, ts *dto.TempStorage) error {
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
	_, err = stdinPipe.Write([]byte("./CPSBOX " + ts.CmdLine + "\n"))
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
	// fmt.Println(esr)
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
	r, err := strconv.Atoi(esr[strings.IndexByte(esr, 'c')+1 : strings.IndexByte(esr, 'r')])
	if err != nil {
		log.Error("%v", err)
		return err
	}
	tc.ActualCpuTime = r
	r, err = strconv.Atoi(esr[strings.IndexByte(esr, 'r')+1 : strings.IndexByte(esr, '$')])
	if err != nil {
		log.Error("%v", err)
		return err
	}
	tc.ActualRealTime = r
	r, err = strconv.Atoi(esr[strings.IndexByte(esr, 'm')+1 : strings.IndexByte(esr, 'c')])
	if err != nil {
		log.Error("%v", err)
		return err
	}
	tc.RealMemory = r
	tc.ErrorOutput = esr[strings.IndexByte(esr, '$')+1:]
	if len(res) >= 64000 {
		log.Error("output length:%v", len(res))
		tc.Flag = "OLE"
		tc.Score = 0
		tc.RealOutput = res
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
			tc.Flag = "CE"
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
		tc.Flag = "CE"
		return nil
	}
	return nil
}

func (Base) run(form *dto.JudgeForm, i int, ts *dto.TempStorage) error {
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
	_, err = stdinPipe.Write([]byte("./RTSBOX " + ts.CmdLine + "\n"))
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

func (Base) judge(form *dto.JudgeForm, i int) error {
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
		if strings.Contains(tc.ErrorOutput, "Traceback (") {
			tc.Flag = "RE"
		} else if strings.Contains(tc.ErrorOutput, "Error") {
			tc.Flag = "CE"
		} else {
			tc.Flag = "RE"
		}
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

func (Base) spj(form *dto.JudgeForm, i int, ts *dto.TempStorage) error {
	tc := &form.TestCase[i]
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(form.MaxRealTime))
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
	_, err = stdinPipe.Write([]byte(ts.CmdLine))
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

func (b Base) afterRun(ts *dto.TempStorage) {
	// if b.rt.needCompile() {
	// 	_ = os.Remove(ts.FilePath)
	// }
	// _ = os.Remove(ts.FilePath + b.rt.getSuffix())
	// _ = os.Remove(ts.FilePath + "_input.txt")
	if ts.UseSPJ && ts.SPJPath != "" {
		if b.sp.needCompile() {
			_ = os.Remove(ts.SPJPath)
		}
		_ = os.Remove(ts.SPJPath + b.sp.getSuffix())
		_ = os.Remove(ts.SPJPath + "_input.txt")
		_ = os.Remove(ts.SPJPath + "_expectOutput.txt")
		_ = os.Remove(ts.SPJPath + "_realOutput.txt")
	}
}
