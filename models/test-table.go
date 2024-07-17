package models

type TestTable struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (TestTable) TableName() string {
	return "test_tables"
}
