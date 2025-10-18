package seed

import (
	"log"

	"Golang-App/models" // adjust to your actual module path

	"gorm.io/gorm"
)

// SeedUsers inserts sample users into the database
func SeedUsers(db *gorm.DB) {
	users := []models.User{
		{Name: "Alice Smith", Email: "alice@example.com", Age: 25},
		{Name: "Bob Johnson", Email: "bob@example.com", Age: 30},
		{Name: "Charlie Brown", Email: "charlie@example.com", Age: 22},
		{Name: "Diana Prince", Email: "diana@example.com", Age: 28},
		{Name: "Ethan Hunt", Email: "ethan@example.com", Age: 35},
	}

	for _, u := range users {
		// FirstOrCreate checks if a user with the same email exists
		if err := db.FirstOrCreate(&u, models.User{Email: u.Email}).Error; err != nil {
			log.Printf("Failed to seed user %s: %v\n", u.Email, err)
		} else {
			log.Printf("Seeded user %s\n", u.Email)
		}
	}
}
