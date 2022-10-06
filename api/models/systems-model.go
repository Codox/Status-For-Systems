package models

import (
  "time"
)

type System struct {
	ID          uint64      `gorm:"primaryKey; column:id" json:"-"`
	Name        string      `gorm:"column:name" json:"name"`
	GroupID     uint64      `gorm:"column:group_id" json:"-"`
	SystemGroup SystemGroup `gorm:"foreignKey:GroupID" json:"systemGroup"`
	CreatedAt   time.Time   `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time   `gorm:"column:updated_at" json:"updatedAt"`
}

type SystemGroup struct {
	ID        uint64    `gorm:"primaryKey; column:id" json:"-"`
	Name      string    `gorm:"column:name" json:"name"`
	Systems   []System  `gorm:"foreignKey:GroupID" json:"systems"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}