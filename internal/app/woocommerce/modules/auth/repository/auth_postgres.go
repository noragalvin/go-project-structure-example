package repository

import (
	"ecommerce-integrations/driver"
)

type authImpl struct {
	database *driver.Database
}

func NewAuthRepository(db *driver.Database) AuthRepository {
	return &authImpl{
		database: db,
	}
}

// func (instance *authImpl) BulkInsertAuths(someArr []interface{}) ([]models.Auth, error) {
// 	auths := []models.Auth{}
// 	mapstructure.Decode(someArr, &auths)
// 	// err := instance.database.Conn.Create(&auths).Error
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	return auths, nil
// }
