#include <stdio.h>
#include <windows.h>


unsigned char sub(unsigned char b, int i) {
    return ((b >> i) | (b << (8 - (i % 8)))) & 255;
}

void sub_04f243(const char* encrypted_text, char* decrypted_text) {
    int encrypted_length = strlen(encrypted_text);
    int decrypted_index = 0;
    
    for (int i = 0; i < encrypted_length; i += 8) {
        const char* block_encrypted = &encrypted_text[i];
        char block_decrypted[9];
        block_decrypted[8] = '\0'; 
        
        for (int j = 0; j < 8; j++) {
             block_decrypted[j] = sub(((block_encrypted[j]+j)%256), j);    //(block_encrypted[j]+j)%256) :- add
        }
       
        int trimmed_length = 8;
        while (trimmed_length > 0 && block_decrypted[trimmed_length - 1] == '\x00') {
            trimmed_length--; 
        }
        
    
        strncpy(&decrypted_text[decrypted_index], block_decrypted, trimmed_length);
        decrypted_index += trimmed_length;
    }
    decrypted_text[decrypted_index] = '\0'; 
}

int main() {
    const char* directoryPath = "./n00b/"; 
    const char* fileName = "/flag.txt"; 
    char buffer[1000]; 
    HANDLE hConsole = GetStdHandle(STD_OUTPUT_HANDLE);
    DWORD written;
   
    DWORD dwAttrib = GetFileAttributesA(directoryPath);
    if (dwAttrib != INVALID_FILE_ATTRIBUTES && (dwAttrib & FILE_ATTRIBUTE_DIRECTORY)) {
          
        char filePath[MAX_PATH];
        snprintf(filePath, MAX_PATH, "%s\\%s", directoryPath, fileName);
        
        dwAttrib = GetFileAttributesA(filePath);
        if (dwAttrib != INVALID_FILE_ATTRIBUTES && !(dwAttrib & FILE_ATTRIBUTE_DIRECTORY)) {
                     FILE* file = fopen(filePath, "r");
            if (file != NULL) {
                
                while (fgets(buffer, sizeof(buffer), file) != NULL) {
                    WriteConsole(hConsole, buffer, strlen(buffer), &written, NULL);
                    WriteConsoleA(hConsole, "\n", 1, &written, NULL);
                }
                fclose(file);
            } else {
                printf("Unable to open the file.\n");
            }
        } else {
            printf("File does not exist in the directory.\n");
        }
    } else {
        printf("Directory does not exist.\n");
    }

int a =40; int b= 2000; int c =30; 
if (a<20 ){
const char* encrypted_text = buffer; 
    char decrypted_text[1000]; 
    if(c<10)
    {
        if(b<1000 ){
    sub_04f243(encrypted_text, decrypted_text);
    WriteConsole(hConsole,"Decrypted text is -> ", strlen("Decrypted text is -> "),&written,NULL);
    WriteConsole(hConsole, decrypted_text, strlen(decrypted_text), &written, NULL);
     WriteConsoleA(hConsole, "\n", 1, &written, NULL);
         WriteConsole(hConsole,"Hehe...", strlen("Hehe..."),&written,NULL);
        }else{
        WriteConsole(hConsole,"Decrypted text is -> ", strlen("Decrypted text is -> "),&written,NULL);
        WriteConsole(hConsole, encrypted_text, strlen(encrypted_text), &written, NULL);
        WriteConsoleA(hConsole, "\n", 1, &written, NULL);
         WriteConsole(hConsole,"Hehe...", strlen("Hehe..."),&written,NULL);
        }
}}
    return 0;
}




