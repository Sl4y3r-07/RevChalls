package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)
var xLPIULBatJUMIZAFJ string= "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
func T3wDBrdtfkL4OCCEr(chunk string, bit_Length int) string {
	var reversed string
	for i := len(chunk)- 1; i >= 0; i-- {
		reversed += string(chunk[i])
	}
	return fmt.Sprintf("%0*s", bit_Length, reversed) // Important: Ensure it matches the bit length
}

func LOdaKcTJeJJulnUHC(data string, mask int) string {
	var masked_data string
    for i := 0; i < len(data); i += 8 {
	    chunk := data[i : i+8]
        num, err := strconv.ParseInt(chunk, 2, 64)
		if err != nil {
			fmt.Println("Error parsing binary:", err)
			return ""
		}
	
		result := num ^ int64(mask)
		binresult := fmt.Sprintf("%08b", result)
	    masked_data += binresult
	}

	return masked_data
}

func HrjZ1eDoZCypd2kqH(data string) string{

	var bin_data string
	for _, char := range data {
		bin_data += fmt.Sprintf("%08b", char)
	}
	bin_data= LOdaKcTJeJJulnUHC(bin_data, 0xaa)
	padding_length := (6 - len(bin_data) % 6) % 6
    bin_data += strings.Repeat("0", padding_length)
    var chunks []string
	for i := 0; i < len(bin_data); i += 6 {
		chunks = append(chunks, bin_data[i:i+6])
	}
	for i, chunk := range chunks {
		chunks[i] = T3wDBrdtfkL4OCCEr(chunk, 6)
	}
	for i, j := 0, len(chunks)-1; i < j; i, j = i+1, j-1 {
		chunks[i], chunks[j] = chunks[j], chunks[i]
	}

	var enc_string string
	for _, chunk := range chunks {
		index := 0
		for i := 0; i < len(xLPIULBatJUMIZAFJ); i++ {
			if fmt.Sprintf("%06b", i) == chunk {
				index = i
				break
			}
		}
		enc_string += string(xLPIULBatJUMIZAFJ[index])
	}


	return enc_string
}


func cEriifsb8nZJS7BbQ(input string, wg *sync.WaitGroup, result chan bool){
	defer wg.Done()
	var str string = input[0:5]
	if(HrjZ1eDoZCypd2kqH(str)=="NGZI1mz"){   
		result <- true
		return 
	}
	result<-false
}