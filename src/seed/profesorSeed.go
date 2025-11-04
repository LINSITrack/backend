package seed

import (
	"log"

	"github.com/LINSITrack/backend/src/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ProfesorSeed(db *gorm.DB) {

	// Seed de Profesor
	var existingProfesor models.Profesor
	result := db.Where("email = ?", "martin@linsi.com").First(&existingProfesor)
	if result.Error == nil {
		log.Println("Profesor 'martin@linsi.com' already exists")
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("martin"), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Failed to hash password for profesor: %v\n", err)
			return
		}

		newProfesor := models.Profesor{
			BaseUser: models.BaseUser{
				Nombre:   "Martin",
				Apellido: "Jorge",
				Email:    "martin@linsi.com",
				Password: string(hashedPassword),
			},
			Legajo: "12345",
		}

		if err := db.Create(&newProfesor).Error; err != nil {
			log.Printf("Failed to create profesor: %v\n", err)
		} else {
			log.Println("Profesor 'martin@linsi.com' created successfully")
		}
	}
	
    // Log
	log.Println("Profesor seed completed successfully")
}
