package model

import "time"

// User 用户
type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"type:varchar(64);uniqueIndex;not null;comment:用户名"`
	Email     string    `gorm:"type:varchar(128);uniqueIndex;not null;comment:邮箱"`
	Password  string    `gorm:"type:varchar(255);not null;comment:密码哈希"`
	PostNum   int64     `gorm:"type:bigint;not null;default:0;comment:发帖数量"`
	CreatedAt time.Time `gorm:"comment:创建时间"`
	UpdatedAt time.Time `gorm:"comment:更新时间"`

	// 关联
	Posts    []Post    `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:UserID"`
}
