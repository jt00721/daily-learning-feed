package domain

import "time"

type Resource struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	URL       string `gorm:"not null;unique"`
	Category  string
	Source    string
	DateAdded time.Time `gorm:"autoCreateTime"`
}
