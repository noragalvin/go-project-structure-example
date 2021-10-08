package models

import "time"

type Product struct {
	Model

	// ShopId uint `json:"shop_id"`
	// MediaId uint `json:"media_id"`
	// BaseVariantId uint `json:"base_variant_id"`
	// BrandId uint `json:"brand_id"`

	Title           string     `json:"title" gorm:"size:256;not null"`
	Handle          string     `json:"handle" gorm:"size:256;not null"`
	Description     string     `json:"description" gorm:"type:text"`
	Sku             string     `json:"sku" gorm:"size:256"`
	TemplateSuffix  string     `json:"template_suffix" gorm:"size:256"`
	Status          string     `json:"status" gorm:"size:50"`
	TitleMeta       string     `json:"title_meta" gorm:"size:256"`
	DescriptionMeta string     `json:"description_meta" gorm:"size:500"`
	CurrentVersion  string     `json:"current_version" gorm:"size:15"`
	AverageRating   float64    `json:"average_rating" gorm:"default:0"`
	ReviewCount     int        `json:"review_count" gorm:"default:0"`
	PublishedAt     *time.Time `json:"published_at"`
}
