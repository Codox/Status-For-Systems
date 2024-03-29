package models

import "time"

type IncidentStatus struct {
	ID        uint64    `gorm:"primaryKey; column:id" json:"-"`
	UUID      string    `gorm:"column:uuid" json:"uuid"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

type IncidentUpdate struct {
	ID               uint64         `gorm:"primaryKey; column:id" json:"-"`
	UUID             string         `gorm:"column:uuid" json:"uuid"`
	Text             string         `gorm:"column:text;type:text" json:"text"`
	IncidentID       uint64         `gorm:"column:incident_id" json:"-"`
	IncidentStatusID uint64         `gorm:"column:incident_status_id" json:"-"`
	IncidentStatus   IncidentStatus `gorm:"foreignKey:IncidentStatusID;references:ID" json:"status"`
	CreatedAt        time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt        time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt        time.Time      `gorm:"column:deleted_at" json:"deletedAt"`
}

type Incident struct {
	ID               uint64           `gorm:"primaryKey; column:id" json:"-"`
	UUID             string           `gorm:"column:uuid" json:"uuid"`
	Title            string           `gorm:"column:title" json:"title"`
	IncidentStatusID uint64           `gorm:"column:incident_status_id" json:"-"`
	Status           IncidentStatus   `gorm:"foreignKey:IncidentStatusID;references:ID" json:"status"`
	Updates          []IncidentUpdate `gorm:"foreignKey:IncidentID;references:ID" json:"updates,omitempty"`
	Description      string           `gorm:"column:description;type:text" json:"description"`
	CreatedAt        time.Time        `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt        time.Time        `gorm:"column:updated_at" json:"updatedAt"`
}

type IncidentSystem struct {
	ID         uint64    `gorm:"primaryKey; column:id" json:"-"`
	UUID       string    `gorm:"column:uuid" json:"uuid"`
	SystemID   uint64    `gorm:"column:system_id" json:"-"`
	IncidentID uint64    `gorm:"column:incident_id" json:"-"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updatedAt"`
}
