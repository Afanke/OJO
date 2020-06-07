package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/afanke/OJO/JudgeServer/dto"
	"github.com/afanke/OJO/utils/log"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	form := dto.JudgeForm{
		UseSPJ:      false,
		MaxCpuTime:  1000,
		MaxRealTime: 1000,
		MaxMemory:   1000000,
		TotalScore:  0,
		Id:          0,
		Lid:         1,
		SPJLid:      1,
		SPJCode:     ``,
		Code: `#include <stdio.h>
int main(){
int a,b;
scanf("%d %d",&a,&b);
printf("%d",a+b);
}`,
		// Code:`print("Hello World")`,
		Flag: "",
		TestCase: []dto.TestCase{
			{
				Input:        "1 2",
				ExpectOutput: "3",
				Score:        10,
			},
			{
				Input:        "4 5",
				ExpectOutput: "9",
				Score:        10,
			},
		},
	}
	client := &http.Client{
		Timeout: time.Duration(form.MaxRealTime) * time.Duration(len(form.TestCase)) * time.Second * 2,
	}
	buff, err := json.Marshal(&form)
	if err != nil {
		log.Error("error:%v\n", err)
		return
	}
	// res, err := client.Post("http://192.168.111.135:2333/Python3", "application/json", bytes.NewBuffer(buff))
	res, err := client.Post("http://192.168.111.138:2333/judge", "application/json", bytes.NewBuffer(buff))
	// res, err := client.Post("http://49.234.91.99:2333/Python3", "application/json", bytes.NewBuffer(buff))
	if err != nil {
		log.Error("error:%v\n", err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("error:%v\n", err)
		return
	}
	err = json.Unmarshal(body, &form)
	if err != nil {
		log.Error("error:%v\n", err)
		return
	}
	fmt.Printf("%+v", form)
}
