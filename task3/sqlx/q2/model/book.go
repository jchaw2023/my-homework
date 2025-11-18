package model

// Book 图书信息表对应的 sqlx 模型
// 对应表：books
type Book struct {
	ID     uint    `db:"id"`     // 图书唯一ID（主键自增）
	Title  string  `db:"title"`  // 图书标题（支持多语言，最长100字符）
	Author string  `db:"author"` // 作者（单作者或多作者，逗号分隔）
	Price  float64 `db:"price"`  // 图书价格（精确到分，默认0.00）
}


