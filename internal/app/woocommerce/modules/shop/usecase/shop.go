package usecase

import (
	"ecommerce-integrations/internal/app/models"
	"ecommerce-integrations/internal/app/woocommerce/modules/shop/repository"

	"github.com/bxcodec/faker/v3"
)

type ShopUseCase struct {
	shopRepo repository.ShopRepository
}

func NewShopUseCase(pr repository.ShopRepository) ShopUseCase {
	return ShopUseCase{
		shopRepo: pr,
	}
}

func (instance *ShopUseCase) GenerateShops(n int) ([]models.Shop, error) {
	shopFakeData := make([]interface{}, 0)
	for i := 0; i < n; i++ {
		p := struct {
			Title string `json:"title" faker:"name"`
		}{}
		_ = faker.FakeData(&p)

		shopFakeData = append(shopFakeData, p)
	}

	shops, err := instance.shopRepo.BulkInsertShops(shopFakeData)

	if err != nil {
		return shops, err
	}
	return shops, nil
}
