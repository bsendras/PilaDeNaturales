package main

import (
	"errors"
	"sync"
)

// Stack representa una pila
type Stack struct {
	mutex sync.Mutex
	items []int
	mins  []int
}

// Push inserta un elemento al tope de la pila
func (s *Stack) Push(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("Only snatural numbers are allowed")
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.items = append(s.items, n)
	s.handleInsertMin(n)

	return n, nil
}

// Pop remueve el tope de la pila y devuelve el valor almacenado
func (s *Stack) Pop() (int, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	count := len(s.items)
	if count <= 0 {
		return 0, errors.New("Stack is empty")
	}

	n := s.items[count-1]
	s.items = s.items[:count-1]
	s.handleRemoveMin(n)

	return n, nil
}

// PeekMin retorna el valor minimo de la pila en tiempo constante O(1)
func (s *Stack) PeekMin() (int, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	count := len(s.mins)

	if count <= 0 {
		return 0, errors.New("Stack is empty")
	}

	return s.mins[count-1], nil
}

// GetItems devuelve una copia de todos los elementos de la pila
// En este caso, acceder a los elementos de la pila, constituye
// un anti patrón con fines demostrativos.
func (s *Stack) GetItems() []int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.items
}

// Funciones no exportadas. Si "stack" fuese un package
// separado, estas funciones no serían accesibles desde
// el exterior del mismo
func (s *Stack) handleInsertMin(n int) {
	count := len(s.mins)
	if count <= 0 {
		s.mins = append(s.mins, n)
		return
	}

	min := s.mins[count-1]
	if n <= min {
		s.mins = append(s.mins, n)
	}
}

func (s *Stack) handleRemoveMin(n int) {
	count := len(s.mins)
	if count <= 0 {
		return
	}

	min := s.mins[count-1]
	if n == min {
		s.mins = s.mins[:count-1]
	}
}

// NewStack crea una nueva instancia de la pila
func NewStack() *Stack {
	return &Stack{
		items: []int{},
		mins:  []int{},
	}
}
