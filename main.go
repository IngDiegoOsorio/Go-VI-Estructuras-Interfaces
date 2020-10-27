package main

import "./multimedia"
func main ()  {

	img1 := multimedia.Imagen{"imagen1","jpg","4"}
	aud1 := multimedia.Audio{"Audio1","Mp3","10.5"}

	img1.Mostrar()
	aud1.Mostrar()
}