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
	form := dto.JudgeForm{
		UseSPJ:      false,
		MaxCpuTime:  1000,
		MaxRealTime: 1000,
		MaxMemory:   10000,
		TotalScore:  0,
		Id:          0,
		SPJCode: `import sys
f1=open(sys.argv[1],'r').read()
f2=open(sys.argv[2],'r').read()
f3=open(sys.argv[3],'r').read()
if f1==f3:
	print("AC",end="")`,
		Code: `
while True:
    print("1000"*1000)`,
		Flag: "",
		TestCase: []dto.TestCase{
			{
				Input:        "qwe",
				ExpectOutput: "qwe",
				Score:        10,
			},
			{
				Input:        "123",
				ExpectOutput: "123",
				Score:        10,
			},
			{
				Input:        "1241",
				ExpectOutput: "1241",
				Score:        10,
			},
		},
	}
	client := &http.Client{
		Timeout: time.Duration(form.MaxRealTime) * time.Duration(len(form.TestCase)) * time.Second * 2,
	}
	buff, err := json.Marshal(form)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	// res, err := client.Post("http://192.168.111.135:2333/Python3", "application/json", bytes.NewBuffer(buff))
	res, err := client.Post("http://49.234.91.99:2333/Python3", "application/json", bytes.NewBuffer(buff))
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
	// fmt.Println(string(body))
	err = json.Unmarshal(body, &form)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	fmt.Printf("%+v", form)
}
