#include<stdio.h>

#define NUMBER  5

void main(void){
    FILE *fp;
    int age = 0;
    int height = 0;

    fp = fopen("abc.txt","r+");
    if(fp == NULL){
        printf("error\n");
    }else{
        fprintf(fp,"%d",60);
        while((age = fgetc(fp)) != EOF){
            putchar(age);
        }
        fclose(fp);
    }
}
