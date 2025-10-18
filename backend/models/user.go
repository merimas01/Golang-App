package models

//DB model
type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Age   int    `json:"age"`
}

//DTO
type UserInsert struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" gorm:"unique" binding:"required,email"`
	Age   int    `json:"age" binding:"required,gte=0"`
}

//DTO
type UserUpdate struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" gorm:"unique" binding:"required,email"`
	Age   int    `json:"age" binding:"required,gte=0"`
}
