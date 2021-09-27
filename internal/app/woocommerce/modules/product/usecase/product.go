package usecase

import (
	"go-project-structure-example/internal/app/custom_types"
	"go-project-structure-example/internal/app/models"
	"go-project-structure-example/internal/app/woocommerce/modules/product/repository"

	"github.com/bxcodec/faker/v3"
)

type ProductUseCase interface {
	GetProductByID(id uint, fields []string) (models.Product, error)
	GetProducts(paginate custom_types.Paginate) ([]models.Product, error)
	GenerateProducts(n int) ([]models.Product, error)
}

type productImpl struct {
	productRepo repository.ProductRepository
}

func NewProductUseCase(pr repository.ProductRepository) ProductUseCase {
	return &productImpl{
		productRepo: pr,
	}
}

func (this *productImpl) GetProductByID(id uint, fields []string) (models.Product, error) {
	product, err := this.productRepo.FindProductByID(id, fields)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (this *productImpl) GetProducts(paginate custom_types.Paginate) ([]models.Product, error) {
	products, err := this.productRepo.GetProductsWithPaginate(paginate)
	if err != nil {
		return products, err
	}
	return products, nil
}

func (this *productImpl) GenerateProducts(n int) ([]models.Product, error) {
	productFakeData := make([]interface{}, 0)
	for i := 0; i < n; i++ {
		p := struct {
			Title string `json:"title" faker:"name"`
		}{}
		_ = faker.FakeData(&p)

		productFakeData = append(productFakeData, p)
	}

	products, err := this.productRepo.BulkInsertProducts(productFakeData)

	if err != nil {
		return products, err
	}
	return products, nil
}
