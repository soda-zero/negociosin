package backend

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id           int       `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Cost_price   float32   `json:"cost_price,omitempty"`
	Sell_price   float32   `json:"sell_price,omitempty"`
	Provider     string    `json:"provider,omitempty"`
	Category_id  int       `json:"category_id,omitempty"`
	Quantity     int       `json:"quantity,omitempty"`
	Iva          float32   `json:"iva,omitempty"`
	Internal_tax float32   `json:"internal_tax,omitempty"`
	Created_at   time.Time `json:"created_at,omitempty"`
	Updated_at   time.Time `json:"updated_at,omitempty"`
}
type ProductWithCategory struct {
	Id                      int       `json:"id,omitempty"`
	Name                    string    `json:"name,omitempty"`
	Cost_price              float32   `json:"cost_price,omitempty"`
	Sell_price              float32   `json:"sell_price,omitempty"`
	Provider                string    `json:"provider,omitempty"`
	Category_id             int       `json:"category_id,omitempty"`
	Quantity                int       `json:"quantity,omitempty"`
	Iva                     float32   `json:"iva,omitempty"`
	Internal_tax            float32   `json:"internal_tax,omitempty"`
	Category_Name           string    `json:"category_name,omitempty"`
	Category_Profit_percent float32   `json:"category_profit_percent,omitempty"`
	Created_at              time.Time `json:"created_at,omitempty"`
	Updated_at              time.Time `json:"updated_at,omitempty"`
}

type Provider struct {
	Name string `json:"name,omitempty"`
}

type Category struct {
	Id             int       `json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	Profit_percent float32   `json:"profit_percent,omitempty"`
	Created_at     time.Time `json:"created_at,omitempty"`
	Updated_at     time.Time `json:"updated_at,omitempty"`
}

func (a *App) Database() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("Failed to get User Config Directory: %w", err)
	}
	configFolder := filepath.Join(configDir, "productsusi")
	if err := os.MkdirAll(configFolder, 0755); err != nil {
		return fmt.Errorf("Error creating folder: %w", err)
	}

	a.db, err = sql.Open("sqlite3", configFolder+"/test.db")
	if err != nil {
		return fmt.Errorf("Error opening database connection: %w", err)
	}

	if err = a.db.Ping(); err != nil {
		return fmt.Errorf("Error pinging database: %w", err)
	}

	createCategoryTableQuery := `
        CREATE TABLE IF NOT EXISTS category (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            profit_percent REAL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
        `
	createProductTableQuery := `
        CREATE TABLE IF NOT EXISTS product (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            cost_price REAL NOT NULL,
            sell_price REAL,
            provider TEXT,
            category_id INTEGER,
            quantity INTEGER,
            iva REAL,
            internal_tax REAL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (category_id) REFERENCES category(id)
        );
        `

	createPriceUpdateTriggerQuery := `
        CREATE TRIGGER update_sell_price
            AFTER INSERT ON product
            BEGIN
                UPDATE product SET cost_price = cost_price + (cost_price * (iva / 100)) + internal_tax WHERE id = NEW.id;
                UPDATE product SET cost_price = cost_price / quantity WHERE id = NEW.id;
                UPDATE product SET sell_price = cost_price * (1 + (SELECT profit_percent / 100 FROM category WHERE category.id = NEW.category_id)) WHERE id = NEW.id;
        END;

        CREATE TRIGGER update_sell_price_on_category_update
            AFTER UPDATE OF profit_percent ON category
            BEGIN
                UPDATE product SET sell_price = sell_price / (1 + (OLD.profit_percent / 100)) * (1 + (NEW.profit_percent / 100)) WHERE category_id = NEW.id;
        END;
        `
	_, err = a.db.Exec(createCategoryTableQuery)
	if err != nil {
		return fmt.Errorf("Error creating category table: %w", err)
	}

	_, err = a.db.Exec(createProductTableQuery)
	if err != nil {
		return fmt.Errorf("Error creating product table: %w", err)
	}
	_, err = a.db.Exec(createPriceUpdateTriggerQuery)
	if err != nil {
		return fmt.Errorf("Error creating price update trigger query: %w", err)

	}

	return nil
}
func (a *App) CreateProduct(product Product) error {
	query := `INSERT INTO product(name, cost_price, provider, category_id, quantity, iva, internal_tax) VALUES (?, ?, ?, ?, ?, ?, ?)`
	statement, err := a.db.Prepare(query)
	defer statement.Close()

	result, err := statement.Exec(product.Name, product.Cost_price, product.Provider, product.Category_id, product.Quantity, product.Iva, product.Internal_tax)
	if err != nil {
		return fmt.Errorf("Error executing query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error getting row affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No product was created")
	}

	fmt.Println("The product was created")
	return nil
}
func (a *App) GetProductById(id int) (*Product, error) {
	var product Product
	query := `SELECT * FROM product WHERE id = ?`
	err := a.db.QueryRow(query, id).Scan(&product.Id, &product.Provider, &product.Updated_at, &product.Created_at, &product.Name, &product.Cost_price, &product.Quantity, &product.Iva, &product.Internal_tax)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Product with ID %d not found", id)
		}
		return nil, fmt.Errorf("Error getting product: %w", err)
	}

	return &product, nil
}

func (a *App) GetAllProducts() ([]ProductWithCategory, error) {
	query := `SELECT p.*, c.name, c.profit_percent
                FROM product p
              LEFT JOIN category c ON p.category_id = c.id;`
	var products []ProductWithCategory
	rows, err := a.db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Error querying providers: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product ProductWithCategory
		if err := rows.Scan(&product.Id, &product.Name, &product.Cost_price, &product.Sell_price, &product.Provider, &product.Category_id, &product.Quantity, &product.Iva, &product.Internal_tax, &product.Created_at, &product.Updated_at, &product.Category_Name, &product.Category_Profit_percent); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("Error scanning provider: %w", err)

		}
		products = append(products, product)
	}

	fmt.Println(products)
	return products, nil
}

func (a *App) DeleteProductById(id int) error {
	query := `DELETE FROM product WHERE id = ?`
	statement, err := a.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("Error preparing SQL statement: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(id)
	if err != nil {
		return fmt.Errorf("Error executing SQL statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No product was deleted")
	}

	fmt.Println("The product was deleted.")
	return nil
}

func (a *App) UpdateProduct(id int, product Product) error {
	query := `UPDATE product SET name = COALESCE(?, name), cost_price = COALESCE(?, cost_price), provider = COALESCE(?, provider), updated_at = CURRENT_TIMESTAMP WHERE id = ?`
	statement, err := a.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("Error preparing SQL statement: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(product.Name, product.Cost_price, product.Provider, id)
	if err != nil {
		return fmt.Errorf("Error executing SQL statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No product was updated")
	}
	fmt.Println("The product was updated")

	return nil
}
func (a *App) GetAllProviders() ([]Provider, error) {
	query := `SELECT provider FROM product ORDER BY id`
	var providers []Provider
	rows, err := a.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error querying providers: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var provider Provider
		if err := rows.Scan(&provider.Name); err != nil {
			return nil, fmt.Errorf("Error scanning provider: %w", err)

		}
		providers = append(providers, provider)
	}

	return providers, nil
}

func (a *App) UpdateByProvider(value float32, provider string) error {
	fmt.Printf("Received provider: %+v\n", provider)
	query := `UPDATE product 
                SET cost_price = cost_price * (1 + (? / 100)), updated_at = CURRENT_TIMESTAMP
                WHERE provider = ?`
	statement, err := a.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("Error preparing SQL statement: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(value, provider)
	if err != nil {
		return fmt.Errorf("Error executing SQL statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No product where updated by Provider")
	}

	fmt.Println("The product were updated by Provider.")
	return nil

}

func (a *App) GetAllCategories() ([]Category, error) {
	query := `SELECT * FROM category ORDER BY id`
	var categories []Category
	rows, err := a.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error querying categories: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.Id, &category.Name, &category.Profit_percent, &category.Created_at, &category.Updated_at); err != nil {
			return nil, fmt.Errorf("Error scanning category: %w", err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (a *App) CreateCategory(category Category) error {
	query := `INSERT INTO category(name, profit_percent) VALUES(?,?)`
	statement, err := a.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("Error preparing SQL statement: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(category.Name, category.Profit_percent)
	if err != nil {
		return fmt.Errorf("Error executing SQL statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error getting row affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No category was created")
	}

	fmt.Println("The category was created")
	return nil
}

func (a *App) DeleteCategoryById(id int) error {
	query := `DELETE FROM category WHERE id = ?`
	statement, err := a.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("Error preparing SQL statement: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(id)
	if err != nil {
		return fmt.Errorf("Error executing SQL statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error getting row affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No category was deleted")
	}

	fmt.Println("The category was deleted")
	return nil
}
