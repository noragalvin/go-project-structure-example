package models

type ProductVariant struct {
	Model

	ProductId uint `json:"product_id"`
	// Media uint `json:"media_id"`

	Title       string  `json:"title" gorm:"size:256;not null"`
	Description string  `json:"description" gorm:"type:text"`
	CostPrice   float64 `json:"cost_price"`
	SalePrice   float64 `json:"sale_price"`
	Price       float64 `json:"price"`

	// These attribute should be in product_variant_metas
	// SalePriceFrom *time.Time `json:"sale_price_from"`
	// SalePriceTo *time.Time `json:"sale_price_to"`

	Width              float64 `json:"width"`
	Length             float64 `json:"length"`
	Height             float64 `json:"height"`
	Weight             float64 `json:"weight"`
	Sku                string  `json:"sku" gorm:"size:256"`
	Barcode            string  `json:"barcode" gorm:"size:256"`
	InventoryQuantity  int     `json:"inventory_quantity"`
	InventoryStatus    string  `json:"inventory_status" gorm:"size:10"` // instock, outofstock
	ManageInventory    bool    `json:"manage_inventory" gorm:"default:false"`
	LowInventoryAmount int     `json:"low_inventory_amount"`
	SoldIndividually   bool    `json:"sold_individually" gorm:"default:false"`
	InventoryPolicy    string  `json:"inventory_policy" gorm:"size:10"`
	IsDigital          bool    `json:"is_digital" gorm:"default:false"`
}
