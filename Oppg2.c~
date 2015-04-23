//gcc -std=gnu99 -Wall -g -o Oppg2 Oppg2.c -lpthread

#include <pthread.h>
#include <stdio.h>


int i = 0;
int N = 100;


// Note the return type: void*
void* Inkrementer(void* lock){
    for (int j = 0; j < N; j++){
	pthread_mutex_lock(lock);
    	i+=2;
	pthread_mutex_unlock(lock);
    }
    return NULL;
}

void* Dekrementer(void* lock){
    for (int j = 0; j < N; j++){
    	pthread_mutex_lock(lock);
    	i--;
	pthread_mutex_unlock(lock);
    }
    return NULL;
}



int main(){
    pthread_mutex_t lock;
    pthread_mutex_init(&lock,NULL);
    
    pthread_t someThread;
    pthread_t someThread2;
    pthread_create(&someThread, NULL, Inkrementer,(void*) &lock);
    pthread_create(&someThread2, NULL, Dekrementer,(void*) &lock);
    
    pthread_join(someThread, NULL);
    pthread_join(someThread2, NULL);
    printf("%d\n", i);

    pthread_mutex_destroy(&lock);
    return 0;
    
}

