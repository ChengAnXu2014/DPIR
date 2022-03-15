package main

import(
	"os"
	"io"
	"fmt"
	"log"
	"strings"
)


func main(){
	fStr, err:=os.Open("test.dpir-str-utf8")
	if err!=nil{log.Panicln(err)}

	var bld strings.Builder
	var mpStrs=make(map[string]int)
	var count =1

	for{
		var data=make([]byte, 1024)
		n, err:=fStr.Read(data)

		for i:=0;i<n;i++{
			if data[i]=='\r'{continue}
			if data[i]=='\n'{
				mpStrs[bld.String()]=count
				count++
				bld.Reset()
			}

			if data[i]!='\r'&&data[i]!='\n'{bld.WriteByte(data[i])}
			

		}//for bld.Write
		if err==io.EOF{
			if bld.Len()!=0{mpStrs[bld.String()]=count}
			break
		}// if io.EOF
		if err!=nil{log.Panicln(err)}


	}//for fStr.Read

	fmt.Printf("%#v", mpStrs)
}// func main




