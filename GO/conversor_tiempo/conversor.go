package main

import (
	"fmt"
)

// Convierte horas a minutos y segundos
func convertirHoras(horas float64) (float64, float64) {
	minutos := horas * 60
	segundos := horas * 3600
	return minutos, segundos
}

// Convierte minutos a horas y segundos
func convertirMinutos(minutos float64) (float64, float64) {
	horas := minutos / 60
	segundos := minutos * 60
	return horas, segundos
}

// Convierte segundos a horas y minutos
func convertirSegundos(segundos float64) (float64, float64) {
	horas := segundos / 3600
	minutos := segundos / 60
	return horas, minutos
}

func main() {
	fmt.Println("=== CONVERSOR DE TIEMPO === ")

	for {
		fmt.Println("1. Convertir horas a minutos y segundos")
		fmt.Println("2. Convertir minutos a horas y segundos")
		fmt.Println("3. Convertir segundos a horas y minutos")
		fmt.Println("4. Salir")

		var opcion int
		fmt.Println("Selecciona una opción: ")
		fmt.Scan(&opcion)

		if opcion == 4 {
			fmt.Println("Saliendo del programa...")
			break
		}

		switch opcion {
		case 1:
			var horas float64
			fmt.Print("Introduce las horas: ")
			fmt.Scan(&horas)
			min, seg := convertirHoras(horas)
			fmt.Printf("%.2f horas son %.2f minutos o %.2f segundos.\n", horas, min, seg)

		case 2:
			var minutos float64
			fmt.Print("Introduce los minutos: ")
			fmt.Scan(&minutos)
			horas, seg := convertirMinutos(minutos)
			fmt.Printf("%.2f minutos son %.2f horas o %.2f segundos.\n", minutos, horas, seg)

		case 3:
			var segundos float64
			fmt.Print("Introduce los segundos: ")
			fmt.Scan(&segundos)
			horas, min := convertirSegundos(segundos)
			fmt.Printf("%.2f segundos son %.2f horas o %.2f minutos.\n", segundos, horas, min)

		default:
			fmt.Println("Opción no válida. Inténtalo de nuevo.")
		}

	}
}
