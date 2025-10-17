package models

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Age   int    `json:"age"`
}

type UserInsert struct {
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Age   int    `json:"age"`
}

type UserUpdate struct {
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Age   int    `json:"age"`
}
