package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("soma", soma(1,1))
	fmt.Println("subtracao" , subtracao(1,1))
	fmt.Println("multiplicacao" , multiplicacao(2,1))
	fmt.Println("divisao" , divisao(2,1))
	fmt.Println("adicionarpercentual" , adicionarpercentual(20 , 100))
	fmt.Println("descobrirpercentual" , descobrirpercentual(100 , 10))
	fmt.Println("pow" , pow(2, 4))
}
func soma(x int, y int) int {
	return x + y
}
func subtracao(x float64, y float64) float64{
	return x - y
}
func multiplicacao(x float64, y float64) float64{
	return x * y
}
func divisao(x float64, y float64) float64{
	return x / y
}
func adicionarpercentual(valor float64, porcentagem float64) float64{
	return valor + (valor * porcentagem / 100)
}
func descobrirpercentual(valor float64, valor2 float64) float64{
	return (valor2 / valor) *100
}
func pow(valor float64 , multiplo float64) float64{
	return math.Pow(valor,multiplo)
}