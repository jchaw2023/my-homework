package model

// Employee 员工信息表对应的 sqlx 模型
type Employee struct {
	ID         uint    `db:"id"`         // 员工唯一ID（主键自增）
	Name       string  `db:"name"`       // 员工姓名（支持中文，最长50字符）
	Department string  `db:"department"` // 所属部门（如：技术部、人事部）
	Salary     float64 `db:"salary"`     // 员工薪水（精确到分，默认0.00）
}


