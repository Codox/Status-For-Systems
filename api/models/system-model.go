package models

import (
	"time"
)

type System struct {
	ID             uint64       `gorm:"primaryKey; column:id" json:"-"`
	UUID           string       `gorm:"column:uuid" json:"uuid"`
	Name           string       `gorm:"column:name" json:"name"`
	Status         SystemStatus `gorm:"foreignKey:SystemStatusID;references:ID" json:"status"`
	GroupID        uint64       `gorm:"column:group_id" json:"-"`
	SystemStatusID uint64       `gorm:"column:system_status_id" json:"-"`
	CreatedAt      time.Time    `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt      time.Time    `gorm:"column:updated_at" json:"updatedAt"`
}

type SystemGroup struct {
	ID        uint64    `gorm:"primaryKey; column:id" json:"-"`
	UUID      string    `gorm:"column:uuid" json:"uuid"`
	Name      string    `gorm:"column:name" json:"name"`
	Systems   []System  `gorm:"foreignKey:GroupID;references:ID;preload:false" json:"systems"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

type SystemStatus struct {
	ID        uint64    `gorm:"primaryKey; column:id" json:"-"`
	UUID      string    `gorm:"column:uuid" json:"uuid"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}
