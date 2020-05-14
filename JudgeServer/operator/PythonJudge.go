package operator

import (
	"github.com/afanke/OJO/JudgeServer/dto"
	"github.com/afanke/OJO/utils/log"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type PythonOperator struct {
}

func (py PythonOperator) Operate(form *dto.OperationForm) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("%v", err)
			form.Flag = "ISE"
			form.Score = 0
		}
		py.afterRun(form)
	}()
	err := py.beforeRun(form)
	if err != nil {
		log.Error("%v", err)
		return
	}
	err = py.run(form)
	if err != nil {
		log.Error("%v", err)
		return
	}
	err = py.judge(form)
	if err != nil {
		log.Error("%v", err)
		return
	}
	if form.UseSPJ {
		err = py.beforeSPJ(form)
		if err != nil {
			log.Error("%v", err)
			return
		}
		err = py.SPJ(form)
		if err != nil {
			log.Error("%v", err)
			return
		}
	}
}

func (py PythonOperator) beforeRun(form *dto.OperationForm) error {
	if form.CmdLine == "version" {
		form.CmdLine = "python3 python3 --version" + " " + strconv.Itoa(form.MaxRealTime) + strconv.Itoa(form.MaxCpuTime) + " " + strconv.Itoa(form.MaxMemory)
		return nil
	}
	path := form.Language + "_" + strconv.Itoa(rand.Int())
	form.CmdLine = "python3 python3 " + path + ".py " + path + "_input.txt " + strconv.Itoa(form.MaxRealTime) + " " + strconv.Itoa(form.MaxCpuTime) + " " + strconv.Itoa(form.MaxMemory)
	file, err := os.Create(path + ".py")
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
	file2, err := os.Create(path + "_input.txt")
	if err != nil {
		log.Error("%v", err)
		return err
	}
	defer file2.Close()
	_, err = file2.WriteString(form.Input)
	if err != nil {
		log.Error("%v", err)
		return err
	}
	return nil
}

func (py PythonOperator) run(form *dto.OperationForm) error {
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
	_, err = stdinPipe.Write([]byte("./SandBoxRunner " + form.CmdLine + "\n"))
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
	// fmt.Println(esr)
	if !strings.HasPrefix(esr, "^") {
		form.Flag = "ISE"
		form.RealOutput = "ISE"
		form.ErrorOutput = "ISE"
		form.ActualRealTime = 0
		form.ActualCpuTime = 0
		form.RealMemory = 0
		form.UseSPJ = false
	}
	form.ErrorOutput = esr
	// fmt.Println("esr", esr)
	form.RealOutput = res
	if len(form.RealOutput) >= 64000 {
		form.Flag = "OLE"
	}
	return nil
}

func (py PythonOperator) judge(form *dto.OperationForm) error {
	if form.Flag != "" {
		form.Score = 0
		return nil
	}
	esr := form.ErrorOutput
	res, err := strconv.Atoi(esr[strings.IndexByte(esr, 'c')+1 : strings.IndexByte(esr, 'r')])
	if err != nil {
		log.Error("%v", err)
		return err
	}
	form.ActualCpuTime = res
	res, err = strconv.Atoi(esr[strings.IndexByte(esr, 'r')+1 : strings.IndexByte(esr, '$')])
	if err != nil {
		log.Error("%v", err)
		return err
	}
	form.ActualRealTime = res
	res, err = strconv.Atoi(esr[strings.IndexByte(esr, 'm')+1 : strings.IndexByte(esr, 'c')])
	if err != nil {
		log.Error("%v", err)
		return err
	}
	form.RealMemory = res
	form.ErrorOutput = esr[strings.IndexByte(esr, '$')+1:]
	if esr[1:2] == "s" {
		sig := esr[2:strings.IndexByte(esr, 'm')]
		form.Score = 0
		switch sig {
		case "24":
			form.Flag = "TLE"
			if form.ActualCpuTime < form.MaxCpuTime*1000 {
				form.ActualCpuTime = form.MaxCpuTime
			}
		case "14":
			form.Flag = "TLE"
		case "11":
			form.Flag = "MLE"
			form.RealMemory = form.MaxMemory
		default:
			form.Flag = "RE"
		}
		form.UseSPJ = false
		return nil
	}
	form.ExpectOutput = strings.ReplaceAll(form.ExpectOutput, "\r\n", "\n")
	if form.ErrorOutput != "" {
		form.Score = 0
		if strings.Contains(form.ErrorOutput, "Traceback (") {
			form.Flag = "RE"
		} else if strings.Contains(form.ErrorOutput, "Error") {
			form.Flag = "CE"
		} else {
			form.Flag = "RE"
		}
		form.UseSPJ = false
		return nil
	}
	if form.UseSPJ {
		return nil
	}
	if form.ExpectOutput == form.RealOutput {
		form.Flag = "AC"
		return nil
	}
	form.Flag = "WA"
	form.Score = 0
	return nil
}

func (py PythonOperator) afterRun(form *dto.OperationForm) {
	_ = os.Remove(form.FilePath + ".py")
	_ = os.Remove(form.FilePath + "_input.txt")
	if form.UseSPJ && form.SPJPath != "" {
		_ = os.Remove(form.SPJPath + ".py")
		_ = os.Remove(form.SPJPath + "_input.txt")
		_ = os.Remove(form.SPJPath + "_expectOutput.txt")
		_ = os.Remove(form.SPJPath + "_realOutput.txt")
	}
}

func (py PythonOperator) beforeSPJ(form *dto.OperationForm) error {
	path := form.Language + "_" + strconv.Itoa(rand.Int())
	form.SPJPath = path
	form.CmdLine = "python3 " + path + ".py" + " " + path + "_input.txt" + " " + path + "_expectOutput.txt" + " " + path + "_realOutput.txt"
	file, err := os.Create(path + ".py")
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
	file2, err := os.Create(path + "_input.txt")
	if err != nil {
		log.Error("%v", err)
		return err
	}
	defer file2.Close()
	_, err = file2.WriteString(form.Input)
	if err != nil {
		log.Error("%v", err)
		return err
	}
	file3, err := os.Create(path + "_expectOutput.txt")
	if err != nil {
		log.Error("%v", err)
		return err
	}
	defer file3.Close()
	_, err = file3.WriteString(form.ExpectOutput)
	if err != nil {
		log.Error("%v", err)
		return err
	}
	file4, err := os.Create(path + "_realOutput.txt")
	if err != nil {
		log.Error("%v", err)
		return err
	}
	defer file4.Close()
	_, err = file4.WriteString(form.RealOutput)
	if err != nil {
		log.Error("%v", err)
		return err
	}
	return nil
}

func (py PythonOperator) SPJ(form *dto.OperationForm) error {
	cmd := exec.Command("/bin/bash")
	stdoutPipe, err := cmd.StdoutPipe()
	stdinPipe, err := cmd.StdinPipe()
	err = cmd.Start()
	if err != nil {
		log.Error("%v", err)
		return err
	}
	_, err = stdinPipe.Write([]byte(form.CmdLine))
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
	res := string(stdout)
	err = cmd.Wait()
	if err != nil {
		log.Error("%v", err)
		return err
	}
	switch res {
	case "AC":
		form.Flag = "AC"
	default:
		form.Flag = "WA"
		form.Score = 0
	}
	return nil
}
