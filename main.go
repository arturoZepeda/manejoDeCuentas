package main

import (
	"fmt"
	"log"
	"strconv"

	database "github.com/arturoZepeda/manejoDeCuentas/db"
	"github.com/arturoZepeda/manejoDeCuentas/extas"
	"github.com/arturoZepeda/manejoDeCuentas/gasto"
	leercsv "github.com/arturoZepeda/manejoDeCuentas/leerCSV"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	/*===================Proceso BBDD=================*/
	//Inicializamos base de datos.
	gastoDB, err := database.NewGastos()
	if err != nil {
		log.Fatalf("Error al inicializar la base de datos: %v", err)
	}
	defer gastoDB.DB.Close()
	/*
		gastos, err := gastoDB.GetGastos()
		if err != nil {
			log.Fatalf("Error al obtener los gastos: %v", err)
		}

		fmt.Println("Gastos en la BBDD: ")
		for _, g := range gastos {
			fmt.Printf("ID: %v, Descripción: %v, Monto: %v, Fecha: %v, Categoría: %v\n", g["id"], g["descripcion"], g["monto"], g["fecha"], g["categoria"])
		}
		idTemp := 1
		gastoID, err := gastoDB.GetGastoID(idTemp)
		if err != nil {
			log.Fatalf("Error al obtener el gato %v descripcion %v", idTemp, err)
		}
		fmt.Printf("ID: %v, Descripción: %v, Monto: %v, Fecha: %v, Categoría: %v\n", gastoID["id"], gastoID["descripcion"], gastoID["monto"], gastoID["fecha"], gastoID["categoria"])

		idTemp = 2
		err = gastoDB.DeleteGasto(idTemp)
		if err != nil {
			log.Fatal("Error al borrar el registro.")
		} else {
			fmt.Println("Se ha eliminado correctamente el registo")
		}
	*/
	/*==================Proceso CSV==================*/
	// genera la lectura del CSV activity y lo almacena en una matriz
	lineasTemp, err := leercsv.LeeCsv("activity.csv")
	if err != nil {
		fmt.Println(err)
	}

	// validamos  si es un amex al verificar la cabecera de los documentos.
	esAmex := leercsv.EsAmex(lineasTemp[0])
	fmt.Println(esAmex)
	// recorremos la matriz para leer cada una de las entradas
	for i, linea := range lineasTemp {

		fmt.Println("========================================")
		fmt.Println("[", i, "]")
		// Asignación de los valores para generar en el nuevo struct.
		fecha, err := extas.ParseFechaEspanol(linea[0])
		if err != nil {
			fmt.Println("error al convertir la fecha")
		}
		titular := linea[3]
		descripcion := linea[2]
		importeStr := linea[5]
		importe, err := strconv.ParseFloat(importeStr, 64)
		if err != nil {
			fmt.Println("Problemas al parsear el importe.", err)
		}
		fmt.Println(fecha, titular, descripcion, importe)
		// generamos nuevo struct
		gastoTemp, err := gasto.New(fecha.Format("2006-01-02"), titular, descripcion, importe)
		if err != nil {
			fmt.Println("Error al crear el gasto:", err)
			continue
		}
		// gastoTemp.Output()
		err = gastoDB.NewGasto(gastoTemp)
		if err != nil {
			fmt.Printf("Error al dar de alta el gasto: %v", err)
		}
	}
	/*
			err = gastoDB.DeleteGastos()
			if err != nil {
				fmt.Printf("Error al limpiar la bbdd: %v", err)
			}
				err = gastoDB.DeleteGasto(390)
				if err != nil {
					fmt.Printf("Error al eliminar el registro ? el error es: ?",390, err)
				}

		err = gastoDB.UpdateCalificador(2, 5)
		if err != nil {
			fmt.Printf("Error al actualizar el calificador: %v", err)
		}
	*/

	gastoTemo, err := gastoDB.GetGastoByCalificador(5)
	if err != nil {
		fmt.Println("Error al obtener los registros: ?", err)
	}
	fmt.Println("Gastos en la BBDD por calificador")
	for _, g := range gastoTemo {
		fmt.Printf("ID: %v, Descripción: %v, Monto: %v, Fecha: %v, Calificador: %v\n", g["id"], g["descripcion"], g["monto"], g["fecha"], g["calificador"])
	}
}
