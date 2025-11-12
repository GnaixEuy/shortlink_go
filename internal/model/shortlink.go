package model

import "time"

type ShortLink struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Code      string `gorm:"uniqueIndex;size:16;not null"`
	URL       string `gorm:"type:text;not null"`
	Clicks    int64  `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpireAt  *time.Time `gorm:"index;null"` // 可选：到期时间
}



