package woocommerce

import (
	"ecommerce-integrations/internal/app/models"

	"github.com/bxcodec/faker/v3"

	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

func createShop(db *gorm.DB) error {
	fakeData := make([]interface{}, 0)
	numberOfFakeData := 10
	for i := 0; i < numberOfFakeData; i++ {
		p := struct {
			DefaultDomain string `json:"default_domain" gorm:"size:50" faker:"username"`
			ShopifyDomain string `json:"shopify_domain" gorm:"size:50" faker:"username"`
		}{}
		_ = faker.FakeData(&p)

		p.DefaultDomain = p.DefaultDomain + ".myshopify.com"
		p.ShopifyDomain = p.ShopifyDomain + ".myshopify.com"

		fakeData = append(fakeData, p)
	}

	shops := []models.Shop{}
	mapstructure.Decode(fakeData, &shops)

	err := db.Create(&shops).Error
	if err != nil {
		return err
	}

	return nil
}
