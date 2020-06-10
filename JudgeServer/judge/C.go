package judge

import (
	"github.com/afanke/OJO/JudgeServer/dto"
	"github.com/afanke/OJO/utils/log"
	"math/rand"
	"os"
	"strconv"
)

type C struct{}

func (c C) getSPJCmpCmd(form *dto.JudgeForm, ts *dto.TempStorage) string {
	return ""
}

func (c C) getSPJRunCmd(form *dto.JudgeForm, ts *dto.TempStorage) string {
	p:=ts.SPJPath
	return "C " + p + ".py" + " " + p + "_input.txt" + " " + p + "_expectOutput.txt" + " " + p + "_realOutput.txt"
}

func (c C) writeSPJCode(form *dto.JudgeForm, ts *dto.TempStorage) error {
	p := "C_" + strconv.Itoa(rand.Int())
	ts.SPJPath = p
	file, err := os.Create(p + ".c")
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

func (c C) writeSPJInput(tc *dto.TestCase, ts *dto.TempStorage) error {
	p:=ts.SPJPath
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

func (c C) getCmpCmd(form *dto.JudgeForm,ts *dto.TempStorage) string {
	lmtStr := strconv.Itoa(form.MaxCpuTime) + " " + (strconv.Itoa(form.MaxRealTime)) + " " + strconv.Itoa(form.MaxMemory)
	inputPath := ts.FilePath + "_input.txt"
	return lmtStr + "gcc "+ts.
}

func (c C) getRunCmd(form *dto.JudgeForm,ts *dto.TempStorage) string {
	lmtStr := strconv.Itoa(form.MaxCpuTime) + " " + (strconv.Itoa(form.MaxRealTime)) + " " + strconv.Itoa(form.MaxMemory)
	inputPath := ts.FilePath + "_input.txt"
	return lmtStr + " " + inputPath + " ./" + ts.FilePath
}

func (c C) writeCode(form *dto.JudgeForm, ts *dto.TempStorage) error {
	p := "C_" + strconv.Itoa(rand.Int())
	ts.FilePath = p
	filePath := p + ".c"
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

func (c C) writeInput(form *dto.JudgeForm, i int, ts *dto.TempStorage) error {
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

func (c C) needCompile() bool {
	return true
}

func (c C) getSuffix() string {
	return ".c"
}

