package models

type Shop struct {
	Model

	DefaultDomain string `json:"default_domain" gorm:"size:50;not null"`
	ShopifyDomain string `json:"shopify_domain" gorm:"size:50;not null"`
}
