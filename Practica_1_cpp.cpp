#include <iostream>
#include <list> // for list operations
#include <bits/stdc++.h>
#include <ctime> 
#include "Bubble_sort.cpp"
#include "dmeas_algoritm.cpp"
#include "Counting_sort.cpp"
#include "merge_sort.cpp"
#include <algorithm>
#include <sys/time.h>
#include <iomanip>



using namespace std;

int main()
{  
	
	int lista_grande[15] = {100,1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 50000}; //// Esta lista, cada valor es la cantidad de numeros que debe tener una lista para ordenarla 5 veces
	int N_l_g = sizeof(lista_grande) / sizeof(lista_grande[0]); /// NLG servira para hacer la iteracion por indice en el sgte bucle for
	
	for (int ii=0; ii<N_l_g ; ii++  )   // un bucle for que es para cada valor de la lista grande, creara una lista random de una cantidad de numeros igual a ese valor
	{
		int n = lista_grande[ii];       // n es el tamaño de la lista, 100 1000 2000 .... 50000, toma el numero de acuerdo al indice indicado lista_grande[0] es 100, lista_grande[15] es 50000
		cout<< n << endl;               // printeamos el tamaño de la losta para verificar
	

		
		int arr[n];  ///  declaramos la variable lista array con la cantidad de numeros que va a tener de acuerdo a cada valor de la lista grande  1000  2000 ... 50000
	 
		srand(time(NULL));
		
		int j = 1;                               /// un contador para el bucle while, debe correrse 5 veces cada uno
		
		while(j<=5)        //// codigo tiene que ejecutarse 5 veces
		{
			for(int i=0;i<n;i++)	{		int num = 0+ rand()%(50-0);		arr[i]=num;	}   // para crear una lista de randoms
			//for(int i=0;i<n;i++)		cout<<arr[i]<<" ";	cout<<endl;  // printeamos la lista de randoms creada
			j+=1;
			//cout<< n << endl;
			unsigned t0, t1;
			
			
			t0=clock();
			// Bubble Sort		
			bubbleSort(arr,n);
					
			t1 = clock();
						
			double time_0 = (double(t1-t0)/CLOCKS_PER_SEC);
			cout << setprecision(5);
			cout << time_0 << endl;		};
		
		int j_1 = 1;
		while(j_1<=5){
		
				
			j_1+=1;
			//cout<< n << endl;
			//for(int i=0;i<n;i++)		cout<<arr[i]<<" ";	cout<<endl; // printeamos la lista de randoms ordenada	
			
			for(int i=0;i<n;i++)	{		int num = 1+ rand()%(50-1);		arr[i]=num;	} 
			//for(int i=0;i<n;i++)		cout<<arr[i]<<" ";	cout<<endl; // printeamos la lista de randoms nueva
			double t0_0=clock();
			// MergeSort
			mergeSort(arr,0,n-1);
			double  t1_0 = clock();
			double time_1 = (double(t1_0-t0_0) / CLOCKS_PER_SEC);
			setprecision(5) ;
			cout << time_1 << endl;
			//for(int i=0;i<n;i++)		cout<<arr[i]<<" ";	cout<<endl; // printeamos la lista de randoms ordenada	
		};
			
		int j_2 = 1;	
		while(j_2<=5){
		
				
			j_2+=1;
			//cout<< n << endl;
			
			for(int i=0;i<n;i++)	{		int num = 1+ rand()%(50-1);		arr[i]=num;	} 
			unsigned t0_1=clock();
			// Selection Sort
			selectionSort(arr,n);
			unsigned t1_1 = clock();
			double time_2 = (double(t1_1-t0_1)/CLOCKS_PER_SEC);
			cout << setprecision(5) ;
			cout << time_2 << endl;};
			
		int j_3 = 1;	
		while(j_3<=5){
		
				
			j_3+=1;
			//cout<< n << endl;
				
			for(int i=0;i<n;i++)	{		int num = 1+ rand()%(50-1);		arr[i]=num;	} 
			unsigned t0_2=clock();
			// Counting Srot
			countSort(arr,n);
			unsigned t1_2 = clock();
			double time_3 = (double(t1_2-t0_2)/CLOCKS_PER_SEC);
			cout << setprecision(5) ;
			cout << time_3 << endl;};
		}	
		return 0;
	}

