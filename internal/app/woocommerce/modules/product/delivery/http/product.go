package http

import (
	"ecommerce-integrations/driver"
	"ecommerce-integrations/internal/app/custom_types"
	"ecommerce-integrations/internal/app/woocommerce/modules/product/repository"
	"ecommerce-integrations/internal/app/woocommerce/modules/product/usecase"
	"ecommerce-integrations/internal/utils/app"
	"ecommerce-integrations/internal/utils/constants"
	"ecommerce-integrations/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUseCase usecase.ProductUseCase
}

func NewProductHandler(db *driver.Database) ProductHandler {
	pRepo := repository.NewProductRepository(db)
	pUseCase := usecase.NewProductUseCase(pRepo)
	return ProductHandler{
		productUseCase: pUseCase,
	}
}

func (instance *ProductHandler) GetProductById(c *gin.Context) {
	app := app.Gin{C: c}

	idParam := c.Param("id")

	id, err := helpers.StringToUint(idParam)

	if err != nil {
		app.BadRequest(gin.H{
			"status_code": constants.INVALID_PARAMS,
		})
		return
	}

	product, err := instance.productUseCase.GetProductByID(id, nil)

	if err != nil {
		app.BadRequest(gin.H{
			"status_code": constants.INVALID_PARAMS,
		})
		return
	}
	app.Ok((gin.H{
		"product": product,
	}))
}

func (instance *ProductHandler) GetProductList(c *gin.Context) {
	app := app.Gin{C: c}

	offsetQuery := c.DefaultQuery("offset", "0")
	limitQuery := c.DefaultQuery("limit", "10")

	offset, err := helpers.StringToInt(offsetQuery)

	if err != nil {
		app.BadRequest(gin.H{
			"status_code": constants.INVALID_PARAMS,
		})
		return
	}

	limit, err := helpers.StringToInt(limitQuery)

	if err != nil {
		app.BadRequest(gin.H{
			"status_code": constants.INVALID_PARAMS,
		})
		return
	}

	paginate := custom_types.Paginate{
		Offset: offset,
		Limit:  limit,
	}

	products, err := instance.productUseCase.GetProducts(paginate)

	if err != nil {
		app.BadRequest(gin.H{
			"status_code": constants.INVALID_PARAMS,
		})
		return
	}

	app.Ok((gin.H{
		"products": products,
	}))
}

func (instance *ProductHandler) GenerateProductsFaker(c *gin.Context) {
	app := app.Gin{C: c}

	number := c.DefaultQuery("n", "10")

	n, err := helpers.StringToInt(number)

	if err != nil {
		app.BadRequest(gin.H{
			"status_code": constants.INVALID_PARAMS,
		})
		return
	}

	products, err := instance.productUseCase.GenerateProducts(n)
	if err != nil {
		app.Error(err.Error())
		return
	}

	app.Ok(gin.H{
		"products": products,
	})
}
