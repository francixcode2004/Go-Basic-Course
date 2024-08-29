package main

import (
	"fmt"
	"math"
)

func main() {
	// variables
	var msg string
	var num int
	var desicion bool
	var pointer *string
	msg = "Hola mundo"
	fmt.Println(msg)
	fmt.Println(num)
	fmt.Println(desicion)
	fmt.Println(pointer)

	//constantes
	const PI = math.Phi
	//tipo de dato $T
	// valor de la variable/constante %v
	// siempre con Printf para el formateador
	fmt.Printf("type %T value %v", PI, PI)

	//operadores aritmeticos e inferencia de datos
	x := 4
	y := 10
	//suma
	fmt.Println(x + y)
	//resta
	fmt.Println(x - y)
	//multiplicacion
	fmt.Println(x * y)
	//division
	//unboxing conversion de datos
	fmt.Println(float64(x) / float64(y))
	//porciento
	fmt.Println(float64(x) * float64(y) / 100)
	//divisor
	fmt.Println(x % y)

	//asignacion de operacion
	//x++, x--, x+=y, x*=y, x/=y
	x += y
	fmt.Println(x)

	//operadores relacionales
	//mayor que >, menor que <, mayor o igual >=, menor o igual <=
	//igualdad ==, desigualdad !=
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
	fmt.Println(x == y)
	fmt.Println(x != y)

	//operadores logicos
	//&& and
	// || or
	// ! negacion
	fmt.Println(x > y && y < x)
	fmt.Println(x > y || y < x)
	fmt.Println(!true)
	fmt.Println(!false)

	// operadores bit a bit
	a := 12 // 1100 en binario
	b := 10 // 1010 en binario

	// AND bit a bit
	fmt.Println("a & b =", a&b) // 1000 en binario, que es 8 en decimal

	// OR bit a bit
	fmt.Println("a | b =", a|b) // 1110 en binario, que es 14 en decimal

	// XOR bit a bit
	fmt.Println("a ^ b =", a^b) // 0110 en binario, que es 6 en decimal

	// NOT bit a bit (complemento)
	fmt.Println("^a =", ^a) // -13 en decimal

	// Desplazamiento a la izquierda
	fmt.Println("a << 1 =", a<<1) // 24 en decimal

	// Desplazamiento a la derecha
	fmt.Println("a >> 1 =", a>>1) // 6 en decimal

	//condicionales
	if x == y {
		fmt.Print("x es igual a y")
	} else if x!=y {
		fmt.Print("x no es igual a y")
	}else{
		fmt.Print("x es igual a y")
	}

	//bucle for
	for i:=0;i<10;i++{
		fmt.Println("Valor de i dentro del loop: ",i)
	}

	//creacio de arrays
	arr:= [4]int{1,2,3,4}
	fmt.Println(arr)
	//valor del array en el indice 0 
	fmt.Println(arr[0])
	// tamaño del array
	fmt.Println(len(arr))
	// tipo de dato del array
	fmt.Println(cap(arr))


	//array multidimensionales
	arr2:= [2][2]int{{1,2},{3,4}}
	for i:=0;i<len(arr2);i++{
		for j:=0;j<len(arr2[i]);j++{
			fmt.Println(arr2[i][j])
		}
	}




	

}
