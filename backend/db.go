package backend

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Product struct {
	Id         int       `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Cost_price float32   `json:"cost_price,omitempty"`
	Provider   string    `json:"provider,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
	Updated_at time.Time `json:"updated_at,omitempty"`
}

func (a *App) Database() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get User Config Directory: %w", err)
	}
	configFolder := filepath.Join(configDir, "productsusi")
	if err := os.MkdirAll(configFolder, 0755); err != nil {
		return fmt.Errorf("error creating folder: %w", err)
	}
	a.db, err = sql.Open("sqlite3", configFolder+"/test.db")
	if err != nil {
		return fmt.Errorf("error opening database connection: %w", err)
	}
	if err = a.db.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}
	_, err = a.db.Exec(`
        CREATE TABLE IF NOT EXISTS products (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            cost_price REAL NOT NULL,
            provider TEXT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `)
	if err != nil {
		return fmt.Errorf("error creating the database table: %w", err)
	}
	return nil
}
func (a *App) CreateProduct(product Product) {
	insertSql := `INSERT INTO products(name, cost_price, provider) VALUES (?, ?, ?)`
	statement, err := a.db.Prepare(insertSql)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(product.Name, product.Cost_price, product.Provider)
	if err != nil {
		log.Fatal("Couldn't create the product", err)
	} else {
		fmt.Println("The product was created")
	}
}
func (a *App) GetProductById(id int) Product {
	var product Product
	query := `SELECT * FROM products WHERE id = ?`
	err := a.db.QueryRow(query, id).Scan(&product.Id, &product.Provider, &product.Updated_at, &product.Created_at, &product.Name, &product.Cost_price)
	if err != nil {
		log.Fatal("Couldn'n get the product", err)
	}
	return product
}
func (a *App) GetAllProducts() []Product {
	var products []Product
	rows, err := a.db.Query("SELECT * FROM products ORDER BY id")
	if err != nil {
		return []Product{}
	}
	defer rows.Close()
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Cost_price, &product.Provider, &product.Created_at, &product.Updated_at); err != nil {
			log.Fatal("Couldn't query it sorry", err)
			return []Product{}
		}
		products = append(products, product)
	}
	return products
}
func (a *App) DeleteProduct(id int) {
	query := `DELETE FROM products WHERE id = ?`
	statement, err := a.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	result, err := statement.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rowsAffected == 0 {
		fmt.Println("No product was deleted.")
	} else {
		fmt.Println("The product was deleted.")
	}
}

func (a *App) UpdateProduct(id int, product Product) {
	query := `UPDATE products SET name = COALESCE(?, name), cost_price = COALESCE(?, cost_price), provider = COALESCE(?, provider), updated_at = ? WHERE id = ?`
	statement, err := a.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	result, err := statement.Exec(product.Name, id, product.Cost_price, product.Updated_at, product.Provider)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rowsAffected == 0 {
		fmt.Println("No product was updated.")
	} else {

		fmt.Println("The product was updated.")
	}

}
