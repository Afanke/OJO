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

const (
	JavaCode = `
			import java.util.Scanner;
			class Hello{
			public static void main(String args[]){
			Scanner sc = new Scanner(System.in);
				int a = sc.nextInt();     
				int b = sc.nextInt();  
				System.out.printf("%d",a+b);
			}
			}
`
	GoCode = `
			package main
			import (
			"fmt"
			)
			
			func main() {
			var a int
			var b int
			fmt.Scanf("%d %d",&a,&b)
			fmt.Printf("%d",a+b)
			}
`
	GoSPJCode = `
			package main
			import(
				"io/ioutil"
				"os"
				"fmt"
			)
			
			func main(){
				// b1, _ := ioutil.ReadFile(os.Args[1]) 
				b2, _ := ioutil.ReadFile(os.Args[2]) 
				b3, _ := ioutil.ReadFile(os.Args[3]) 
				// s1:=string(b1)  // input
				s2:=string(b2)  // expected output
				s3:=string(b3)  // user output
				if s2==s3{
					fmt.Printf("AC")
					return
				}
				fmt.Printf("WA")
			}
`
)

func main() {

	form := dto.JudgeForm{
		UseSPJ:      true,
		MaxCpuTime:  1000,
		MaxRealTime: 1000,
		MaxMemory:   1000000,
		TotalScore:  0,
		Id:          0,
		Lid:         5,
		SPJLid:      5,
		Code:        GoCode,
		SPJCode:     GoSPJCode,
		Flag:        "",
		TestCase: []dto.TestCase{
			{
				Input:          "1 2",
				ExpectedOutput: "3 ",
				Score:          10,
			},
			{
				Input:          "4 5",
				ExpectedOutput: "9",
				Score:          10,
			},
		},
	}
	send("http://192.168.111.139:2333/judge", &form)
	//send("http://49.234.91.99:2333/judge",&form)
}

func send(addr string, form *dto.JudgeForm) {
	client := &http.Client{
		Timeout: time.Duration(form.MaxRealTime) * time.Duration(len(form.TestCase)) * time.Second * 2,
	}
	buff, err := json.Marshal(&form)
	if err != nil {
		log.Error("error:%v\n", err)
		return
	}
	res, err := client.Post(addr, "application/json", bytes.NewBuffer(buff))
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
