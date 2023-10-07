package main

import "C"
import (
    "fmt"
)

//export ShowMessageBox
func ShowMessageBox() {
    caption := "Hello from DLL"
    text := "This is a message from the DLL!"
    C.MessageBox(nil, C.CString(caption), C.CString(text), C.uint(0))
}

func main() {}
