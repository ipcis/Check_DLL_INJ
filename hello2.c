#include <windows.h>

DWORD WINAPI ThreadProc(LPVOID lpParameter) {
    MessageBox(NULL, "Hello, World!", "Message Box", MB_OK);
    return 0;
}

BOOL APIENTRY DllMain(HMODULE hModule, DWORD  ul_reason_for_call, LPVOID lpReserved) {
    if (ul_reason_for_call == DLL_PROCESS_ATTACH) {
        HANDLE hThread = CreateRemoteThread(GetCurrentProcess(), NULL, 0, ThreadProc, NULL, 0, NULL);
        if (hThread != NULL) {
            WaitForSingleObject(hThread, INFINITE);
            CloseHandle(hThread);
        }
    }
    return TRUE;
}
