package http

import (
	"go-project-structure-example/driver"
	"go-project-structure-example/internal/app/custom_types"
	"go-project-structure-example/internal/app/woocommerce/modules/product/repository"
	"go-project-structure-example/internal/app/woocommerce/modules/product/usecase"
	"go-project-structure-example/internal/utils/app"
	"go-project-structure-example/internal/utils/e"
	"go-project-structure-example/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	GenerateProductsFaker(c *gin.Context)
	GetProductList(c *gin.Context)
	GetProductById(c *gin.Context)
}

type productHandler struct {
	productUseCase usecase.ProductUseCase
}

func NewProductHandler(db *driver.Database) ProductHandler {
	pRepo := repository.NewProductRepository(db)
	pUseCase := usecase.NewProductUseCase(pRepo)
	return &productHandler{
		productUseCase: pUseCase,
	}
}

func (this *productHandler) GetProductById(c *gin.Context) {
	app := app.Gin{C: c}

	idParam := c.Param("id")

	id, err := helpers.StringToUint(idParam)

	if err != nil {
		app.BadRequest(gin.H{
			"status_code": e.INVALID_PARAMS,
		})
		return
	}

	product, err := this.productUseCase.GetProductByID(id, nil)

	if err != nil {
		app.BadRequest(gin.H{
			"status_code": e.INVALID_PARAMS,
		})
		return
	}
	app.Ok((gin.H{
		"product": product,
	}))
}

func (this *productHandler) GetProductList(c *gin.Context) {
	app := app.Gin{C: c}

	offsetQuery := c.DefaultQuery("offset", "0")
	limitQuery := c.DefaultQuery("limit", "10")

	offset, err := helpers.StringToInt(offsetQuery)

	if err != nil {
		app.BadRequest(gin.H{
			"status_code": e.INVALID_PARAMS,
		})
		return
	}

	limit, err := helpers.StringToInt(limitQuery)

	if err != nil {
		app.BadRequest(gin.H{
			"status_code": e.INVALID_PARAMS,
		})
		return
	}

	paginate := custom_types.Paginate{
		Offset: offset,
		Limit:  limit,
	}

	products, err := this.productUseCase.GetProducts(paginate)

	if err != nil {
		app.BadRequest(gin.H{
			"status_code": e.INVALID_PARAMS,
		})
		return
	}

	app.Ok((gin.H{
		"products": products,
	}))
}

func (this *productHandler) GenerateProductsFaker(c *gin.Context) {
	app := app.Gin{C: c}

	number := c.DefaultQuery("n", "10")

	n, err := helpers.StringToInt(number)

	if err != nil {
		app.BadRequest(gin.H{
			"status_code": e.INVALID_PARAMS,
		})
		return
	}

	products, err := this.productUseCase.GenerateProducts(n)
	if err != nil {
		app.Error(err.Error())
		return
	}

	app.Ok(gin.H{
		"products": products,
	})
}
