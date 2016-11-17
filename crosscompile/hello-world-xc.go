package main

// #include <stdio.h>
//
// void hello_world() {
//     printf("Hello World\n");
// }
import "C"

func main() {
	C.hello_world()
}
