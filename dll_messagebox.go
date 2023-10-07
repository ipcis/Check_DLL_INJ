package main

import (
    "syscall"
    "unsafe"
)

var (
    user32           = syscall.NewLazyDLL("user32.dll")
    messageBox       = user32.NewProc("MessageBoxW")
    kernel32         = syscall.NewLazyDLL("kernel32.dll")
    getModuleHandleW = kernel32.NewProc("GetModuleHandleW")
)

func MessageBox(caption, text string) int {
    moduleName := syscall.StringToUTF16Ptr("")
    handle, _, _ := getModuleHandleW.Call(uintptr(unsafe.Pointer(moduleName)))
    return int(messageBox.Call(
        0,
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
        uintptr(0),
    ))
}

func main() {}

//export ShowMessageBox
func ShowMessageBox() {
    caption := "Hello from DLL"
    text := "This is a message from the DLL!"
    MessageBox(caption, text)
}
