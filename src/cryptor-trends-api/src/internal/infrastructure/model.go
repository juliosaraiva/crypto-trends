package infrastructure

import "github.com/google/uuid"

type Cryptocurrency struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name        string    `gorm:"type:varchar(50)"`
	Symbol      string    `gorm:"type:varchar(20)"`
	Rank        int       `gorm:"column:rank"`
	MaxSupply   int       `gorm:"column:max_supply"`
	Ciruclating int       `gorm:"column:ciruclating"`
	TotalSupply int       `gorm:"column:total_supply"`
	Price       float64   `gorm:"type:float"`
	TimeStamp   int       `gorm:"column:timestamp"`
	Trend       string    `gorm:"type:varchar(8)"`
}
