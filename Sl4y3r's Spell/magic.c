#include <stdio.h>

int main() {
    FILE *file;
    char filename[] = "a.png";

    file = fopen(filename, "rb");
    if (file != NULL) {
        char magicBytes[8]; 
        size_t bytesRead = fread(magicBytes, 1, sizeof(magicBytes), file);

        if (bytesRead == sizeof(magicBytes)) {
            printf("Magic bytes: ");
            for (int i = 0; i < sizeof(magicBytes); ++i) {
                printf("%02X ", (unsigned char)magicBytes[i]);
            }
            printf("\n");
        } else {
            printf("Unable to read magic bytes\n");
        }

        fclose(file);
    } else {
        printf("Unable to open file\n");
    }

    return 0;
}
