#include <windows.h>
#include <stdio.h>

int main() {
    DWORD processId = 12345; // Ersetzen Sie 12345 durch die tatsächliche Prozess-ID des Zielprozesses

    HANDLE hProcess = OpenProcess(PROCESS_VM_READ, FALSE, processId);
    if (hProcess == NULL) {
        printf("Fehler beim Öffnen des Prozesses. Fehlercode: %d\n", GetLastError());
        return 1;
    }

    // Hier den Speicherbereich definieren, den Sie untersuchen möchten
    LPVOID addressToCheck = (LPVOID)0x00400000; // Ersetzen Sie dies durch die gewünschte Adresse

    MEMORY_BASIC_INFORMATION memInfo;
    SIZE_T bytesRead = VirtualQueryEx(hProcess, addressToCheck, &memInfo, sizeof(memInfo));
    if (bytesRead == 0) {
        printf("Fehler beim Abfragen des Speicherbereichs. Fehlercode: %d\n", GetLastError());
        CloseHandle(hProcess);
        return 1;
    }

    // Überprüfen, ob der Speicherbereich PAGE_EXECUTE_READWRITE ist
    if (memInfo.Protect == PAGE_EXECUTE_READWRITE) {
        printf("Der Speicherbereich ist PAGE_EXECUTE_READWRITE.\n");
    } else {
        printf("Der Speicherbereich ist nicht PAGE_EXECUTE_READWRITE.\n");
    }

    CloseHandle(hProcess);
    return 0;
}
