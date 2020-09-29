package main

import (
	"fmt"
	"sync"
)

// Push apila un item
func Push(n int, s *Stack, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	value, err := s.Push(n)
	if err == nil {
		fmt.Println("Apilado:", value)
	}
}

// Pop desapila un item
func Pop(s *Stack, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	value, err := s.Pop()
	if err == nil {
		fmt.Println("Desapilado:", value)
	}
}

// Peek obtiene y muestra el minimo
func Peek(s *Stack) {
	value, err := s.PeekMin()
	if err == nil {
		fmt.Println("Minimo:", value)
	}
}

// ShowItems Muestra los items de la pila
func ShowItems(s *Stack) {
	fmt.Println("Stack:", s.GetItems())
}
