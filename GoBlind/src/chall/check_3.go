package main

import(
	 "path/filepath"
	 "os"
	 "sync"
	 
	//  "fmt"
)

func QQRDzgcGfEBGY3XF0eLnHbi(input string, wg *sync.WaitGroup, result chan bool) {
	defer wg.Done()
	exe_Path,_:= os.Executable()
	dir_Name := filepath.Dir(exe_Path)
	files, _ := os.ReadDir(dir_Name)
	 counter:=0
	 for _, entry := range files{
		if !entry.IsDir(){
		counter++
		}
	 }
	
	 if counter!=5{
		result<-false
		return
	 }
	 if input[counter+1]!='o'{  // 6th index
        result<- false
		return
	 }
	 result<- true
}