package http

import (
	"ecommerce-integrations/driver"
	"ecommerce-integrations/internal/app/woocommerce/modules/shop/repository"
	"ecommerce-integrations/internal/app/woocommerce/modules/shop/usecase"
	"ecommerce-integrations/internal/utils/app"
	"ecommerce-integrations/internal/utils/constants"
	"ecommerce-integrations/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type ShopHandler struct {
	shopUseCase usecase.ShopUseCase
}

func NewShopHandler(db *driver.Database) ShopHandler {
	tRepo := repository.NewShopRepository(db)
	pUseCase := usecase.NewShopUseCase(tRepo)
	return ShopHandler{
		shopUseCase: pUseCase,
	}
}
func (instance *ShopHandler) Index(c *gin.Context) {
	app := app.Gin{C: c}

	number := c.DefaultQuery("n", "10")

	n, err := helpers.StringToInt(number)

	if err != nil {
		app.BadRequest(gin.H{
			"status_code": constants.INVALID_PARAMS,
		})
		return
	}

	shops, err := instance.shopUseCase.GenerateShops(n)
	if err != nil {
		app.Error(err.Error())
		return
	}

	app.Ok(gin.H{
		"shops": shops,
	})
}
