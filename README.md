# Pila de Numeros Naturales

Ejercicio:

Diseñe una pila de números naturales con las siguientes operaciones:

- `push(n: int): int`

Que inserte un elemento en la cima de la pila y devuelve dicho elemento

- `pop(): int`

Que devuelva el último elemento de la pila y lo reitire de la misma

- `peekMin(): int`

que devuelva el menor elemento de la pila sin afectar su contenido

Criterios:

- La solución debe tener todas las operaciones que se indican arriba.

- Cada una de las operaciones en la solución debe tener una complejidad en tiempo de ejecución constante O(1)

- Se deben detectar e informar errores de underflow.


## Ejecutar el programa


``` sh
go run .
```

La salida será algo similar a lo siguiente
```sh
$ go run .

Apilado: 24
Apilado: 5
Apilado: 4
Desapilado: 4
Apilado: 3
Apilado: 2
Desapilado: 2
Apilado: 23
Apilado: 24
Desapilado: 24
Desapilado: 23
Apilado: 3
Desapilado: 3
Apilado: 25
Desapilado: 10
Apilado: 15
Apilado: 10
Apilado: 30

------------------
Minimo: 3
Stack: [5 15 24 25 3 30]
------------------
```

La impresión en pantalla de las operaciones no sigue un orden determiístico, ya que cada operacion de apilado o desapilado se realiza en un thread diferente.
Sin embargo, al finalizar todas las operaciones, se imprimen el valor minimo de la pila y el contenido final de la misma para verificar que el estado de la misma no posee inconsistencias, producto de la concurrencia.

## Complejidad 0(1)
La pila esta implementada sobre un slice de enteros, de manera que las operaciones `push` y `pop` no presentan dificultades para implementarse con tiempo de ejecusión constante, ya que se implementan como accesos al directos al indice tope de la pila.

Por otra parte, la operación que presenta problemas en la obtención del resultado en complejidad O(1) es `peekMin`, ya que para lograrlo, no es posible iterar la misma y comparar los valores. La solución propuesta, se sirve del uso de una pila de apoyo, donde se apilan los mínimos sucesivos.