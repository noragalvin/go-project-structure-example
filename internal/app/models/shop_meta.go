package models

type ShopMeta struct {
	Model

	ShopId uint   `json:"shop_id"`
	Key    string `json:"key" gorm:"size:50;not null"`
	Value  string `json:"value" `
}
