package judge

import (
	"github.com/afanke/OJO/JudgeServer/dto"
	"github.com/afanke/OJO/utils/log"
	"math/rand"
	"os"
	"strconv"
)

type Python3 struct{}

func (py Python3) getSPJCmpCmd(form *dto.JudgeForm, ts *dto.TempStorage) string {
	return ""
}

func (py Python3) getSPJRunCmd(form *dto.JudgeForm, ts *dto.TempStorage) string {
	p := ts.SPJPath
	return "python3 " + p + ".py" + " " + p + "_input.txt" + " " + p + "_expectOutput.txt" + " " + p + "_realOutput.txt"
}

func (py Python3) writeSPJCode(form *dto.JudgeForm, ts *dto.TempStorage) error {
	p := "Python3_" + strconv.Itoa(rand.Int())
	ts.SPJPath = p
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
	return nil
}

func (py Python3) writeSPJInput(tc *dto.TestCase, ts *dto.TempStorage) error {
	p := ts.SPJPath
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

func (py Python3) getCmpCmd(form *dto.JudgeForm, ts *dto.TempStorage) string {
	return ""
}

func (py Python3) getRunCmd(form *dto.JudgeForm, ts *dto.TempStorage) string {
	lmtStr := strconv.Itoa(form.MaxCpuTime) + " " + (strconv.Itoa(form.MaxRealTime)) + " " + strconv.Itoa(form.MaxMemory)
	inputPath := ts.FilePath + "_input.txt"
	filePath := ts.FilePath + ".py"
	return lmtStr + " " + inputPath + " python3 " + filePath
}

func (py Python3) writeCode(form *dto.JudgeForm, ts *dto.TempStorage) error {
	p := "Python3_" + strconv.Itoa(rand.Int())
	ts.FilePath = p
	filePath := p + ".py"
	file, err := os.Create(filePath)
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
	return nil
}

func (py Python3) writeInput(form *dto.JudgeForm, i int, ts *dto.TempStorage) error {
	inputPath := ts.FilePath + "_input.txt"
	file, err := os.Create(inputPath)
	if err != nil {
		log.Error("%v", err)
		return err
	}
	defer file.Close()
	_, err = file.WriteString(form.TestCase[i].Input)
	if err != nil {
		log.Error("%v", err)
		return err
	}
	return nil
}

func (py Python3) needCompile() bool {
	return false
}

func (py Python3) getSuffix() string {
	return ".py"
}
