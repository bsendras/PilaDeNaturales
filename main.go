package main

import (
	"fmt"
	"sync"
)

/*
	Apilo y desapilo números de forma asincrónica para que el resultado no
	sea deterministico. Solo desapilo durante las iteraciones impares para que al
	finalizar aún queden elementos en la pila. De esta manera puedo verificar
	que tanto el contenido de la pila como el valor minimo registrado, quedan
	en un estado consistente y que las primitivas de la pila son thread safe.
*/

func main() {
	var wg sync.WaitGroup

	// Valores
	var items []int = []int{5, 3, 10, 15, 25, 30, 4, 3, 2, 23, 24, 24}

	// Creo la pila
	stack := NewStack()

	// Apilo y desapilo
	count := len(items)
	for i := 0; i < count; i++ {
		go Push(items[i], stack, &wg)
		if i%2 == 0 {
			go Pop(stack, &wg)
		}
	}

	// Espero las go routines
	wg.Wait()

	// Veo los resultadods
	fmt.Println("")
	fmt.Println("------------------")
	Peek(stack)
	ShowItems(stack)
	fmt.Println("------------------")
}
