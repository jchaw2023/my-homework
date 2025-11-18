package model

import (
	"time"

	"gorm.io/gorm"
)

// Post 文章
type Post struct {
	ID         uint64     `gorm:"primaryKey;autoIncrement"`
	Title      string     `gorm:"type:varchar(255);not null;comment:标题"`
	Content    string     `gorm:"type:longtext;not null;comment:内容"`
	UserID     uint64     `gorm:"not null;index;comment:作者ID"`
	CommentNum uint64     `gorm:"type:bigint unsigned;not null;default:0;comment:评论数"`
	CreatedAt  *time.Time `gorm:"comment:创建时间"`
	UpdatedAt  *time.Time `gorm:"comment:更新时间"`

	// 关联
	User     User      `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	if p.UserID != 0 {
		tx.Model(&User{}).
			Where("id = ? ", p.UserID).
			UpdateColumn("post_num", gorm.Expr("post_num + 1"))
		if tx.Error != nil {
			return tx.Error
		}
		return nil
	}
	return nil
}
func (p *Post) AfterDelete(tx *gorm.DB) (err error) {
	if p.UserID != 0 {
		tx.Model(&User{}).Where("id = ?", p.UserID).
			UpdateColumn("post_num", gorm.Expr("post_num - 1"))
		if tx.Error != nil {
			return tx.Error
		}
		return nil
	}
	return nil
}
