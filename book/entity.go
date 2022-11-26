package book

import "time"

type Book struct {
	ID        int       `gorm:"type:int(10);primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(100)" json:"title"`
	Price     int       `gorm:"type:int(10)" json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updateAt"`
}
