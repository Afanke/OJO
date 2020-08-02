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
	CCode = `
#include <stdio.h>
int main(){
int a,b;
scanf("%d %d",&a,&b);
printf("%d",a+b);
}
`
	CSPJCode = `
#include <stdio.h>
#include <string.h>
char s1[100]={0};
char s2[100]={0};
char s3[100]={0};
int main(int argc, char *argv[]){
FILE* f1=fopen(argv[1],"r");
FILE* f2=fopen(argv[2],"r");
FILE* f3=fopen(argv[3],"r");
fgets(s1,100,f1); // input
fgets(s2,100,f2); // expected output
fgets(s3,100,f3); // user output
if (strcmp(s2,s3)==0){
	printf("AC");
}else{
	printf("WA");
}
fclose(f1);
fclose(f2);
fclose(f3);
}
`
	CPPCode = `
#include <iostream>
using namespace std;
int main(){
int a,b;
cin>>a;
cin>>b;
cout<<a+b;
}
`
	CPPSPJCode = `
#include <string>
#include <fstream>
#include <sstream>
#include <iostream>
#include <stdlib.h>
using namespace std;

string readFileString(char * filename)
{
ifstream ifile(filename);
ostringstream buf;
char ch;
while(buf&&ifile.get(ch))
buf.put(ch);
return buf.str();
}
  
int main(int argc, char *argv[]){
string s1,s2,s3;
s1=readFileString(argv[1]);
s2=readFileString(argv[2]);
s3=readFileString(argv[3]);
if (s2==s3){
	cout<<"AC";
}else{
	cout<<"WA";
}
}
`
	JavaCode = `
			import java.util.Scanner;
			class Test{
			public static void main(String args[]){
			Scanner sc = new Scanner(System.in);
				int a = sc.nextInt();     
				int b = sc.nextInt();  
				System.out.printf("%d",a+b);
			}
			}
`
	JavaSPJCode = `
	import java.io.*;
	class SPJTest{

		public static void main(String args[]){
			var s1 = readToString(args[0]); // input
			var s2 = readToString(args[1]); // expected output
			var s3 = readToString(args[2]); // user output
			if (s2.equals(s3)) {
			System.out.printf("AC");
			return;
			}
			System.out.printf("WA");
		}

		public static String readToString(String fileName) {  
        String encoding = "UTF-8";  
        File file = new File(fileName);  
        Long filelength = file.length();  
        byte[] filecontent = new byte[filelength.intValue()];  
        try {  
            FileInputStream in = new FileInputStream(file);  
            in.read(filecontent);  
            in.close();  
        } catch (FileNotFoundException e) {  
            e.printStackTrace();  
        } catch (IOException e) {  
            e.printStackTrace();  
        }  
        try {  
            return new String(filecontent, encoding);  
        } catch (UnsupportedEncodingException e) {  
            System.err.println("The OS does not support " + encoding);  
            e.printStackTrace();  
            return null;  
        }  
    }
	
	}
	
	`
	Python3Code = `
i=input().split()
a=int(i[0])
b=int(i[1])
print(a+b,end="")
`
	Python3SPJCode = `
import sys
f1 = open(sys.argv[1], 'r') 
f2 = open(sys.argv[2], 'r') 
f3 = open(sys.argv[3], 'r') 
try: 
    s1 = f1.read() # input
    s2 = f2.read() # expected output
    s3 = f3.read() # user output
    if s2==s3:
        print("AC",end='')
    else:
        print("WA",end='')
finally: 
    f1.close()
    f2.close()
    f3.close()
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
		MaxCpuTime:  3000,
		MaxRealTime: 3000,
		MaxMemory:   1000000,
		TotalScore:  0,
		Id:          0,
		Lid:         3,
		SPJLid:      3,
		SPJMp:       2,
		CompMp:      5,
		Code:        JavaCode,
		SPJCode:     JavaSPJCode,
		Flag:        "",
		TestCase: []dto.TestCase{
			{
				Input:          "1 2",
				ExpectedOutput: "3",
				Score:          10,
			},
			{
				Input:          "4 5",
				ExpectedOutput: "9",
				Score:          10,
			},
		},
	}
	//send("http://192.168.111.139:2333/judge", &form)
	send("http://49.234.91.99:2333/judge", &form)
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
