package http

import (
	"ecommerce-integrations/driver"
	"ecommerce-integrations/internal/app/woocommerce/modules/webhook/repository"
	"ecommerce-integrations/internal/app/woocommerce/modules/webhook/usecase"
	"ecommerce-integrations/internal/utils/app"
	"ecommerce-integrations/internal/utils/constants"
	"ecommerce-integrations/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type WebhookHandler struct {
	webhookUseCase usecase.WebhookUseCase
}

func NewWebhookHandler(db *driver.Database) WebhookHandler {
	tRepo := repository.NewWebhookRepository(db)
	pUseCase := usecase.NewWebhookUseCase(tRepo)
	return WebhookHandler{
		webhookUseCase: pUseCase,
	}
}
func (instance *WebhookHandler) Index(c *gin.Context) {
	app := app.Gin{C: c}

	number := c.DefaultQuery("n", "10")

	n, err := helpers.StringToInt(number)

	if err != nil {
		app.BadRequest(gin.H{
			"status_code": constants.INVALID_PARAMS,
		})
		return
	}

	webhooks, err := instance.webhookUseCase.GenerateWebhooks(n)
	if err != nil {
		app.Error(err.Error())
		return
	}

	app.Ok(gin.H{
		"webhooks": webhooks,
	})
}
