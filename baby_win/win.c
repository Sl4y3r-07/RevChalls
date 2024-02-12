#include <Windows.h>
#include <string.h>
#include <stdio.h>

const char* sub() {
    char a[] = "pa55wd";
    char b[] = "A/AGG;$QjbF*\"Rc";
    char* c = (char*)malloc(strlen(b) + 1);
    if (c != NULL) {
        for (int i = 0; i < strlen(b); i++) {
            c[i] = b[i] ^ a[i % strlen(a)];
        }
        c[strlen(b)] = '\0';
    }
    return c;
}
LRESULT CALLBACK WindowProc(HWND hwnd, UINT uMsg, WPARAM wParam, LPARAM lParam) {
    switch (uMsg) {
        case WM_CREATE: {
            CreateWindow("BUTTON", "OK",
                WS_VISIBLE | WS_CHILD | BS_PUSHBUTTON,
                150, 100, 80, 25, hwnd, (HMENU)1, NULL, NULL);
            CreateWindowEx(WS_EX_CLIENTEDGE, "EDIT", "Type the password here !",
                WS_CHILD | WS_VISIBLE | ES_AUTOHSCROLL,
                50, 50, 200, 25, hwnd, (HMENU)2, NULL, NULL);
        }
        break;
        case WM_COMMAND: {
            if (LOWORD(wParam) == 1) { 
                char buffer[100];
                GetWindowText(GetDlgItem(hwnd, 2), buffer, 100); 

                const char* desiredString = sub();
                if (strcmp(buffer, desiredString) == 0) {
                    MessageBox(hwnd, "Password Correct", "Input Check", MB_OK | MB_ICONINFORMATION);
                    PostQuitMessage(0);
                    free((void*)desiredString); // Cast to void* before passing to free()

                }
                else {
                    MessageBox(hwnd, "Password Incorrect", "Input Check", MB_ICONERROR);
                    free((void*)desiredString); // Cast to void* before passing to free()

                }
            }
        }
        break;
        case WM_DESTROY:
            PostQuitMessage(0);
            return 0;
        default:
            return DefWindowProc(hwnd, uMsg, wParam, lParam);
    }
    return 0;
}

int WINAPI WinMain(HINSTANCE hInstance, HINSTANCE hPrevInstance, LPSTR lpCmdLine, int nCmdShow) {
    const char* CLASS_NAME = "Input";
    WNDCLASS wc = {0};

    wc.lpszClassName = CLASS_NAME;
    wc.lpfnWndProc = WindowProc;
    wc.hInstance = hInstance;

    RegisterClass(&wc);
    HWND hwnd = CreateWindowEx(0, CLASS_NAME, "Input program", WS_OVERLAPPEDWINDOW,
        CW_USEDEFAULT, CW_USEDEFAULT, 300, 200, NULL, NULL, hInstance, NULL);
    if (hwnd == NULL)
        return 0;

    ShowWindow(hwnd, nCmdShow);

    MSG msg = { 0 };
    while (GetMessage(&msg, NULL, 0, 0)) {
        TranslateMessage(&msg);
        DispatchMessage(&msg);
    }

    return msg.wParam;
}
