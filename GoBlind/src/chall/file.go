package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

var (
	kernel32        = syscall.NewLazyDLL("kernel32.dll")
	procWriteFile   = kernel32.NewProc("WriteFile")
	procCreateFileW = kernel32.NewProc("CreateFileW")
	procCloseHandle = kernel32.NewProc("CloseHandle")
)

const (
	GENERIC_WRITE      = 0x40000000
	OPEN_ALWAYS        = 4
	FILE_ATTRIBUTE_NORMAL = 0x80
)

func q0aQYGInCONP5tnB6Ry7NW6Y3K(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func uJ16vre3TdtG4rkHjwO(filename string) (uintptr, error) {
	filenameptr, err := syscall.UTF16PtrFromString(filename)
	if err != nil {
		return 0, err
	}
	r1, _, err := procCreateFileW.Call(
		uintptr(unsafe.Pointer(filenameptr)),
		GENERIC_WRITE,
		0,
		0,
		OPEN_ALWAYS,
		FILE_ATTRIBUTE_NORMAL,
		0,
	)
	if r1 == 0 {
		return 0, err
	}
	return r1, nil
}

func b1hNZ7tNUkwvfteBep5(hFile uintptr, data []byte) error {
	var bytes_Written uint32
	r1, _, err := procWriteFile.Call(
		hFile,
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(&bytes_Written)),
		0,
	)
	if r1 == 0 {
		return err
	}
	return nil
}

func xVX6JBcuXtSC1epR9Ra(hObject uintptr) {
	procCloseHandle.Call(hObject)
}

func R7iQNypRBtbrpyaNq1z(input string) {
	enc_data := []byte{
		90, 60, 81, 81, 163, 74, 21, 106, 136, 66, 204, 83, 241, 42, 166, 94,
		77, 90, 191, 251, 10, 144, 29, 92, 111, 35, 52, 105, 38, 119, 200, 8,
		180, 228, 167, 40, 84, 52, 226, 37, 158, 236, 127, 12, 163, 135, 213, 45,
		219, 255, 241, 64, 114, 83, 165, 190, 235, 3, 31, 243, 157, 162, 230, 10,
		180, 242, 102, 211, 163, 35, 248, 246, 95, 222, 174, 237, 75, 63, 216, 26,
		231, 136, 7, 190, 211, 209, 186, 131, 21, 202, 1, 227, 176, 231, 22, 93,
		34, 205, 250, 192, 197, 178, 26, 12, 210, 165, 112, 130, 181, 152, 164, 9,
		133, 183, 154, 85, 9, 9, 98, 246, 97, 55, 88, 88, 244, 80, 112, 112, 204,
		166, 140, 152, 138, 23, 59, 236, 74, 226, 4, 104, 137, 101, 116, 145, 135,
		229, 72, 246, 101, 37, 215, 166, 119, 166, 123, 19, 213, 198, 200, 14,
	}
	
	key := []byte(input) 

	// decrypting the data in chunks of 16 bytes
	decrypted_data := make([]byte, 0)
	for i := 0; i < len(enc_data); i += 16 {
		block := enc_data[i:min(i+16, len(enc_data))]
		decData := OnkQsAFZuhclKr2sBDT(block, key) 
		decrypted_data = append(decrypted_data, decData...)
	}

	// writing the decrypted data 
	output_File, err := os.Create("./secret.txt") 
	if err != nil {
		q0aQYGInCONP5tnB6Ry7NW6Y3K(err)
		return
	}
	defer output_File.Close()

	
	err = b1hNZ7tNUkwvfteBep5(output_File.Fd(), decrypted_data)
	if err != nil {
		q0aQYGInCONP5tnB6Ry7NW6Y3K(err)
		return
	}

	fmt.Print(":)")
}


