// LORENA BAJO REBOLLO	TELEMÁTICA

#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <ucontext.h>
#include <sys/time.h>
#include <time.h>
#include <unistd.h>
#include <sys/types.h>
#include <err.h>

#define NUM_THREADS 32

struct Thread{
	ucontext_t ucp;
	int id;
	long long hora;
	long long hora_despertar;
	void * pila_t;
	int estado;
};
/*	ESTADOS

	0->No activado
	1->Activado, listo para ejecutar
	2->Suspendido
	3->Dormido
*/

int contador=0;
int hilos_activos=0;
int hilos_suspendidos=0;
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
		if (array_threads[p].estado==3){
			if(horams()>=array_threads[p].hora_despertar){
				array_threads[p].estado=1;
			}
		}
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
		if (array_threads[p].estado==3){
			if(horams()>=array_threads[p].hora_despertar){
				array_threads[p].estado=1;
			}
		}
		if (array_threads[p].estado==0){
			free(array_threads[p].pila_t);
			array_threads[p].pila_t=NULL;
		}
		if (array_threads[p].estado==1){
			pos_sig=p;	
			return pos_sig;
		}
	}
	return -1;
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
		nuevoThread.id = contador;
		nuevoThread.hora = horams();
		nuevoThread.estado=1;

		getcontext(&nuevoThread.ucp);
		nuevoThread.ucp.uc_stack.ss_sp=nuevoThread.pila_t;
		nuevoThread.ucp.uc_stack.ss_size=stacksize;
		makecontext(&nuevoThread.ucp,(void (*)(void ))mainf,1,arg);

		hilos_activos++;
		contador++;

		for(int i=0;i<NUM_THREADS;i++){
			if (array_threads[i].estado==0){
				array_threads[i] = nuevoThread;
				return nuevoThread.id;
			}
		}
	}
	return -1;
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

void 
suspendthread(void){
	int pos_sig,s;
	ucontext_t *ctx_actual;

	ctx_actual=&array_threads[pos_t_actual].ucp;
	array_threads[pos_t_actual].estado=2;
	hilos_activos--;
	hilos_suspendidos++;
	if(hilos_activos==0){
		fprintf(stderr,"No quedan threads activos\n" );
		exit(1);
	}
	pos_sig=buscarposicion();
	if (pos_sig==-1){
		for (int i=0;i<NUM_THREADS;i++){
			if (array_threads[i].estado==3){
				s=array_threads[i].hora_despertar-horams();
				if (s<=0){
					array_threads[i].estado=1;
				}
				sleep(s);
				array_threads[i].estado=1;
				pos_t_actual=i;
				swapcontext(ctx_actual,&array_threads[pos_t_actual].ucp);
				return;
			}
		}
	}else{

		pos_t_actual=pos_sig;
		swapcontext(ctx_actual,&array_threads[pos_t_actual].ucp);
	}
}


int
resumethread(int id){
	int pos;

	for (int i=0;i<NUM_THREADS;i++){
		if (array_threads[i].id==id){
			pos=i;
		}
	}

	if (array_threads[pos].estado==2){
		array_threads[pos].estado=1;
		hilos_suspendidos--;
		hilos_activos++;
		return 0;
	}
	return -1;
}

int 
suspendedthreads(int **list){
	int j=0;
	int *array_t_suspendidos;
	array_t_suspendidos=(int*)malloc(hilos_suspendidos*sizeof(int));

	for(int i=0;i<NUM_THREADS;i++){
		if(array_threads[i].estado==2){
			array_t_suspendidos[j]=array_threads[i].id;
			j++;
		}
	}

	if(trazas==1){
		for(int i=0;i<hilos_suspendidos; i++){
			printf("pos %d, id %d\n",i,array_t_suspendidos[i]);
		} 
	}

	*list=array_t_suspendidos;

	return hilos_suspendidos;
}

int 
killthread(int id){
	int pos=-1;

	for (int i=0;i<NUM_THREADS;i++){
		if (array_threads[i].id==id){
			pos=i;
		}
	}
	if(pos==pos_t_actual){
		fprintf(stderr,"Acción inválida. No puedes matarte a ti mismo\n");
		return -1;
	}else if (pos<0){
		fprintf(stderr,"Id inválido. No se puede eliminar un thread que no existe\n");
		return -1;
	}else{
		array_threads[pos].estado=0;
		array_threads[pos].pila_t=NULL;
		hilos_activos--;
		return id;
	}
}

void
sleepthread(int msec){
	int pos_sig;
	long long fin_dormir,hora_actual,s;
	ucontext_t *ctx_actual;

	ctx_actual=&array_threads[pos_t_actual].ucp;
	hora_actual=horams();
	fin_dormir=hora_actual+msec;
	array_threads[pos_t_actual].estado=3;
	array_threads[pos_t_actual].hora_despertar=fin_dormir;
	pos_sig=buscarposicion();
	
	if (pos_sig==-1){
		for (int i=0;i<NUM_THREADS;i++){
			if (array_threads[i].estado==3){
				s=array_threads[i].hora_despertar-horams();
				if (s<=0){
					array_threads[i].estado=1;
				}
				usleep(s);
				array_threads[i].estado=1;
				pos_t_actual=i;
				swapcontext(ctx_actual,&array_threads[pos_t_actual].ucp);
				return;
			}
		}
	}else{

		pos_t_actual=pos_sig;
		swapcontext(ctx_actual,&array_threads[pos_t_actual].ucp);
	}
}