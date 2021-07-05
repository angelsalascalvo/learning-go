package main

import (
	"fmt"
	"strconv" //Parseos string
)

/*
El uso de paquetes implica que las funciones y variables declaradas con la primera letra en mayuscula sean publicas (accesibles entre paquetes)
y con minuscula solo accesibles desde el propio paquete
*/
func main() {
	//Parte1()
	//Parte2()
	//Parte3()
	//Parte4()
}

//------------------------------------------------------------>

// Parte1 :
//Ejemplos basicos 1
func Parte1() {
	//1. Declarar variables
	var a1 = "hola mundo"
	var a2 string = "hola mundo 2"
	a3 := "hola mundo 3"

	var b1, b2, b3 int
	b1 = 1
	b2 = 2
	b3 = 3

	//2. Imprimir
	fmt.Println(a1)
	fmt.Println(a2 + " - " + a3)

	fmt.Println("B1 = " + strconv.Itoa(b1)) //Hay que convertirlo a string para poder imprimirlo con +
	fmt.Println("B2 =", b2, "a")            //No es necesario parsearlo a string si se imprime con ,
	fmt.Printf("B3 = %d", b3)               //Se pueden emplear verbos para imprimir variables dentro de las cadenas con formato (printf) https://golang.org/pkg/fmt/#hdr-Printing

	//3. Leer datos
	var edad int
	fmt.Println("\nIndica tu edad")
	//Uso de verbos (%d) para parsear directamente al datos que deseamos, el \n lo utilizamos para evitar leer el enter pulsado por el usuario
	fmt.Scanf("%d\n", &edad) // &edad empleado para seleccionar la referencia a memoria
	fmt.Println("Has introducido:", edad)

	var nombre string
	fmt.Println("\nIndica tu nombre")
	fmt.Scanf("%s\n", &nombre)
	fmt.Println("Has introducido: " + nombre)

	//4.Condicionales
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")

	} else if edad < 0 {
		fmt.Println("¿Aun no has nacido?")

	} else {
		fmt.Println("Eres menor de edad")
	}

	if nombre == "angel" && edad >= 18 {
		fmt.Println("Hola angel mayor de edad")

	} else if nombre == "angel" {
		fmt.Println("Hola angel")
	}
}

//------------------------------------------------------------>

//Parte2 :
//Ejemplos basicos 2
func Parte2() {
	//1. Bucles, unicamente existe el ciclo for
	for i := 0; i < 10; i++ {
		fmt.Println("Hola mundo", i)
	}

	//Simulacion de while con el bucle for
	i := 0
	for i < 10 {
		fmt.Println("Simulando un while")
		i++
	}

	//Bucle infinito
	for {
		i++
		if i == 20 {
			break
		}
	}
	fmt.Print("\nSe ha roto el bucle infinito :(\n\n")

	//Bucle saltando iteraciones
	j := 0
	for j < 10 {
		j++
		if j == 3 {
			continue
		}
		//Si se ejecuta el continue se salta la iteracion y esta parte no se ejecuta
		fmt.Println(j)
	}
}

//------------------------------------------------------------>

func Parte3() {
	//1. Array
	var array [10]int
	fmt.Println(array)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Print(array2)
	fmt.Print(" Logitud: ", len(array2), "\n")
	fmt.Println("Primera posicion: " + strconv.Itoa(array2[0]))

	//2. Slices  (array dinamicos)
	var slice []int //Definicion igual que un array pero sin indicar el tamanno
	if slice == nil {
		//Al inicializar un slice si inicia a nulo
		fmt.Println("Está vacio, tamaño " + strconv.Itoa(len(slice)))
	}

	//De este modo podemos crear un slice con 3 elementos y una capacidad 4 (que es el tamanno maximo del slice)
	slice2 := make([]int, 3, 4)
	fmt.Println(slice2)
	fmt.Println("Longitud actual: ", len(slice2), " Capacidad máxima: ", cap(slice2))
	slice2 = append(slice2, -1) //agregar un nuevo dato al slice de forma dinamica
	fmt.Println(slice2)
	//Si superamos la capacidad establecida inicialmente, append redimensionara el slice dandole más capacidad (Operacion menos optima que la asignacion)
	slice2 = append(slice2, -8)
	fmt.Println(slice2)
	fmt.Print("Capacidad ampliada hasta ", cap(slice2), "\n")

	//3. Copiar slices
	var sliceOrigin = []int{9, 8, 7, 6}
	var sliceCopy = make([]int, 3)

	copy(sliceCopy, sliceOrigin)

	fmt.Println("\n", sliceCopy)

	//4. Punteros
	var puntero1, puntero2 *int //Variables que hacen referencia a un espacio de memoria
	entero := 5

	//Con & obtenemos el puntero correspondiente a una variable
	//(entero => valor de la variable, &entero => espacio de memoria donde se almacena el valor de la variable)
	puntero1 = &entero

	fmt.Println(puntero1)
	fmt.Println(puntero2)

	//Alterar el valor de un espacio de memoria
	//(puntero => espacio de memoria, *puntero => valor del espacio de memoria)
	*puntero1 = 6

	fmt.Println("Valor asociadodel puntero ", *puntero1, "(Direccion de memoria ", puntero1, ")")
	fmt.Println("Valor de entero ", entero, "(Direccion de memoria ", &entero, ")")

}

//------------------------------------------------------------>

type User struct {
	edad             int
	nombre, apellido string
}

func (user User) GetNombreCompleto() string {
	return user.nombre + " " + user.apellido
}

//Si en el metodo set hacemos referencia al puntero de user podremos cambiar el valor,
//en otro caso lo que se envía al metodo es una copia de la estructura y cualquier cambio solo afecta al ambito del pripio metodo
func (user *User) SetName(nuevoNombre string) {
	user.nombre = nuevoNombre
}

type Animal struct {
	nombre string
}
type Colibri struct {
	Animal      //Campo anonimo
	colorPlumas string
}

func Parte4() {
	//1.Estructuras
	var angel User
	angel.edad = 21
	angel.nombre = "angel"
	angel.apellido = "salas"
	fmt.Println(angel)

	tobias := User{11, "Tobias", "Dupty"}
	fmt.Println(tobias)
	fmt.Println(tobias.GetNombreCompleto())

	manolo := User{nombre: "Manolo", edad: 72}
	fmt.Println(manolo)

	manolo.SetName("Juan")
	fmt.Println(manolo.nombre)

	//2. Campos anonimos en estructuras (simulacion de herencia)
	colibri := Colibri{Animal{"colibri"}, "amarillo y negro"}
	fmt.Println(colibri.nombre)

}
