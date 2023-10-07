# 

```
C:\Users\MWuser\go_projekt\Check_DLL_INJ>go build -o do_dll_inj.exe do_dll_inj.go
C:\Users\MWuser\go_projekt\Check_DLL_INJ>go mod init dll_injector
C:\Users\MWuser\go_projekt\Check_DLL_INJ>go mod tidy
C:\Users\MWuser\go_projekt\Check_DLL_INJ>go build -o do_dll_inj.exe do_dll_inj.go
```



compile dll
```
set CGO_ENABLED=1
go -o simple_dll.dll -buildmode=c-shared simple_dll.go
go build -o mydll.dll -buildmode=c-shared mydll.go

gcc -shared -o hello.dll hello.c -Wl,--out-implib,libhello.a

```
C:\Program Files\Go\pkg\tool\windows_amd64\link.exe: running gcc failed: exec: "gcc": executable file not found in %PATH%

set PATH=%PATH%;C:\Users\MWuser\Downloads\x86_64-13.2.0-release-posix-seh-msvcrt-rt_v11-rev0\mingw64\bin
