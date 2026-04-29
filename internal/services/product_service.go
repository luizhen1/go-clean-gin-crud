package services

import (
	"github.com/luizhen1/go-clean-gin-crud/internal/domain"
)

type ProductService struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *domain.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) FetchProducts() ([]domain.Product, error) {
	return s.repo.Fetch()
}

func (s *ProductService) GetProductByID(id int64) (*domain.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) UpdateProduct(product *domain.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id int64) error {
	return s.repo.Delete(id)
}
