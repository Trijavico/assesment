package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"size:55;unique;not null"`
	Password []byte `gorm:"not null"`
}

type Note struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `gorm:"size:150;not null"`
	Content  string `gorm:"type:TEXT"`
	Archived bool   `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
	User     User   `gorm:"constraint:OnDelete:CASCADE"`
}
