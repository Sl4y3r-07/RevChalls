// #include <stdio.h>
// #include <windows.h>

// int main()
// {
//     const char* dirName ="./n00b";
//     const  char* imagefileName ="./n00b/damn.encrypted.png";
//     const char* textFileName= "./n00b/cool.encrypted.txt";
//     const char* fileNames[]={imagefileName,textFileName};
//     char key[40];

//     const  char* dec_imagefileName ="./n00b/damn.png";
//     const char* dec_textFileName= "./n00b/cool.txt";
//     const char* dec_fileNames[]={dec_imagefileName,dec_textFileName};

//     printf("\n");
//     printf("************************************************************\n");
//     printf("*                                                          *\n");
//     printf("*                       Need Help !!                       *\n");
//     printf("*                                                          *\n");
//     printf("* I encrypted my files a long time ago, which contain a    *\n");
//     printf("* secret of mine. But I lost the key which decrypts them.  *\n");
//     printf("* Can you please help me find out the key?                 *\n");
//     printf("*                                                          *\n");
//     printf("************************************************************\n\n");
//     printf("Gimme Key: ");

//     scanf("%s",key);

//     for(int j=0;j<2;j++){
//      FILE *file;
//      printf("File is: %s",fileNames[j]);
//      printf("\n");
//      file = fopen(fileNames[j], "rb");
//      char magicbytes[8]; int dec_magicbytes; int add;
//      if(file!=NULL)
//      {
//         size_t bytes = fread(magicbytes,1,sizeof(magicbytes),file);
//          printf("Bytes read: %zu\n", bytes);
//         for(int i=0; i<sizeof(magicbytes);i++)
//         {
//             printf("%02x",(unsigned char)magicbytes[i]);
//             printf(" ");
//         }
//         printf("\n");
//         fclose(file);

//         file = fopen(dec_fileNames[j], "wb");
//         fseek(file,0,SEEK_SET);
//         for (int i =0;i<bytes;i++)
//         {   
//             add = (unsigned char)magicbytes[i] -i;
//             dec_magicbytes = (add^key[i]);
//             printf("%02x",(unsigned char)dec_magicbytes); printf(" ");
          
//             fwrite(&dec_magicbytes, 1, 1, file);
//         }
           
        
//         printf("\n");
//         fclose(file);
//     }
//     }  
// }
#include <stdio.h>
#include <windows.h>

int main()
{
    const char *dirName = "./n00b";
    const char *imageFileName = "./n00b/damn.encrypted.png";
    const char* textFileName= "./n00b/cool.encrypted.txt";
    const char* fileNames[]={imageFileName,textFileName};
    const char *dec_imageFileName = "./n00b/damn.png"; 
    const char* dec_textFileName= "./n00b/cool.txt";
    const char* dec_fileNames[]={dec_imageFileName,dec_textFileName};
    char key[40];

    printf("\n");
    printf("************************************************************\n");
    printf("*                                                          *\n");
    printf("*                       Need Help !!                       *\n");
    printf("*                                                          *\n");
    printf("* I encrypted my files a long time ago, which contain a    *\n");
    printf("* secret of mine. But I lost the key which decrypts them.  *\n");
    printf("* Can you please help me find out the key?                 *\n");
    printf("*                                                          *\n");
    printf("************************************************************\n\n");
    printf("Gimme Key: ");

    scanf("%s", key);
    for(int j=0;j<2;j++)
    {
    FILE *imageFile = fopen(fileNames[j], "rb");  
    FILE *decImageFile = fopen(dec_fileNames[j], "wb");  

    if (imageFile != NULL && decImageFile != NULL)
    {
       
        fseek(imageFile, 0, SEEK_END);
        long fileSize = ftell(imageFile);
        fseek(imageFile, 0, SEEK_SET);
        char *fileBuffer = (char*)malloc(fileSize);

        if (fileBuffer != NULL)
        {
            fread(fileBuffer, 1, fileSize, imageFile);
            if (j==0){         
            for (int i = 0; i < 8; i++)
            {
                int add = (unsigned char)fileBuffer[i] - i;
                int dec_magicbyte = (add ^ key[i]);
                fileBuffer[i] = (char)dec_magicbyte;
            }
            }
            else if(j==1)
            {
                for (int i=0;i<37;i++)
                {
                  fileBuffer[i]=fileBuffer[i]^key[i%strlen(key)];
                }
            }
            fwrite(fileBuffer, 1, fileSize, decImageFile);
            free(fileBuffer);
        }
        else
        {
            printf("Error, memory for buffer can't be allocated.\n");
        }

        fclose(imageFile);
        fclose(decImageFile);

        printf("Decryption and copying completed. Check the file: %s\n", dec_fileNames[j]);
    }
    else
    {
        printf("Error opening files.\n");
    }
    }
    return 0;
}
