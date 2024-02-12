#include <stdio.h>

int main()
{
    FILE *file;
    char filename[] ="b.txt";
    file = fopen(filename, "rb");
    if(file!=NULL)
    {
        char magicbytes[8];
        size_t bytes = fread(magicbytes,1,sizeof(magicbytes), file);
        
        for (int i = 0; i < sizeof(magicbytes); ++i) {
                printf("%02X ", (unsigned char)magicbytes[i]);
            }
    }
    else{
        printf("Unable to open the file\n");
    }
    
}