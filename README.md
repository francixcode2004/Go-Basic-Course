# Curso Básico de : Conceptos Fundamentales

Este repositorio contiene un ejemplo de código en que cubre los conceptos básicos del lenguaje. Está diseñado para ayudar a los principiantes a comprender los fundamentos de , incluyendo variables, constantes, operadores, condicionales, bucles, arrays y más.

## Contenido

1. [Variables y Tipos de Datos](#variables-y-tipos-de-datos)
2. [Constantes](#constantes)
3. [Operadores Aritméticos](#operadores-aritméticos)
4. [Operadores Relacionales](#operadores-relacionales)
5. [Operadores Lógicos](#operadores-lógicos)
6. [Operadores Bit a Bit](#operadores-bit-a-bit)
7. [Condicionales](#condicionales)
8. [Bucles](#bucles)
9. [Arrays](#arrays)
10. [Arrays Multidimensionales](#arrays-multidimensionales)

[Variables y Tipos de Datos](#variables-y-tipos-de-datos)
`Variables y Tipos de Datos`
var msg string
var num int
var desicion bool
var pointer \*string
msg = "Hola mundo"
fmt.Println(msg)
fmt.Println(num)
fmt.Println(desicion)
fmt.Println(pointer)

`Constantes`
Las constantes se declaran utilizando la palabra clave const y son útiles para valores que no cambian durante la ejecución del programa:

const PI = math.Phi
fmt.Printf("type %T value %v\n", PI, PI)

`Operadores Aritméticos`
Los operadores aritméticos se utilizan para realizar operaciones matemáticas básicas:

Suma (+)
Resta (-)
Multiplicación (_)
División (/)
Módulo (%)
x := 4
y := 10
fmt.Println(x + y)
fmt.Println(x - y)
fmt.Println(x _ y)
fmt.Println(float64(x) / float64(y))
fmt.Println(x % y)

`Operadores Relacionales`
Los operadores relacionales comparan dos valores y devuelven un resultado booleano:

Mayor que (>)
Menor que (<)
Mayor o igual que (>=)
Menor o igual que (<=)
Igualdad (==)
Desigualdad (!=)

fmt.Println(x > y)
fmt.Println(x < y)
fmt.Println(x >= y)
fmt.Println(x <= y)
fmt.Println(x == y)
fmt.Println(x != y)

`Operadores Lógicos`
Los operadores lógicos se utilizan para combinar expresiones booleanas:

AND (&&)
OR (||)
Negación (!)

fmt.Println(x > y && y < x)
fmt.Println(x > y || y < x)
fmt.Println(!true)
fmt.Println(!false)

`Operadores Bit a Bit`
Los operadores bit a bit permiten manipular los bits individuales de un número:

AND bit a bit (&)
OR bit a bit (|)
XOR bit a bit (^)
NOT bit a bit (^) (complemento)
Desplazamiento a la izquierda (<<)
Desplazamiento a la derecha (>>)

a := 12 // 1100 en binario
b := 10 // 1010 en binario

fmt.Println("a & b =", a&b) // 8
fmt.Println("a | b =", a|b) // 14
fmt.Println("a ^ b =", a^b) // 6
fmt.Println("^a =", ^a) // -13
fmt.Println("a << 1 =", a<<1) // 24
fmt.Println("a >> 1 =", a>>1) // 6
`Condicionales`
Los condicionales permiten ejecutar diferentes bloques de códi en función de ciertas condiciones:

if x == y {
fmt.Print("x es igual a y")
} else if x != y {
fmt.Print("x no es igual a y")
} else {
fmt.Print("x es igual a y")
}
`Bucles`
El bucle for es el único tipo de bucle en y se utiliza para repetir un bloque de códi varias veces:

for i := 0; i < 10; i++ {
fmt.Println("Valor de i dentro del loop: ", i)
}
`Arrays`
Los arrays en son colecciones de elementos del mismo tipo, con un tamaño fijo:

arr := [4]int{1, 2, 3, 4}
fmt.Println(arr)
fmt.Println(arr[0]) // Valor en el índice 0
fmt.Println(len(arr)) // Tamaño del array
fmt.Println(cap(arr)) // Capacidad del array
Arrays Multidimensionales
Los arrays multidimensionales son arrays que contienen otros arrays como elementos:

arr2 := [2][2]int{{1, 2}, {3, 4}}
for i := 0; i < len(arr2); i++ {
for j := 0; j < len(arr2[i]); j++ {
fmt.Println(arr2[i][j])
}
}
`Contribución`
Este es un ejemplo básico para aquellos que están comenzando con . Si deseas contribuir o agregar más ejemplos, siéntete libre de crear un Pull Request.

Licencia
Este proyecto está licenciado bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.

Este `README.md` proporciona una explicación detallada de cada parte del códi y los conceptos básicos que cubre, lo que es ideal para principiantes que aprenden .
