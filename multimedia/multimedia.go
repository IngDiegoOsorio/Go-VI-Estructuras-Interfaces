package multimedia

import "fmt"

type Multimedia interface {
	Mostrar()
}

type Imagen struct{
	Titulo string
	Formato string
	Canales string
}

func (img *Imagen) Mostrar() {
	atributos := "Titulo: " + img.Titulo + "Formato: " + img.Formato + "Canales: " + img.Canales
	fmt.Println(atributos)
}

type Audio struct{
	Titulo string
	Formato string
	Duracion string
}

func (aud *Audio) Mostrar() {
	atributos := "Titulo: " + aud.Titulo + "Formato: " + aud.Formato + "Canales: " + aud.Duracion
	fmt.Println(atributos)
}

type Video struct{
	Titulo string
	Formato string
	Frames string
}	

func (vid *Video) Mostrar() {
	atributos := "Titulo: " + vid.Titulo + "Formato: " + vid.Formato + "Canales: " + vid.Frames
	fmt.Println(atributos)
}