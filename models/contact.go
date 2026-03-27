package models

import "time"

type Contact struct {
	ID        uint      `gorm:"primaryKey;column:id;type:BIGINT UNSIGNED AUTO_INCREMENT"`
	FullName  string    `gorm:"column:full_name;type:VARCHAR(100);not null"`
	Email     string    `gorm:"column:email_address;type:VARCHAR(100);not null"`
	Phone     string    `gorm:"column:phone_number;type:VARCHAR(20);not null"`
	Message   string    `gorm:"column:message_text;type:TEXT;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:DATETIME;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:DATETIME;autoCreateTime"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:DATETIME;index"`
}

// untuk gorm. gorm akan cari tablename
func (Contact) TableName() string {
	return "contact_messages"
}
