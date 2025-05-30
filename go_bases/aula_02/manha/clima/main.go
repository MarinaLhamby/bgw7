package main

import "fmt"

// Utilizei todos como inteiro por levar em consideração o display utilizado em celulares
func main() {
	var (
		temperatura int = 19
		umidade     int = 60
		pressao     int = 1025
	)

	fmt.Printf("Previsão do tempo atual:\nTemperatura: %d graus celsius\nUmidade relativa do ar: %d%%\nPressão atmosférica: %d hPa\n",
		temperatura, umidade, pressao)
}
