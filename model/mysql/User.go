package mysql

type User struct {
	ID int64 `gorm:"primarkKey"`
	Name string
	Age int
}
