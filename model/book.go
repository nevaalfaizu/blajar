package model

type Book struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	Author     string   `json:"author"`
	Year       int      `json:"year"`
	CategoryID int      `json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"category"`
}
