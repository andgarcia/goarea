package goarea

import "math"

//Pi é uma proporção numerica definida pela relacao entre o perimetro de uma circunferencia e o seu diametro
const Pi = 3.14116

//Circ é responsavel por calcular a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect é responsavel por calcular a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Triangulo é a funcao publica para calcular o triangulo
func Triangulo(base, altura float64) float64 {
	return _TrianguloEq(base, altura)
}

//não é visivel por causa do _
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
