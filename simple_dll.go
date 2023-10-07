package main

import "C"
import (
 "syscall"

 "golang.org/x/sys/windows"
)

func init() {
 windows.MessageBox(windows.HWND(0), syscall.StringToUTF16Ptr("Injected"), syscall.StringToUTF16Ptr("Injection works"), windows.MB_OK)
}

func main() {}
