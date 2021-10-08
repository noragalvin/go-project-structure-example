package usecase

import (
	"ecommerce-integrations/internal/app/woocommerce/modules/auth/repository"
)

type AuthUseCase struct {
	authRepo repository.AuthRepository
}

func NewAuthUseCase(pr repository.AuthRepository) AuthUseCase {
	return AuthUseCase{
		authRepo: pr,
	}
}

// func (instance *AuthUseCase) GenerateAuths(n int) ([]models.Auth, error) {
// 	authFakeData := make([]interface{}, 0)
// 	for i := 0; i < n; i++ {
// 		p := struct {
// 			Title string `json:"title" faker:"name"`
// 		}{}
// 		_ = faker.FakeData(&p)

// 		authFakeData = append(authFakeData, p)
// 	}

// 	auths, err := instance.authRepo.BulkInsertAuths(authFakeData)

// 	if err != nil {
// 		return auths, err
// 	}
// 	return auths, nil
// }
