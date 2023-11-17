package main

import "fmt"

func main() {
	var base, limite, producto int

	fmt.Print("Ingrese la base de la tabla: ")
	fmt.Scan(&base)
	fmt.Print(("Ingrese el limite de la tabla: "))
	fmt.Scan(&limite)
	fmt.Println("La tabla del ", base)
	
	for contador := 0; contador <= limite; contador = contador + 1 {
		producto = base * contador
		fmt.Println(contador, "*", base, "=", producto)
	}
}
