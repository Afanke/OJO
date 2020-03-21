package main

import (
	"encoding/json"
	"fmt"
	"github.com/afanke/OJO/JudgeServer/dto"
	tcp "github.com/afanke/OJO/utils/tcp"
)

func main() {
	conn, err := tcp.Dial("49.234.91.99:2333")
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	var forms = make([]dto.OperationForm, 0, 1)
	form1 := dto.OperationForm{
		Language: "Python3",
		FilePath: "",
		CmdLine:  "",
		Code: `print("")print("")

		print("")
		print("")print("")print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")
		print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")
		print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")
		print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")
		print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")
		print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")
		print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")
		print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")
		print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")
		print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")
		print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("")print("")
		print("")print("")
		print("")

		print("")print("")

		print("")
		print("")print("") `,
		Input:          "321",
		ExpectOutput:   "32\n",
		RealOutput:     "",
		ErrorOutput:    "",
		Flag:           "",
		MaxCpuTime:     2,
		ActualCpuTime:  0,
		MaxRealTime:    2,
		ActualRealTime: 0,
		MaxMemory:      20000000,
		RealMemory:     0,
		Score:          10,
		PcId:           0,
	}
	forms = append(forms, form1)
	bytes, err := json.Marshal(&forms)
	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &forms)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	lens, err := conn.Send(bytes)
	fmt.Println("send", lens)
	fmt.Println(len(bytes))
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	_, recv, err := conn.Recv()
	if err != nil {
		fmt.Printf("eof:error:%v", err)
		return
	}
	err = json.Unmarshal(recv, &forms)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	fmt.Printf("%#v", forms)
	fmt.Println(forms)
}
