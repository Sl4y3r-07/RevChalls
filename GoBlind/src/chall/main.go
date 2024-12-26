package main

import (
	"fmt"
	"sync"
	"golang.org/x/sys/windows"
)

func NcfAwgRxhT4ab() bool{
	dll_name := ktUHtuzBYPCRqkfdQFG(0x8f7ee672)
	kernel32:= windows.NewLazySystemDLL(dll_name)
    isDebuggerPresent:= kernel32.NewProc("IsDebuggerPresent")
	ret, _, _ := isDebuggerPresent.Call()
	return ret != 0
}

func ismHQJZnri0vV(flags uint32) error {
	user32 := windows.NewLazySystemDLL("user32.dll")
	exitWindowsEx_ := user32.NewProc("ExitWindowsEx")
	ret, _, err := exitWindowsEx_.Call(uintptr(flags), 0)
	if ret == 0 { 
		return fmt.Errorf("ExitWindowsEx failed: %v", err)
	}
	return nil
}

func pQbwBvcFxubQ9() {
	if NcfAwgRxhT4ab() {
	   ismHQJZnri0vV(windows.EWX_LOGOFF)
	} 
}

func O6t7LeVzef(input string) bool {
	if len(input)^0x10 != 0 {
		return false
	}
	return true
}


func tUuCkSywvuVtXI() {
	fmt.Println(" _____                                   _     __  __        ")
	fmt.Println("|  __ \\                                 | |   |  \\/  |       ")
	fmt.Println("| |  | |  ___   ___  _ __  _   _  _ __  | |_  | \\  / |  ___  ")
	fmt.Println("| |  | | / _ \\ / __|| '__|| | | || '_ \\ | __| | |\\/| | / _ \\ ")
	fmt.Println("| |__| ||  __/| (__ | |   | |_| || |_) || |_  | |  | ||  __/ ")
	fmt.Println("|_____/  \\___| \\___||_|    \\__, || .__/  \\__| |_|  |_| \\___| ")
	fmt.Println("                             __/ || |                        ")
	fmt.Println("                            |___/ |_|                        ")
}


func KHjPRKrWguCBOPh3uiv(){
		fmt.Println("\n\nI have somehow recovered this file from a compromised system but...")
		fmt.Println("It asks for a password to reveal some super secret information..")
		fmt.Println("Can you please help me recovering the secret ? Good Luck!!")
}
func main(){
	var input string
	var wg sync.WaitGroup 
	pQbwBvcFxubQ9()  
    tUuCkSywvuVtXI()
	KHjPRKrWguCBOPh3uiv()
	fmt.Print("Enter the password: ")
    fmt.Scanf("%s", &input)
	res:=O6t7LeVzef(input)        //input check
    if !res{
		fmt.Print("Invalid Length") // wallpaper chanege krdunga idhar
        return
	}
	result := WO7XAiKm70O1Llq7foIP(input)            // check_6
	wg.Add(result)
	chan1:= make(chan bool,1)
	chan4:= make(chan bool,1)
	chan5:= make(chan bool,1)
	chan6:= make(chan bool,1)
	go PxIuE72uAaCXX3HmXchd(input, &wg, chan1)   // check_1
	go JcZGhfFWaFsVCIwGz6lVoX(input, &wg, chan4)      // check_2
	go QQRDzgcGfEBGY3XF0eLnHbi(input, &wg,chan5) // check_3
	go cEriifsb8nZJS7BbQ(input, &wg, chan6)     // check_4
	wg.Wait()
	system_Date_OK := <-chan1
	file_Name_Ok := <- chan4
	no_file_ok:= <- chan5
	check_custom_b64:= <- chan6
	if !system_Date_OK|| !file_Name_Ok || !no_file_ok || !check_custom_b64{
		ismHQJZnri0vV(windows.EWX_LOGOFF)    
	    // fmt.Print("wrong")
		return
		}	
	R7iQNypRBtbrpyaNq1z(input)
	Z2JEqRKSdxHoRaSCjUshuqZUVD(input)           // check_7
}




