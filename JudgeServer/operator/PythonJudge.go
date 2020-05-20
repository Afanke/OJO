package operator

import (
	"context"
	"fmt"
	"github.com/afanke/OJO/JudgeServer/dto"
	"github.com/afanke/OJO/utils/log"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type PythonOperator struct {
}

func (py PythonOperator) operate(form *dto.JudgeForm, i int) {
	ts := &dto.TempStorage{
		FilePath: "",
		SPJPath:  "",
		CmdLine:  "",
	}
	useSPJ := false
	defer func() {
		if err := recover(); err != nil {
			log.Error("%v", err)
			i := 0
			for {
				pc, fileName, line, ok := runtime.Caller(i)
				if !ok {
					form.TestCase[i].Flag = "ISE"
					form.TestCase[i].Score = 0
					return
				}
				funcName := runtime.FuncForPC(pc).Name()
				fileName = path.Base(fileName)
				fmt.Println(fileName, funcName, line)
				i++
			}
		}
		py.afterRun(useSPJ, ts)
	}()
	err := py.beforeRun(form, i, ts)
	if err != nil {
		form.TestCase[i].Flag = "ISE"
		log.Error("%v", err)
		return
	}
	err = py.run(form, i, ts)
	if err != nil {
		form.TestCase[i].Flag = "ISE"
		log.Error("%v", err)
		return
	}
	err = py.judge(form, i)
	if err != nil {
		form.TestCase[i].Flag = "ISE"
		log.Error("%v", err)
		return
	}
	useSPJ = form.UseSPJ && form.TestCase[i].Flag == "JUG"
	if useSPJ {
		err = py.beforeSPJ(form, i, ts)
		if err != nil {
			form.TestCase[i].Flag = "ISE"
			log.Error("%v", err)
			return
		}
		err = py.spj(form, i, ts)
		if err != nil {
			form.TestCase[i].Flag = "ISE"
			log.Error("%v", err)
			return
		}
	}
}

func (py PythonOperator) loopOperate(form *dto.JudgeForm) {
	for i := 0; i < len(form.TestCase); i++ {
		form.TestCase[i].Flag = "JUG"
		py.operate(form, i)
		log.Debug("i=%v", i)
	}
}

func (py PythonOperator) Mark(form *dto.JudgeForm) {
	py.loopOperate(form)
	py.summary(form)
}

func (py PythonOperator) summary(form *dto.JudgeForm) {
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

func (py PythonOperator) beforeRun(form *dto.JudgeForm, i int, ts *dto.TempStorage) error {
	if ts.CmdLine == "version" {
		ts.CmdLine = "python3 python3 --version" + " " + strconv.Itoa(form.MaxRealTime) + strconv.Itoa(form.MaxCpuTime) + " " + strconv.Itoa(form.MaxMemory)
		return nil
	}
	p := "Python3_" + strconv.Itoa(rand.Int())
	ts.FilePath = p
	ts.CmdLine = "python3 python3 " + p + ".py " + p + "_input.txt " + strconv.Itoa(form.MaxRealTime) + " " + strconv.Itoa(form.MaxCpuTime) + " " + strconv.Itoa(form.MaxMemory)
	file, err := os.Create(p + ".py")
	if err != nil {
		log.Error("%v", err)
		return err
	}
	defer file.Close()
	_, err = file.WriteString(form.Code)
	if err != nil {
		log.Error("%v", err)
		return err
	}
	file2, err := os.Create(p + "_input.txt")
	if err != nil {
		log.Error("%v", err)
		return err
	}
	defer file2.Close()
	_, err = file2.WriteString(form.TestCase[i].Input)
	if err != nil {
		log.Error("%v", err)
		return err
	}
	return nil
}

func (py PythonOperator) run(form *dto.JudgeForm, i int, ts *dto.TempStorage) error {
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
	// 	panic(err)
	// }
	_, err = stdinPipe.Write([]byte("./SandBoxRunner " + ts.CmdLine + "\n"))
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
		tc.RealOutput = "ISE"
		tc.ErrorOutput = "ISE"
		tc.ActualRealTime = 0
		tc.ActualCpuTime = 0
		tc.RealMemory = 0
	}
	tc.ErrorOutput = esr
	tc.RealOutput = res
	if len(tc.RealOutput) >= 64000 {
		tc.Flag = "OLE"
	}
	return nil
}

func (py PythonOperator) judge(form *dto.JudgeForm, i int) error {
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
	if esr[1:2] == "s" {
		sig := esr[2:strings.IndexByte(esr, 'm')]
		tc.Score = 0
		switch sig {
		case "24":
			tc.Flag = "TLE"
			if tc.ActualCpuTime < form.MaxCpuTime*1000 {
				tc.ActualCpuTime = form.MaxCpuTime
			}
		case "14":
			tc.Flag = "TLE"
		case "11":
			tc.Flag = "MLE"
			tc.RealMemory = form.MaxMemory
		default:
			tc.Flag = "RE"
		}
		return nil
	}
	tc.ExpectOutput = strings.ReplaceAll(tc.ExpectOutput, "\r\n", "\n")
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

func (py PythonOperator) afterRun(useSPJ bool, ts *dto.TempStorage) {
	_ = os.Remove(ts.FilePath + ".py")
	_ = os.Remove(ts.FilePath + "_input.txt")
	if useSPJ && ts.SPJPath != "" {
		_ = os.Remove(ts.SPJPath + ".py")
		_ = os.Remove(ts.SPJPath + "_input.txt")
		_ = os.Remove(ts.SPJPath + "_expectOutput.txt")
		_ = os.Remove(ts.SPJPath + "_realOutput.txt")
	}
}

func (py PythonOperator) beforeSPJ(form *dto.JudgeForm, i int, ts *dto.TempStorage) error {
	tc := &form.TestCase[i]
	p := "Python3_" + strconv.Itoa(rand.Int())
	ts.SPJPath = p
	ts.CmdLine = "python3 " + p + ".py" + " " + p + "_input.txt" + " " + p + "_expectOutput.txt" + " " + p + "_realOutput.txt"
	file, err := os.Create(p + ".py")
	if err != nil {
		log.Error("%v", err)
		return err
	}
	defer file.Close()
	_, err = file.WriteString(form.SPJCode)
	if err != nil {
		log.Error("%v", err)
		return err
	}
	file2, err := os.Create(p + "_input.txt")
	if err != nil {
		log.Error("%v", err)
		return err
	}
	defer file2.Close()
	_, err = file2.WriteString(tc.Input)
	if err != nil {
		log.Error("%v", err)
		return err
	}
	file3, err := os.Create(p + "_expectOutput.txt")
	if err != nil {
		log.Error("%v", err)
		return err
	}
	defer file3.Close()
	_, err = file3.WriteString(tc.ExpectOutput)
	if err != nil {
		log.Error("%v", err)
		return err
	}
	file4, err := os.Create(p + "_realOutput.txt")
	if err != nil {
		log.Error("%v", err)
		return err
	}
	defer file4.Close()
	_, err = file4.WriteString(tc.RealOutput)
	if err != nil {
		log.Error("%v", err)
		return err
	}
	return nil
}

func (py PythonOperator) spj(form *dto.JudgeForm, i int, ts *dto.TempStorage) error {
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
