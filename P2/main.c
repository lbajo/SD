// LORENA BAJO REBOLLO	TELEMÁTICA

#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include "threads.h"

void 
e1(){
	printf("¡HOLA! Soy Thread %d\n     Estoy en e1\n",curidthread());
	printf("	Me voy a dormir\n");
	sleepthread(3000);
	printf("¡HOLA! Soy Thread %d\n     Me estoy despertando!\n",curidthread());
	exitsthread();
}

void 
e2(){
	printf("¡HOLA! Soy Thread %d\n     Estoy en e2\n",curidthread());
	suspendthread();
	printf("¡HOLA! Soy Thread %d\n     Vuelvo a estar activo\n",curidthread());
	exitsthread();
}

void
e3(){
	printf("¡HOLA! Soy Thread %d\n     Estoy en e3\n",curidthread());
	exitsthread();
}

void 
bucle(){
	for(int i=0;i<2;i++){
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
	killthread(6);
	exitsthread();
}

int 
main(int argc, char *argv[])
{
	int *lista;

	initthreads();
	sleep(1);
	createthread(e1,NULL,5*1024);
	printf("¡HOLA! Soy Thread %d\n     Soy main \n",curidthread());
	yieldthread();
	createthread(e2,NULL,20*1024);
	printf("¡HOLA! Soy Thread %d\n     Soy main \n",curidthread());	
	sleep(1);
	yieldthread();
	printf("¡HOLA! Soy Thread %d\n     Soy main \n",curidthread());
	createthread(e2,NULL,5*1024);
	yieldthread();
	createthread(e3,NULL,5*1024);
	sleep(1);
	yieldthread();
	printf("¡HOLA! Soy Thread %d\n     Soy main \n",curidthread());
	sleep(1);
	yieldthread();
	createthread(bucle,NULL,5*1024);
	printf("¡HOLA! Soy Thread %d\n     Soy main \n",curidthread());
	resumethread(2);
	yieldthread();
	sleep(1);
	createthread(suma,NULL,5*1024);
	yieldthread();
	printf("¡HOLA! Soy Thread %d\n     Soy main \n",curidthread());
	createthread(e2,NULL,5*1024);
	suspendedthreads(&lista);
	killthread(7);
	printf("¡HOLA! Soy Thread %d\n     Soy main, Bye! \n",curidthread());
	sleep(1);
	free(lista);
	exitsthread();
}