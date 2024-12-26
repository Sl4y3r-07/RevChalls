package main

import (
	"time"
	"sync"
)

func PxIuE72uAaCXX3HmXchd(input string,wg *sync.WaitGroup, result chan bool) {
	defer wg.Done()
	currDate := time.Now()
	year := currDate.Year()
	month := currDate.Month()
	var reference_year = 2004
	if year-reference_year != 0 || t5gftiAC(month.String()) != 0xad018838 { // for July
		result <- false
		return
	}
	if input[int(month)]!= 'N'  {           // 7th index
        result<- false
		return
	}
	if input[int(month)+2] != '$' || input[int(month)+3] != 'G' {        // 9th index and 10th
		result<- false
		return
	}
	val:=CWlj47zuCalSy1YWM(input)  // check_5
	if(val!=true){
		result<- false
		return
	}
	result <- true
}