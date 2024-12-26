package main

import (
    "os"
	"path/filepath"
	"sync"
)

func JcZGhfFWaFsVCIwGz6lVoX( input string, wg*sync.WaitGroup, result chan bool){  
	defer wg.Done()
	 exePath,_:= os.Executable()
	 exeName := filepath.Base(exePath)
	 if exeName[8]==input[8] && exeName[5]==input[12] {     // 8th and 12th index
           result <- true
	       return
	 }
	 result <- false
}