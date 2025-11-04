package seed

import (
	"log"

	"github.com/LINSITrack/backend/src/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AlumnoSeed(db *gorm.DB) {
	
    // Seed de Alumno
	var existingAlumno models.Alumno
	result := db.Where("email = ?", "joaquin@linsi.com").First(&existingAlumno)
	if result.Error == nil {
		log.Println("Alumno 'joaquin@linsi.com' already exists")
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("joaquin"), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Failed to hash password for alumno: %v\n", err)
			return
		}

		newAlumno := models.Alumno{
			BaseUser: models.BaseUser{
				Nombre:   "Joaquin",
				Apellido: "Botteri",
				Email:    "joaquin@linsi.com",
				Password: string(hashedPassword),
			},
			Legajo: "54321",
		}

		if err := db.Create(&newAlumno).Error; err != nil {
			log.Printf("Failed to create alumno: %v\n", err)
		} else {
			log.Println("Alumno 'joaquin@linsi.com' created successfully")
		}
	}

    // Log
	log.Println("Alumno seed completed successfully")
}
