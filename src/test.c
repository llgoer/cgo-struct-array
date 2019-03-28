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

int main() {
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