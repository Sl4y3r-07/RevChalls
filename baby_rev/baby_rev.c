
#include <stdio.h>
#include <string.h>

int main() {
    printf("Hey, just started with CTFs ??");
   printf("Cool, this is also just a beginning !!! Good Luck.");
    char v6[] = "akf`|Ubq4u26i`X6tXarii&&z";
    int v3 = 7;
    char v[100] = "";

    for (int i = 0; i < strlen(v6); i++) {
        v[i] = v6[i] ^ v3;
    }

   

    return 0;
}


