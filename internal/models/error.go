package models

type Error struct {
	Model
	Type     string      `json:"type" gorm:"type:varchar(100);"`
	Place    string      `json:"place" gorm:"type:text;"`
	Data     interface{} `json:"data" gorm:"type:text;"`
	FileName string      `json:"file_name"`
	Function string      `json:"function"`
	Line     int         `json:"line"`
}
