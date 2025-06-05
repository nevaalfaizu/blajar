package model

type Inventory struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Quantity   int      `json:"quantity"`
	Price      float64  `json:"price"`
	Year       int      `json:"year"`
	CategoryID int      `json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"category"`
}
