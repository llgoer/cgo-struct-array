// run GODEBUG=cgocheck=0 go run main.go

package main

/*
#include <stdio.h>
#include <stdlib.h>

typedef struct {
    int a;
    int b;
} Foo;

void pass_struct(Foo *in) { printf("A: %d\tB: %d\n", in->a, in->b); }


void pass_array(Foo **in, int len) {
    for(int i = 0; i < len; i++) {
        printf("A: %d\tB: %d\n", (*in+i)->a, (*in+i)->b);
    }
}

void test_pass_array() {
	Foo in;
	in.a = 1;
	in.b = 2;
	pass_struct(&in);

	Foo *ins;
	ins = (Foo*) malloc(2 * sizeof(Foo));
	for (int i = 0; i < 2; ++i)
	{
		scanf("%d%d", &(ins+i)->a, &(ins+i)->b);
	}
	pass_array(&ins, 2);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type Foo struct{ a, b int32 }

func main() {
	foo := Foo{10, 20}
	foos := []*Foo{&Foo{1, 2}, &Foo{3, 4}}

	fmt.Println("from C land")
	C.pass_struct((*C.Foo)(unsafe.Pointer(&foo)))
	C.pass_struct((*C.Foo)(unsafe.Pointer(foos[0])))
	fmt.Println("from C array land")
	C.pass_array((**C.Foo)(unsafe.Pointer(&foos[0])), C.int(len(foos)))
	fmt.Println("a & b should have incremented with 1")

	fmt.Println("from Go land")
	for _, foo := range foos {
		fmt.Printf("%d : %d\n", foo.a, foo.b)
		C.pass_struct((*C.Foo)(unsafe.Pointer(foo)))
	}

	// C.test_pass_array()
}
