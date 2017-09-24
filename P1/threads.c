// LORENA BAJO REBOLLO	TELEMÁTICA

#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <ucontext.h>
#include <sys/time.h>
#include <time.h>
#include <unistd.h>

#define NUM_THREADS 32

struct Thread{
	ucontext_t ucp;
	int id;
	long long hora;
	void * pila_t;
	int estado;
};

int contador=0;
int hilos_activos=0;
int pos_t_actual=0;
int trazas=0;

typedef struct Thread Thread;

Thread array_threads[NUM_THREADS];

long long 
horams(void){
	struct timeval t1;
	gettimeofday(&t1, NULL);
	return (t1.tv_sec*1000) + (t1.tv_usec /1000);
}

int 
buscarposicion(void){
	int pos_sig;
	
	for (int p=pos_t_actual+1;p<NUM_THREADS;p++){
		if (array_threads[p].estado==0){
			free(array_threads[p].pila_t);
			array_threads[p].pila_t=NULL;
		}
		if (array_threads[p].estado==1){
			pos_sig=p;
			return pos_sig;
		}
	}
	for(int p=0;p<=pos_t_actual;p++){
		if (array_threads[p].estado==0){
			free(array_threads[p].pila_t);
			array_threads[p].pila_t=NULL;
		}
		if (array_threads[p].estado==1){
			pos_sig=p;	
			return pos_sig;
		}
	}
	return -9999;
}

int 
temporizador(void){
	int cuanto=200;
	int tiempo_ejecucion;
	long long hora_actual=horams();
	tiempo_ejecucion=hora_actual- array_threads[pos_t_actual].hora;
	if(tiempo_ejecucion>=cuanto){
		return 1;
	}
	return 0;
}

void 
initthreads(void){

	Thread nuevoThread;

	nuevoThread.id = contador;
	nuevoThread.hora=horams();
	nuevoThread.pila_t=NULL;
	nuevoThread.estado=1;
	array_threads[0]=nuevoThread;
	contador++;
	hilos_activos++;
}

int
createthread(void(*mainf)(void*),void*arg,int stacksize){

	if(hilos_activos <= NUM_THREADS){
		Thread nuevoThread;

		nuevoThread.pila_t=malloc(stacksize);
		getcontext(&nuevoThread.ucp);
		nuevoThread.ucp.uc_stack.ss_sp=nuevoThread.pila_t;
		nuevoThread.ucp.uc_stack.ss_size=stacksize;
		makecontext(&nuevoThread.ucp,(void (*)(void ))mainf,1,arg);

		nuevoThread.id = contador;
		nuevoThread.hora = horams();
		nuevoThread.estado=1;

		hilos_activos++;
		contador++;

		for(int i=0;i<NUM_THREADS;i++){
			if (array_threads[i].estado==0){
				array_threads[i] = nuevoThread;
				return nuevoThread.id;
			}
		}
	}
	return -9999;
}

int 
curidthread(void){
	return array_threads[pos_t_actual].id;
}

void 
exitsthread(void){
	ucontext_t *ctx_actual;
	int pos;
	if (trazas==1){
		printf("EXITSTHREAD --> Lista actual\n");
		printf("|| Thread(id) || posición ||  estado  ||    hora       || \n");
		for(int i=0;i<NUM_THREADS;i++){
			if(array_threads[i].estado == 1){
				printf("       %d            %d            %d       %lld\n",array_threads[i].id,i, array_threads[i].estado,array_threads[i].hora);
			}
		}
		printf("\n");
	}

	pos=buscarposicion();
	ctx_actual=&array_threads[pos_t_actual].ucp;
	
	if(pos_t_actual==pos){
		if (trazas==1){
			printf("No hay más threads. Fin de programa \n");
		}
		free(array_threads[pos_t_actual].pila_t);
		array_threads[pos_t_actual].pila_t=NULL;
		hilos_activos--;
		exit(0);

	}else{
		if (trazas==1){
			printf("Quedan %d threads\n",hilos_activos);
			printf("Soy %d en la posición %d --> ", curidthread(),pos_t_actual);
		}

		array_threads[pos_t_actual].estado=0;
		array_threads[pos].hora=horams();
		pos_t_actual=pos;
		hilos_activos--;

		if (trazas==1){
			printf("Salto a %d en la posición %d\n", curidthread(),pos);
		}
		swapcontext(ctx_actual,&array_threads[pos_t_actual].ucp);
	}
}

void 
yieldthread(void){
	ucontext_t *ctx_actual;
	int pos;

	if (trazas==1){
		printf("YIELDTHREAD --> Lista actual \n");
		printf("|| Thread(id) || posición ||  estado  ||    hora       || \n");
		for(int i=0;i<NUM_THREADS;i++){
			if(array_threads[i].estado == 1){
				printf("       %d            %d            %d       %lld\n",array_threads[i].id,i, array_threads[i].estado,array_threads[i].hora);
			}
		}
		printf("\n");
	}
	if (temporizador()==1){

		pos=buscarposicion();
		if(pos==pos_t_actual){
			if (trazas==1){
				printf("Continúo, soy: %d, en la posición %d\n",curidthread(),pos_t_actual);
			}

			array_threads[pos_t_actual].hora=horams();
			return;
		}

		ctx_actual=&array_threads[pos_t_actual].ucp;

		if (trazas==1){
			printf("Soy %d en la posición %d --> ", curidthread(),pos_t_actual);
		}
		
		pos_t_actual=pos;
		array_threads[pos_t_actual].hora=horams();	

		if (trazas==1){			
			printf("Salto a %d en la posición %d\n", curidthread(),pos);
		}
		
		swapcontext(ctx_actual,&array_threads[pos_t_actual].ucp);
		return;	
	}
}
