package multimedia

import "fmt"

type imagen struct{
	Titulo string
	Formato string
	Canales string
}

func (img *imagen) Mostrar() {
	atributos := "Titulo: " + img.Titulo + "Formato: " + img.Formato + "Canales: " + img.Canales
	fmt.Println(atributos)
}

type audio struct{
	Titulo string
	Formato string
	Duracion string
}

func (aud *audio) Mostrar() {
	atributos := "Titulo: " + aud.Titulo + "Formato: " + aud.Formato + "Canales: " + aud.Duracion
	fmt.Println(atributos)
}

type video struct{
	Titulo string
	Formato string
	Frames string
}	

func (vid *video) Mostrar() {
	atributos := "Titulo: " + vid.Titulo + "Formato: " + vid.Formato + "Canales: " + vid.Frames
	fmt.Println(atributos)
}