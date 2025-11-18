package model

import (
	"time"

	"gorm.io/gorm"
)

// Comment 评论
type Comment struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement"`
	Content   string     `gorm:"type:varchar(512);not null;comment:评论内容"`
	UserID    uint64     `gorm:"not null;index;comment:用户ID"`
	PostID    uint64     `gorm:"not null;index;comment:文章ID"`
	CreatedAt *time.Time `gorm:"comment:创建时间"`
	UpdatedAt *time.Time `gorm:"comment:更新时间"`

	// 关联
	User User `gorm:"foreignKey:UserID"`
	Post Post `gorm:"foreignKey:PostID"`
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	if c.PostID != 0 { //每删除一条评论，文章的评论数量减1
		tx.Model(&Post{}).Where("id = ? AND user_id = ? AND comment_num > 0", c.PostID, c.UserID).
			UpdateColumn("comment_num", gorm.Expr("comment_num - 1"))
		if tx.Error != nil {
			return tx.Error
		}
		return nil
	}
	return nil
}

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	if c.PostID != 0 {
		tx.Model(&Post{}).Where("id = ? AND user_id = ?", c.PostID, c.UserID).
			UpdateColumn("comment_num", gorm.Expr("comment_num + 1"))
		if tx.Error != nil {
			return tx.Error
		}
		return nil
	}
	return nil
}
