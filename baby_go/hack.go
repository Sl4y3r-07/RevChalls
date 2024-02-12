package main

import "fmt"

func main(){
	fmt.Println("This is a baby challenge !!")
     str := "fnac}A6Wzo|?^}=~uMp vom9c"
	 res :=""
	 st  :=""
	 for i:=0; i<len(str);i++{
		if i%2 == 0 {
			st += string(str[i] + 1)
		} else {
			st += string(str[i])
		}
		
	 }
	 for i:=0; i<len(str);i++{
		res += string(st[i]^byte(i+1))
	 }
	
}
