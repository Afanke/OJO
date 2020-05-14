package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/afanke/OJO/JudgeServer/dto"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	var forms = make([]dto.OperationForm, 0, 1)
	form1 := dto.OperationForm{
		Language:     "Python3",
		FilePath:     "",
		CmdLine:      "",
		Code:         `print(input())`,
		Input:        "321",
		ExpectOutput: "321\n",
		RealOutput:   "",
		ErrorOutput:  "",
		Flag:         "",
		MaxCpuTime:   2,
		MaxRealTime:  2,
		MaxMemory:    20000000,
		Score:        10,
		UseSPJ:       true,
		SPJCode: `import sys
f1=open(sys.argv[1], mode='r')
f2=open(sys.argv[3],mode='r')
if f1.read()+'\n'==f2.read():
	print("AC",end="")`,
	}
	forms = append(forms, form1)
	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	buff, err := json.Marshal(forms)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	res, err := client.Post("http://192.168.111.132:2333/"+forms[0].Language, "application/json", bytes.NewBuffer(buff))
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	err = json.Unmarshal(body, &forms)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	fmt.Printf("%#v", forms)
}
