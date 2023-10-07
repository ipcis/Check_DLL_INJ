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

type MEMORY_BASIC_INFORMATION struct {
    BaseAddress       uintptr
    AllocationBase    uintptr
    AllocationProtect uint32
    RegionSize        uintptr
    State             uint32
    Protect           uint32
    Type              uint32
}

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

    var mbi MEMORY_BASIC_INFORMATION
    addr := uintptr(0)

    for {
        ret, _, _ := syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualQueryEx").Call(
            uintptr(hProcess),
            addr,
            uintptr(unsafe.Pointer(&mbi)),
            unsafe.Sizeof(mbi),
        )
        if ret == 0 {
            break
        }

        if mbi.Type == 0x10000 { // MEM_IMAGE
            fmt.Printf("Region Start: %#x, Size: %#x\n", mbi.BaseAddress, mbi.RegionSize)
        }

        addr += mbi.RegionSize
    }
}
