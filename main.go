package main

import (
	"fmt"

	"github.com/arturoZepeda/manejoDeCuentas/gasto"
	leercsv "github.com/arturoZepeda/manejoDeCuentas/leerCSV"
)

func main() {
	// var lineas []gasto.Gasto

	gasto, err := gasto.New("2024-11-17", "Titular", "Descripcion", 100000.01)
	if err != nil {
		fmt.Println(err)
	}
	gasto.Output()
	gasto.Clear()
	gasto.Output()
	lineasTemp, err := leercsv.LeeCsv("activity.csv")
	if err != nil {
		fmt.Println(err)
	}
	esAmex := leercsv.EsAmex(lineasTemp[0])
	fmt.Println(esAmex)
	for i, linea := range lineasTemp {
		fmt.Println("[", i, "]", linea)
		fechaTemp := linea[0]
		titular := linea[3]
		descripcion := linea[2]
		importe := linea[5]
		fmt.Println(fechaTemp, titular, descripcion, importe)
		// lineas = append(lineas,  )
	}
}
