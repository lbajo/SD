// LORENA BAJO REBOLLO	TELEMÁTICA

#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include "threads.h"

void 
e1(){
	printf("¡HOLA! Soy Thread %d\n     Estoy en e1\n",curidthread());
	exitsthread();
}

void 
e2(){
	printf("¡HOLA! Soy Thread %d\n     Estoy en e2\n",curidthread());
	exitsthread();
}

void
e3(){
	printf("¡HOLA! Soy Thread %d\n     Estoy en e3\n",curidthread());
	exitsthread();
}

void 
bucle(){
	for(int i=0;i<5;i++){
		printf("¡HOLA! Soy Thread %d\n     Estoy en bucle, vuelta %d\n",curidthread(),i);
	}
	exitsthread();

}

void 
suma(){
	int a,b,c;
	a=5;
	b=4;
	c=a+b;
	printf("¡HOLA! Soy Thread %d\n     Estoy en suma, resultado-> %d\n",curidthread(),c);
	exitsthread();
}

int 
main(int argc, char *argv[])
{
	initthreads();
	createthread(e1,NULL,2*1024);
	printf("¡HOLA! Soy Thread %d\n     Soy main \n",curidthread());
	sleep(1);
	yieldthread();
	createthread(e2,NULL,2*1024);
	printf("¡HOLA! Soy Thread %d\n     Soy main \n",curidthread());
	sleep(1);
	yieldthread();
	createthread(e3,NULL,2*1024);
	printf("¡HOLA! Soy Thread %d\n     Soy main \n",curidthread());
	sleep(1);
	createthread(bucle,NULL,2*1024);
	printf("¡HOLA! Soy Thread %d\n     Soy main \n",curidthread());
	yieldthread();
	sleep(1);
	createthread(suma,NULL,2*1024);
	printf("¡HOLA! Soy Thread %d\n     Soy main \n",curidthread());
	createthread(suma,NULL,2*1024);
	yieldthread();
	sleep(1);
	printf("¡HOLA! Soy Thread %d\n     Soy main \n",curidthread());
	exitsthread();
}