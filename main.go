package main

import (
	"./multimedia"
	"fmt"
)


func main ()  {

	datosmultimedia := multimedia.ContenidoWeb{}
	var opcion int
	var nombre, formato, extra string

	fmt.Println("1.- Agregar Imagen")
	fmt.Println("2.- Agregar Audio")
	fmt.Println("3.- Agregar Video")
	fmt.Println("4.- Mostrar")

	fmt.Scanln(&opcion)

	switch opcion {

	case 1:

		fmt.Println("Dame el nombre: ")
		fmt.Scanln(&nombre)

		fmt.Println("Dame el formato: ")
		fmt.Scanln(&formato)

		fmt.Println("Dame el numero de canales: ")
		fmt.Scanln(&extra)

		img := multimedia.Imagen{nombre,formato,extra}
		datosmultimedia = append(img)
	case 2:

		fmt.Println("Dame el nombre: ")
		fmt.Scanln(&nombre)

		fmt.Println("Dame el formato: ")
		fmt.Scanln(&formato)

		fmt.Println("Dame el numero de duracion: ")
		fmt.Scanln(&extra)

		aud := multimedia.Audio{nombre,formato,extra}
		datosmultimedia = append(aud)

	case 3:

		fmt.Println("Dame el nombre: ")
		fmt.Scanln(&nombre)

		fmt.Println("Dame el formato: ")
		fmt.Scanln(&formato)

		fmt.Println("Dame el numero de frames: ")
		fmt.Scanln(&extra)

		vid := multimedia.Video{nombre,formato,extra}
		datosmultimedia = append(vid)
	case 4:
		
	}


	for _, contenido := range datosmultimedia.Slicemultimedia{
		contenido.Mostrar()
	}
	
}