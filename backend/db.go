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
	Id                  int       `json:"id,omitempty"`
	Name                string    `json:"name,omitempty"`
	Original_cost_price float32   `json:"original_cost_price,omitempty"`
	Unit_Cost_price     float32   `json:"unit_cost_price,omitempty"`
	Unit_Sell_price     float32   `json:"unit_sell_price,omitempty"`
	Category_id         int       `json:"category_id,omitempty"`
	Provider_id         int       `json:"provider_id,omitempty"`
	Quantity            int       `json:"quantity,omitempty"`
	Iva                 float32   `json:"iva,omitempty"`
	Internal_tax        float32   `json:"internal_tax,omitempty"`
	Created_at          time.Time `json:"created_at,omitempty"`
	Updated_at          time.Time `json:"updated_at,omitempty"`
}
type ProductWithCategoryAndProvider struct {
	Id                      int       `json:"id,omitempty"`
	Name                    string    `json:"name,omitempty"`
	Original_cost_price     float32   `json:"original_cost_price,omitempty"`
	Unit_Cost_price         float32   `json:"unit_cost_price,omitempty"`
	Unit_Sell_price         float32   `json:"unit_sell_price,omitempty"`
	Category_id             int       `json:"category_id,omitempty"`
	Provider_id             int       `json:"provider_id,omitempty"`
	Quantity                int       `json:"quantity,omitempty"`
	Iva                     float32   `json:"iva,omitempty"`
	Internal_Tax            float32   `json:"internal_tax,omitempty"`
	Category_Profit_Percent float32   `json:"category_profit_percent,omitempty"`
	Created_at              time.Time `json:"created_at,omitempty"`
	Updated_at              time.Time `json:"updated_at,omitempty"`
	Category_Name           string    `json:"category_name,omitempty"`
	Provider_Name           string    `json:"provider,omitempty"`
}

type Provider struct {
	Id           int       `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Phone_number string    `json:"phone_number,omitempty"`
	Created_at   time.Time `json:"created_at,omitempty"`
	Updated_at   time.Time `json:"updated_at,omitempty"`
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

	createProviderTableQuery := `
        CREATE TABLE IF NOT EXISTS provider (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            phone_number TEXT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
        INSERT OR IGNORE INTO provider(id, name, phone_number) VALUES (999, 'Ninguno', '0000-0000');
        `

	createCategoryTableQuery := `
        CREATE TABLE IF NOT EXISTS category (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            profit_percent REAL,
            CREATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            UPDATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
        INSERT OR IGNORE INTO category(id, name, profit_percent) VALUES (999, 'Ninguna', 1);
        `
	createProductTableQuery := `
        CREATE TABLE IF NOT EXISTS product (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            original_cost_price REAL NOT NULL,
            unit_cost_price REAL,
            unit_sell_price REAL,
            quantity INTEGER NOT NULL,
            iva REAL,
            internal_tax REAL,
            category_id INTEGER,
            provider_id INTEGER,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (category_id) REFERENCES category(id),
            FOREIGN KEY (provider_id) REFERENCES provider(id)
        )
        `
	createUpdateTrigger := `
        CREATE TRIGGER IF NOT EXISTS update_unit_prices_trigger
        AFTER UPDATE OF name, original_cost_price, quantity, iva, internal_tax, category_id, provider_id
        ON product
        FOR EACH ROW
        WHEN NEW.unit_sell_price IS NULL OR OLD.unit_sell_price = NEW.unit_sell_price
        BEGIN
            UPDATE product SET
                unit_cost_price = (NEW.original_cost_price + NEW.original_cost_price * NEW.iva / 100 + NEW.internal_tax) / NEW.quantity,
                unit_sell_price = (NEW.original_cost_price + NEW.original_cost_price * NEW.iva / 100 + NEW.internal_tax) / NEW.quantity * (1 + (SELECT profit_percent / 100 FROM category WHERE id = NEW.category_id))
            WHERE id = NEW.id;
        END;
        `
	createInsertTrigger := `
        CREATE TRIGGER IF NOT EXISTS product_insert_trigger
        AFTER INSERT ON product
        BEGIN
        UPDATE product SET
          unit_cost_price = (NEW.original_cost_price + NEW.original_cost_price * NEW.iva / 100 + NEW.internal_tax) / NEW.quantity,
          unit_sell_price = (NEW.original_cost_price + NEW.original_cost_price * NEW.iva / 100 + NEW.internal_tax) / NEW.quantity * (1 + (SELECT profit_percent / 100 FROM category WHERE id = NEW.category_id))
        WHERE id = NEW.id;
        END;
        `
	createPriceUpdateTriggerQuery := `
        CREATE TRIGGER IF NOT EXISTS update_sell_price_on_category_update
        AFTER UPDATE OF profit_percent ON category
        BEGIN
        UPDATE product SET 
            unit_sell_price = unit_sell_price / (1 + (OLD.profit_percent / 100)) * (1 + (NEW.profit_percent / 100)) WHERE category_id = NEW.id;
        END;
        `
	_, err = a.db.Exec(createProviderTableQuery)
	if err != nil {
		return fmt.Errorf("Error creating provider table: %w", err)
	}

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
	_, err = a.db.Exec(createInsertTrigger)
	if err != nil {
		return fmt.Errorf("Error creating product_insert_trigger: %w", err)
	}
	_, err = a.db.Exec(createUpdateTrigger)
	if err != nil {
		return fmt.Errorf("Error creating update_unit_prices_trigger: %w", err)
	}

	return nil
}
func (a *App) CreateProduct(product Product) error {
	query := `
    INSERT INTO product (name, original_cost_price, unit_sell_price, category_id, provider_id, quantity, iva, internal_tax) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
   `

	statement, err := a.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("Error preparing query: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(
		product.Name,
		product.Original_cost_price,
		product.Unit_Cost_price,
		product.Category_id,
		product.Provider_id,
		product.Quantity,
		product.Iva,
		product.Internal_tax,
	)
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
	err := a.db.QueryRow(query, id).Scan(&product.Id, &product.Name, &product.Original_cost_price, &product.Unit_Cost_price, &product.Unit_Sell_price, &product.Category_id, &product.Provider_id, &product.Quantity, &product.Iva)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Product with ID %d not found", id)
		}
		return nil, fmt.Errorf("Error getting product: %w", err)
	}

	return &product, nil
}

func (a *App) GetAllProducts() ([]ProductWithCategoryAndProvider, error) {
	query := `
    SELECT product.id, product.name, product.original_cost_price, product.unit_cost_price, product.unit_sell_price, 
       product.category_id, product.provider_id, product.quantity, product.iva, 
       product.internal_tax, 
       product.created_at, product.updated_at,
       category.name,
       category.profit_percent, provider.name
    FROM product 
    INNER JOIN category ON product.category_id = category.id 
    INNER JOIN provider ON product.provider_id = provider.id;
        `
	var products []ProductWithCategoryAndProvider
	rows, err := a.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error querying providers: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product ProductWithCategoryAndProvider
		if err := rows.Scan(&product.Id, &product.Name, &product.Original_cost_price, &product.Unit_Cost_price, &product.Unit_Sell_price, &product.Category_id, &product.Provider_id, &product.Quantity, &product.Iva, &product.Internal_Tax, &product.Created_at, &product.Updated_at, &product.Category_Name, &product.Category_Profit_Percent, &product.Provider_Name); err != nil {
			return nil, fmt.Errorf("Error scanning provider: %w", err)

		}
		products = append(products, product)
	}

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

func (a *App) UpdateProductById(id int, product Product) error {
	query := `
        UPDATE product SET 
            name = COALESCE(?, name),
            original_cost_price = COALESCE(?, original_cost_price),
            provider_id = COALESCE(?, provider_id),
            category_id = COALESCE(?, category_id),
            updated_at = CURRENT_TIMESTAMP,
            internal_tax = COALESCE(?, internal_tax),
            quantity = COALESCE(?, quantity),
            iva = COALESCE(?, iva)
        WHERE id = ?
        `
	statement, err := a.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("Error preparing SQL statement: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(
		product.Name,
		product.Original_cost_price,
		product.Provider_id,
		product.Category_id,
		product.Internal_tax,
		product.Quantity,
		product.Iva,
		id)
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

func (a *App) UpdateByProvider(percent_number float32, provider string) error {
	fmt.Printf("Received provider: %+v\n", provider)
	query := `UPDATE product 
                SET cost_price = cost_price * (1 + (? / 100)), updated_at = CURRENT_TIMESTAMP
                WHERE provider = ?`
	statement, err := a.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("Error preparing SQL statement: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(percent_number, provider)
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

func (a *App) UpdateCategoryById(id int, category Category) error {
	query := `UPDATE category SET name = COALESCE(?, name), profit_percent = COALESCE(?, profit_percent), updated_at = CURRENT_TIMESTAMP WHERE id = ?`
	statement, err := a.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("Error preparing SQL statement: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(category.Name, category.Profit_percent, id)
	if err != nil {
		return fmt.Errorf("Error executing SQL statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No category was updated")
	}
	fmt.Println("The category was updated")

	return nil
}

func (a *App) GetAllProviders() ([]Provider, error) {
	query := `SELECT * FROM provider ORDER BY id`
	var providers []Provider
	rows, err := a.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error querying providers: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var provider Provider
		if err := rows.Scan(&provider.Id, &provider.Name, &provider.Phone_number, &provider.Created_at, &provider.Updated_at); err != nil {
			return nil, fmt.Errorf("Error scanning provider: %w", err)

		}
		providers = append(providers, provider)
	}

	return providers, nil
}

func (a *App) CreateProvider(provider Provider) error {
	query := `INSERT INTO provider(name, phone_number) VALUES (?, ?)`
	statement, err := a.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("Error preparing SQL statement: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(provider.Name, provider.Phone_number)
	if err != nil {
		return fmt.Errorf("Error executing SQL statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No provider was created")
	}
	fmt.Println("The provider was created")

	return nil

}

func (a *App) UpdateProviderById(id int, provider Provider) error {
	query := `UPDATE provider SET name = COALESCE(?, name), phone_number = COALESCE(?, phone_number), updated_at = CURRENT_TIMESTAMP WHERE id = ?`
	statement, err := a.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("Error preparing SQL statement: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(provider.Name, provider.Phone_number, id)
	if err != nil {
		return fmt.Errorf("Error executing SQL statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No provider was updated")
	}
	fmt.Println("The provider was updated")

	return nil
}

func (a *App) DeleteProviderById(id int) error {
	query := `DELETE FROM provider WHERE id = ?`
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
		return fmt.Errorf("No provider was deleted")
	}

	fmt.Println("The provider was deleted.")
	return nil
}
