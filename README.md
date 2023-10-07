# 
C:\Users\MWuser\go_projekt\Check_DLL_INJ>go build -o do_dll_inj.exe do_dll_inj.go
do_dll_inj.go:5:5: no required module provides package github.com/JamesHovious/w32: go.mod file not found in current directory or any parent directory; see 'go help modules'

C:\Users\MWuser\go_projekt\Check_DLL_INJ>go mod init dll_injector
go: creating new go.mod: module dll_injector
go: to add module requirements and sums:
        go mod tidy

C:\Users\MWuser\go_projekt\Check_DLL_INJ>go mod tidy
go: finding module for package github.com/JamesHovious/w32
go: downloading github.com/JamesHovious/w32 v1.2.0
go: found github.com/JamesHovious/w32 in github.com/JamesHovious/w32 v1.2.0

C:\Users\MWuser\go_projekt\Check_DLL_INJ>go build -o do_dll_inj.exe do_dll_inj.go
