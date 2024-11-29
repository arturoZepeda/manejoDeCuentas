package serv

import (
	"encoding/json"
	"net/http"

	"github.com/arturoZepeda/manejoDeCuentas/database"
	"github.com/goji/goji"
	"github.com/goji/goji/mux"
)

func SetupRoutes(db *database.GastoDB) *goji.Mux {
	// Crea un router Goji
	router := goji.NewMux()

	// Define rutas
	router.HandleFunc(mux.Pattern("/api/gastos"), func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleGetGastos(db, w, r)
		} else if r.Method == "POST" {
			handleNewGasto(db, w, r)
		}
	})

	router.HandleFunc(mux.Pattern("/api/gastos/:id"), func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleGetGastoByID(db, w, r)
		} else if r.Method == "DELETE" {
			handleDeleteGasto(db, w, r)
		}
	})

	return router
}

// Controladores
func handleGetGastos(db *database.GastoDB, w http.ResponseWriter, r *http.Request) {
	gastos, err := db.GetGastos()
	if err != nil {
		http.Error(w, "Error al obtener los gastos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gastos)
}

func handleNewGasto(db *database.GastoDB, w http.ResponseWriter, r *http.Request) {
	var gasto map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&gasto)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	// Aquí puedes crear un nuevo registro en la base de datos usando los datos enviados
	w.WriteHeader(http.StatusCreated)
}
