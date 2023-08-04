package main

import (
	"fmt"
	"math/rand"
	"time"
)

// funcion de ejecucion
func main() {
	Lista_listas := []int{100, 1000} //, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 50000}
	var tiempo_bubble_sort, tiempo_Counting_sort time.Duration

	var s_b []string
	//var s_m []int
	//var s_s []int
	//var s_c []int

	//t := time.Now().UTC()

	for i := 0; i < len(Lista_listas); i++ {
		Lista_de_lista := make([]int, Lista_listas[i]) /// creo una lista de ceros con tamaño de cada valor de la lista_listas {0,0,0,0,0}
		fmt.Println(len(Lista_de_lista))
		contador_1 := 0

		for contador_1 < 5 {
			contador_1++

			for i_1 := 0; i_1 < len(Lista_de_lista); i_1++ {
				Lista_de_lista[i_1] = rand.Intn(10) /// mete valores rand a la lista de ceros
			}

			//fmt.Println(Lista_de_lista)
			tiempo_bubble_sort = Bubble_sort(Lista_de_lista)
			//fmt.Println(Lista_de_lista, "lista ordena")
			fmt.Printf("%v\n", tiempo_bubble_sort)
			ra := tiempo_bubble_sort.String()
			s_b = append(s_b, ra)

		}

		contador_3 := 0

		for contador_3 < 5 {
			contador_3++

			for i_1 := 0; i_1 < len(Lista_de_lista); i_1++ {
				Lista_de_lista[i_1] = rand.Intn(10) //mete valores rand a la lista de ceros
			}

			//fmt.Println(Lista_de_lista)
			t_start_1 := time.Now()
			mergeSort(Lista_de_lista)
			t_end_1 := time.Now()
			tiempo_ejec_merge := t_end_1.Sub(t_start_1)
			///fmt.Println(ordenada_merge, "lista ordena")
			fmt.Printf("%v\n", tiempo_ejec_merge)
			//s_m = append(s_m, tiempo_ejec_merge)

		}

		contador_2 := 0

		for contador_2 < 5 {
			contador_2++

			for i_1 := 0; i_1 < len(Lista_de_lista); i_1++ {
				Lista_de_lista[i_1] = rand.Intn(10) //mete valores rand a la lista de ceros
			}

			//fmt.Println(Lista_de_lista)
			t_start := time.Now()
			Selection_Sort(Lista_de_lista, len(Lista_de_lista))
			t_end := time.Now()
			tiempo_ejec_selection := t_end.Sub(t_start)
			//fmt.Println(ordenada_select, "lista ordena")
			fmt.Printf("%v\n", tiempo_ejec_selection)
			//s_s = append(s_s, tiempo_ejec_selection)
		}

		contador_4 := 0

		for contador_4 < 5 {
			contador_4++

			for i_1 := 0; i_1 < len(Lista_de_lista); i_1++ {
				Lista_de_lista[i_1] = rand.Intn(10) //mete valores rand a la lista de ceros
			}

			//fmt.Println(Lista_de_lista)
			tiempo_Counting_sort = Counting_sort(Lista_de_lista)
			//fmt.Println(Lista_de_lista, "lista ordena")
			fmt.Printf("%v\n", tiempo_Counting_sort)
			//s_c = append(s_c, tiempo_Counting_sort)
		}
	}
	fmt.Println(s_b)
}

// ALGORITMOS
func Bubble_sort(Lista []int) time.Duration {
	//var ListaB = make([]int, len(Lista)) // Declaras una lista con tun tamano X
	//copy(ListaB, Lista)                  // copias en ListaB lo de Lista
	//fmt.Println("Burbuja1")
	//fmt.Println(Lista)
	//fmt.Println(ListaB)
	t_start := time.Now() //Inicio de reloj
	// Algoritmo
	swapped := 0 //
	for i := 0; i < len(Lista)-1; i++ {
		for j := 0; j < len(Lista)-i-1; j++ {
			if Lista[j] > Lista[j+1] {
				swapped = 1
				Lista[j], Lista[j+1] = Lista[j+1], Lista[j]
			}
		}
		if swapped == 0 {
			return 0
		}
	}
	t_end := time.Now()
	//fmt.Println(ListaB)
	return (t_end.Sub(t_start))
}

// Selection_sort ordenacion por burbuja
func Selection_Sort(array []int, size int) []int {
	var min_index int
	var temp int
	for i := 0; i < size-1; i++ {
		min_index = i
		// Find index of minimum element
		for j := i + 1; j < size; j++ {
			if array[j] < array[min_index] {
				min_index = j
			}
		}
		temp = array[i]
		array[i] = array[min_index]
		array[min_index] = temp
	}
	return array
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}
func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

// Counting_sort ordenacion por burbuja
func Counting_sort(Lista []int) time.Duration {
	ListaC := Lista[:]
	//fmt.Println("Counting")
	//fmt.Println(ListaC)
	t_start := time.Now() //Inicio de reloj
	// Algoritmo
	size := len(ListaC)
	N := 10
	output := make([]int, size) // [0]*size
	count := make([]int, N)     //[0]*N
	for i := 0; i < size; i++ {
		count[ListaC[i]] += 1
	}
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	//	i := size - 1
	for i := 0; i < size; i++ {
		output[count[ListaC[i]]-1] = ListaC[i]
		count[ListaC[i]] -= 1
		//i -= 1
	}
	for i := 0; i < size; i++ { // in range(0, size):
		ListaC[i] = output[i]
	}
	t_end := time.Now()
	//fmt.Println(ListaC)
	return (t_end.Sub(t_start))
}
