package usecase

import (
	"ecommerce-integrations/internal/app/custom_types"
	"ecommerce-integrations/internal/app/models"
	"ecommerce-integrations/internal/app/woocommerce/modules/product/repository"

	"github.com/bxcodec/faker/v3"
)

type ProductUseCase struct {
	productRepo repository.ProductRepository
}

func NewProductUseCase(pr repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		productRepo: pr,
	}
}

func (instance *ProductUseCase) GetProductByID(id uint, fields []string) (models.Product, error) {
	product, err := instance.productRepo.FindProductByID(id, fields)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (instance *ProductUseCase) GetProducts(paginate custom_types.Paginate) ([]models.Product, error) {
	products, err := instance.productRepo.GetProductsWithPaginate(paginate)
	if err != nil {
		return products, err
	}
	return products, nil
}

func (instance *ProductUseCase) GenerateProducts(n int) ([]models.Product, error) {
	productFakeData := make([]interface{}, 0)
	for i := 0; i < n; i++ {
		p := struct {
			Title string `json:"title" faker:"name"`
		}{}
		_ = faker.FakeData(&p)

		productFakeData = append(productFakeData, p)
	}

	products, err := instance.productRepo.BulkInsertProducts(productFakeData)

	if err != nil {
		return products, err
	}
	return products, nil
}
