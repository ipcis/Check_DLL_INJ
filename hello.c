#include <windows.h>

__declspec(dllexport) void ShowMessageBox() {
    MessageBox(NULL, "Hello, World!", "DLL Message", MB_OK);
}
