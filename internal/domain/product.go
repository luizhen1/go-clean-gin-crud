package domain

// Product representa a nossa entidade de negócio
// @Description Informações do produto
// @name Product
type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductRepository interface {
	Create(product *Product) error
	Fetch() ([]Product, error)
	GetByID(id int64) (*Product, error)
	Update(product *Product) error
	Delete(id int64) error
}
