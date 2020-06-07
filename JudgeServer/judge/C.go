package judge

import (
	"github.com/afanke/OJO/JudgeServer/dto"
	"github.com/afanke/OJO/utils/log"
	"math/rand"
	"os"
	"strconv"
)

type C struct{}

func (c C) beforeCompile(form *dto.JudgeForm, i int, ts *dto.TempStorage) error {
	p := "C_" + strconv.Itoa(rand.Int())
	ts.FilePath = p
	lmtStr := strconv.Itoa(form.MaxCpuTime) + " " + (strconv.Itoa(form.MaxRealTime)) + " " + strconv.Itoa(form.MaxMemory)
	filePath := p + ".c"
	ts.CmdLine = lmtStr + " 0 gcc " + filePath + " -o " + p
	log.Debug("%v", ts.CmdLine)
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

func (c C) needCompile() bool {
	return true
}

func (c C) beforeRun(form *dto.JudgeForm, i int, ts *dto.TempStorage) error {
	lmtStr := strconv.Itoa(form.MaxCpuTime) + " " + (strconv.Itoa(form.MaxRealTime)) + " " + strconv.Itoa(form.MaxMemory)
	inputPath := ts.FilePath + "_input.txt"
	ts.CmdLine = lmtStr + " " + inputPath + " ./" + ts.FilePath
	log.Debug("%v", ts.CmdLine)
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

func (c C) beforeSPJ(form *dto.JudgeForm, i int, ts *dto.TempStorage) error {
	panic("implement me")
}

func (c C) getSuffix() string {
	return ".c"
}
