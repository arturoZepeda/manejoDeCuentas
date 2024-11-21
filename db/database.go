package database

import (
	"database/sql"
	"log"
)

type GastoDB struct {
	DB *sql.DB
}

const (
	file   string = "Gasto.db"
	create string = `CREATE TABLE Gasto(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  descripcion TEXT NOT NULL,
  monto REAL NOT NULL,
  fecha TEXT NOT NULL,
  categoria TEXT NOT NULL);`
)

func Conectar() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Conexion exitosa a la base de datos: ", file)
	return db, nil
}

func NewGastos() (*GastoDB, error) {
	db, err := Conectar()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(create)
	if err != nil {
		db.Close()
		return nil, err
	}

	log.Println("Base de datos creada exitosamente.")
	return &GastoDB{DB: db}, nil
}

func (g *GastoDB) GetGastos() ([]map[string]interface{}, error) {
	rows, err := g.DB.Query(`SELECT * FROM Gasto`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var gastos []map[string]interface{}
	for rows.Next() {
		var id int
		var descripcion, fecha, categoria string
		var monto float64
		if err := rows.Scan(&id, &descripcion, &monto, &fecha, &categoria); err != nil {
			return nil, err
		}
		gastos = append(gastos, map[string]interface{}{
			"id":          id,
			"descripcion": descripcion,
			"monto":       monto,
			"fecha":       fecha,
			"categoria":   categoria,
		})
	}
	return gastos, nil
}
