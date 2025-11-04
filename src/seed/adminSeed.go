package seed

import (
	"log"

	"github.com/LINSITrack/backend/src/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AdminSeed(db *gorm.DB) {

	// Seed de Admin
	var existingAdmin models.Admin
	result := db.Where("email = ?", "mateo@linsi.com").First(&existingAdmin)
	if result.Error == nil {
		log.Println("Admin 'mateo@linsi.com' already exists")
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("mateo"), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Failed to hash password for admin: %v\n", err)
			return
		}

		newAdmin := models.Admin{
			BaseUser: models.BaseUser{
				Nombre:   "Mateo",
				Apellido: "Polci",
				Email:    "mateo@linsi.com",
				Password: string(hashedPassword),
			},
		}

		if err := db.Create(&newAdmin).Error; err != nil {
			log.Printf("Failed to create admin: %v\n", err)
		} else {
			log.Println("Admin 'admin@linsi.com' created successfully")
		}
	}

    // Log
	log.Println("Admin seed completed successfully")
}
