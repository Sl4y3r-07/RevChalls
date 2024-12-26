package main

import (
	"strings"
	"unicode/utf16"
	"unsafe"

	"golang.org/x/sys/windows"
)


func BTb00Dqff2hrcUJ4e70(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}
	utf16ptr := (*[1 << 20]uint16)(unsafe.Pointer(ptr))[:]
	for i, char := range utf16ptr {
		if char == 0 {
			return string(utf16.Decode(utf16ptr[:i]))
		}
	}
	return ""
}

func t5gftiAC(input string) uint64 {
	var hash uint64
	for _, c := range input {
		hash = (hash*0x1003F + uint64(c)) & 0xffffffff
	}
	return hash
}


func ktUHtuzBYPCRqkfdQFG(dllHash uint64) string {
    peb := windows.RtlGetCurrentPeb()
    ldr := (*windows.PEB_LDR_DATA)(unsafe.Pointer(peb.Ldr))
    head := &ldr.InMemoryOrderModuleList
    current := ldr.InMemoryOrderModuleList.Flink

    for current != head && current != nil {
        entry := (*windows.LDR_DATA_TABLE_ENTRY)(unsafe.Pointer(current))
        dll_name := BTb00Dqff2hrcUJ4e70(uintptr(unsafe.Pointer(entry.FullDllName.Buffer)))
        dll_name = strings.ToLower(dll_name)
        if dllHash == t5gftiAC(dll_name) {
            return dll_name
        }
        current = current.Flink
    }
    return ""
}

