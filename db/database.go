package database

import (
	"database/sql"
	"log"

	"github.com/arturoZepeda/manejoDeCuentas/gasto"
)

type GastoDB struct {
	DB *sql.DB
}

const (
	file   string = "db/Gasto.db"
	create string = `CREATE TABLE IF NOT EXISTS Gasto(
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

func (g *GastoDB) GetGastoID(id int) (map[string]interface{}, error) {
	// r,err := g.DB.QueryRow(query string, args ...any)
	row := g.DB.QueryRow(`SELECT * FROM Gasto WHERE id = ?`, id)
	var idr int
	var descripcion, fecha, categoria string
	var monto float64
	err := row.Scan(&idr, &descripcion, &monto, &fecha, &categoria)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"id":          idr,
		"descripcion": descripcion,
		"monto":       monto,
		"fecha":       fecha,
		"categoria":   categoria,
	}, nil
}

func (g *GastoDB) NewGasto(gasto *gasto.Gasto) error {
	query := `INSERT INTO Gasto (descripcion, monto, fecha, categoria) VALUES (?, ?, ?, ?)`
	_, err := g.DB.Exec(query, gasto.Descripcion, gasto.Importe, gasto.FechaDeCompra.Format("2006-01-02"), gasto.Titular)
	if err != nil {
		return err
	}
	return nil
}

func (g *GastoDB) DeleteGastos() error {
	query := `DELETE FROM Gasto`
	_, err := g.DB.Exec(query)
	if err != nil {
		return err
	}
	_, err = g.DB.Exec(`DELETE FROM sqlite_sequence WHERE name ='Gasto'`)
	if err != nil {
		return err
	}
	return nil
}

func (g *GastoDB) DeleteGasto(id int) error {
	query := `DELETE FROM Gasto WHERE id = ?`
	_, err := g.DB.Exec(query, id)
	if err != nil {
		return err
	}

	// Reinicia el contador de autoincremento
	_, err = g.DB.Exec(`DELETE FROM sqlite_sequence WHERE name='Gasto'`)
	if err != nil {
		return err
	}

	return nil
}

func (g *GastoDB) UpdateGasto(id int, gasto *gasto.Gasto) error {
	query := `UPDATE Gasto SET descripcion = ?, monto = ?, fecha = ?, categoria = ? WHERE id = ?`
	_, err := g.DB.Exec(query, gasto.Descripcion, gasto.Importe, gasto.FechaDeCompra.Format("2006-01-02"), gasto.Titular, id)
	if err != nil {
		return err
	}
	return nil
}

func (g *GastoDB) Close() error {
	return g.DB.Close()
}
