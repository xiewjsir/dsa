#include "textflag.h"

GLOBL ·helloWorld(SB),NOPTR,$24            // var helloworld []byte("Hello World!")
DATA ·helloWorld+0(SB)/8,$text<>(SB) // StringHeader.Data
DATA ·helloWorld+8(SB)/8,$12         // StringHeader.Len
DATA ·helloWorld+16(SB)/8,$16        // StringHeader.Cap

GLOBL text<>(SB),NOPTR,$16 //text<> ;text只是一个常量名,叫什么都行,<>表示私有变量
DATA text<>+0(SB)/8,$"Hello Wo"      // ...string data...
DATA text<>+8(SB)/8,$"rld!"          // ...string data...
