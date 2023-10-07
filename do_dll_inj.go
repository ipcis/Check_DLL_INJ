package main

import (
    "fmt"
    "github.com/JamesHovious/w32"
    "syscall"
)

func main(){
    //just a random dll
    dllPath := "C:\\Windows\\SysWOW64\\wevtapi.dll"
    procID := 8124
    hProc, handleErr  := w32.OpenProcess(w32.PROCESS_ALL_ACCESS, false, uint32(procID))
    if handleErr != nil{
        fmt.Println("handleErr:", handleErr)
    }
    kernel32DLL, dllLoadErr := syscall.LoadLibrary("kernel32.dll")
    if dllLoadErr != nil{
        fmt.Println("dllLoadErr:", dllLoadErr)
    }
    addr, addrErr := syscall.GetProcAddress(syscall.Handle(kernel32DLL), "LoadLibraryA")
    if addrErr != nil{
        fmt.Println("addrErr:", addrErr)
    }
    arg, allocErr := w32.VirtualAllocEx(hProc, 0, len(dllPath)*2, w32.MEM_RESERVE | w32.MEM_COMMIT, w32.PAGE_READWRITE)
    if allocErr != nil{
        fmt.Println("allocErr:", allocErr)
    }
    writeErr := w32.WriteProcessMemory(hProc, uint32(arg), []byte(dllPath), 0)
    if writeErr != nil{
        fmt.Println("writeErr:", writeErr)
    }
    _, _, threadErr := w32.CreateRemoteThread(hProc, nil, 0, uint32(addr), arg, 0)
    if threadErr != nil{
        fmt.Println("threadErr:", threadErr)
    }
}
