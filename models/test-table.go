package models

type TestTable struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (TestTable) TableName() string {
	return "test_tables"
}
