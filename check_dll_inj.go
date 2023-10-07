package main

import (
    "fmt"
    "os"
    "strconv"
    "syscall"
    "unsafe"
)

const (
    PROCESS_QUERY_INFORMATION = 0x0400
    PROCESS_VM_READ           = 0x0010
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: dll-inject-checker <process_pid>")
        os.Exit(1)
    }

    pidStr := os.Args[1]
    targetPID, err := strconv.Atoi(pidStr)
    if err != nil {
        fmt.Println("Ungültige PID:", err)
        os.Exit(1)
    }

    hProcess, err := syscall.OpenProcess(PROCESS_QUERY_INFORMATION|PROCESS_VM_READ, false, uint32(targetPID))
    if err != nil {
        fmt.Println("Fehler beim Öffnen des Prozesses:", err)
        os.Exit(1)
    }
    defer syscall.CloseHandle(hProcess)

    var mbi syscall.MemoryBasicInformation
    var bytesRead uintptr
    addr := uintptr(0)

    for {
        ret, _, _ := syscall.VirtualQueryEx(hProcess, addr, &mbi, uint32(unsafe.Sizeof(mbi)))
        if ret == 0 {
            break
        }

        if mbi.Type&syscall.MEM_IMAGE != 0 {
            fmt.Printf("Region Start: %#x, Size: %#x\n", mbi.BaseAddress, mbi.RegionSize)
        }

        addr += uintptr(mbi.RegionSize)
    }
}
