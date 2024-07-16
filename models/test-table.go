package models

type TestTable struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
}

func (TestTable) TableName() string {
	return "test_tables"
}
