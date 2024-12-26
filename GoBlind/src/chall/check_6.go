package main 

func WO7XAiKm70O1Llq7foIP(input string) int{
	char := input[11]           //11th index
    result := (int(char) >> 2)   
    result = (result % 10) + 2   // (2 for 4 routines)
	return result
}