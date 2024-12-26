package main

import (
	
	"runtime"
)

func nOsylsKWBryE2UkXS() int {
	stack := make([]uintptr, 100)
	depth := runtime.Callers(0, stack)
	stack = stack[:depth]
    return depth
}
func CWlj47zuCalSy1YWM(input string) bool {
    res := nOsylsKWBryE2UkXS()
	char_Xor := input[13] ^ input[14]   // 13th 14th index 
    val := int((char_Xor >> 1) % 14)  // check 
    if (val == res) {         // 5   
        return true
	}
	return false
}