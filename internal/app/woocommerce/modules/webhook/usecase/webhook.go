package usecase

import (
	"ecommerce-integrations/internal/app/models"
	"ecommerce-integrations/internal/app/woocommerce/modules/webhook/repository"

	"github.com/bxcodec/faker/v3"
)

type WebhookUseCase struct {
	webhookRepo repository.WebhookRepository
}

func NewWebhookUseCase(pr repository.WebhookRepository) WebhookUseCase {
	return WebhookUseCase{
		webhookRepo: pr,
	}
}

func (instance *WebhookUseCase) GenerateWebhooks(n int) ([]models.Webhook, error) {
	webhookFakeData := make([]interface{}, 0)
	for i := 0; i < n; i++ {
		p := struct {
			Title string `json:"title" faker:"name"`
		}{}
		_ = faker.FakeData(&p)

		webhookFakeData = append(webhookFakeData, p)
	}

	webhooks, err := instance.webhookRepo.BulkInsertWebhooks(webhookFakeData)

	if err != nil {
		return webhooks, err
	}
	return webhooks, nil
}
