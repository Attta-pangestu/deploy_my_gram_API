package models

import (
	"time"
)

type GormModel struct {
    ID        uint           `gorm:"primary_key;autoIncrement" json:"id"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
