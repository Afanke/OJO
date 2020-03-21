package main

import (
	"encoding/json"
	"fmt"
	"github.com/afanke/OJO/JudgeServer/dto"
	"github.com/afanke/OJO/JudgeServer/operator"
	tcp "github.com/afanke/OJO/utils/tcp"
)

var operationMap = map[string]operator.Operator{
	"Python3": operator.PythonOperator{},
}

func handle(conn tcp.Conn) {
	defer func() {
		conn.Close()
		fmt.Println("close")
	}()
	_, recv, err := conn.Recv()
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	var forms []dto.OperationForm
	err = json.Unmarshal(recv, &forms)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	for i := 0; i < len(forms); i++ {
		o := operationMap[forms[i].Language]
		if o != nil {
			o.Operate(&forms[i])
		}
	}
	res, err := json.Marshal(forms)
	if err != nil {
		fmt.Printf("error:%v", err)
	}
	_, err = conn.Send(res)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
}

func main() {
	listen, err := tcp.Listen(":2333")
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	fmt.Println("start")
	fmt.Println(listen.Addr())
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		fmt.Println("connect")
		if err != nil {
			fmt.Printf("error:%v", err)
			return
		}
		go handle(conn)
	}
}
