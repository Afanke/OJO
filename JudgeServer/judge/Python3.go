package judge

import (
	"github.com/afanke/OJO/JudgeServer/dto"
	"github.com/afanke/OJO/utils/log"
	"math/rand"
	"os"
	"strconv"
)

type Python struct{}

func (py Python) beforeCompile(form *dto.JudgeForm, i int, ts *dto.TempStorage) error {
	// do nothing
	return nil
}

func (py Python) needCompile() bool {
	return false
}

func (py Python) getSuffix() string {
	return ".py"
}

func (py Python) beforeRun(form *dto.JudgeForm, i int, ts *dto.TempStorage) error {
	p := "Python3_" + strconv.Itoa(rand.Int())
	ts.FilePath = p
	lmtStr := strconv.Itoa(form.MaxCpuTime) + " " + (strconv.Itoa(form.MaxRealTime)) + " " + strconv.Itoa(form.MaxMemory)
	inputPath := p + "_input.txt"
	filePath := p + ".py"
	ts.CmdLine = lmtStr + " " + inputPath + " python3 " + filePath
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
	file2, err := os.Create(inputPath)
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

func (py Python) beforeSPJ(form *dto.JudgeForm, i int, ts *dto.TempStorage) error {
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
