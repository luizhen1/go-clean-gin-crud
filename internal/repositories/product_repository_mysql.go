package repositories

import (
	"database/sql"

	"github.com/luizhen1/go-clean-gin-crud/internal/domain"
)

type mysqlProductRepository struct {
	db *sql.DB
}

// NewMySQLProductRepository cria uma instância do repositório
func NewMySQLProductRepository(db *sql.DB) domain.ProductRepository {
	return &mysqlProductRepository{db: db}
}

func (r *mysqlProductRepository) Create(product *domain.Product) error {
	query := "INSERT INTO products (name, price) VALUES (?, ?)"

	// Executamos a query passando os valores da struct
	result, err := r.db.Exec(query, product.Name, product.Price)
	if err != nil {
		return err
	}

	// Recuperamos o ID gerado pelo MySQL e atribuímos de volta à struct
	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	product.ID = lastID

	return nil
}

// Por enquanto, vamos deixar os outros métodos vazios (stubs) para satisfazer a interface
func (r *mysqlProductRepository) Fetch() ([]domain.Product, error) {
	query := "SELECT id, name, price FROM products"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product

	// Percorremos os resultados do banco
	for rows.Next() {
		var product domain.Product
		// Fazemos o scan de cada coluna para a nossa struct
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *mysqlProductRepository) GetByID(id int64) (*domain.Product, error) {
	query := "SELECT id, name, price FROM products WHERE id = ?"

	var product domain.Product
	// QueryRow executa a query e já permite fazer o Scan diretamente
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Retorna nada se o produto não existir
		}
		return nil, err
	}

	return &product, nil
}

func (r *mysqlProductRepository) Update(product *domain.Product) error {
	query := "UPDATE products SET name = ?, price = ? WHERE id = ?"

	_, err := r.db.Exec(query, product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *mysqlProductRepository) Delete(id int64) error {
	query := "DELETE FROM products WHERE id = ?"

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
