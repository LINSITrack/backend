package main

import (
	"log"
	"os"

	"github.com/LINSITrack/backend/src/db"
	"github.com/LINSITrack/backend/src/middleware"
	"github.com/LINSITrack/backend/src/models"
	"github.com/LINSITrack/backend/src/routes"
	"github.com/LINSITrack/backend/src/seed"
	"github.com/LINSITrack/backend/src/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func initAuth() {
	_ = godotenv.Load()

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("ERROR: SECRET_KEY is not set")
	}
	middleware.SetSecretKey(secretKey)
}

func main() {

	// Database connection
	db, err := db.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	// Carga de la secret key
	initAuth()

	// Setup de router
	router := gin.Default()

	// Configuración CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// Logs
	const Reset, Cyan = "\033[0m", "\033[36m"
	log.SetPrefix(Cyan)
	log.SetFlags(log.LstdFlags)
	log.Printf("-----------------------------------------------: %s\n", Reset)
	log.Printf("SERVER RUNNING ON: %s %s\n", "http://localhost:8080", Reset)
	log.Printf("-----------------------------------------------: %s\n", Reset)

	// Ruta unprotected
	router.GET("/unprotected", func(c *gin.Context) {
		c.String(200, "Hello from Gin! Server is up and running. (Unprotected Route)")
	})

	// Ruta protected
	router.GET("/protected", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.String(200, "Hello from Gin! Server is up and running. (Protected Route)")
	})

	// (¡Peligro: Borra la base de datos al descomentar! Excepto las instancias creadas con la seed al iniciar el servidor)
	db.Migrator().DropTable(
		&models.Profesor{}, 
		&models.Admin{}, 
		&models.Alumno{}, 
		&models.Materia{}, 
		&models.Comision{}, 
		&models.Cursada{}, 
		&models.Notificacion{}, 
		&models.ProfesorXComision{},
	)

	// Automigraciones
	if err := db.AutoMigrate(
		&models.Profesor{},
		&models.Admin{},
		&models.Alumno{},
		&models.Materia{},
		&models.Comision{},
		&models.Cursada{},
		&models.Notificacion{},
		&models.ProfesorXComision{},
	); err != nil {
		log.Fatalf("Error during auto migration: %v\n", err)
	}

	// Ejecutar seed inicial de la DB
	seed.AdminSeed(db)
	seed.ProfesorSeed(db)
	seed.AlumnoSeed(db)
	seed.MateriaSeed(db)
	seed.ComisionSeed(db)
	seed.CursadaSeed(db)
	seed.NotificacionSeed(db)
	seed.ProfesorXComisionSeed(db)

	// Setup de services
	authService := services.NewAuthService(db)
	profesorService := services.NewProfesorService(db)
	adminService := services.NewAdminService(db)
	alumnoService := services.NewAlumnoService(db)
	materiaService := services.NewMateriaService(db)
	comisionService := services.NewComisionService(db)
	cursadaService := services.NewCursadaService(db)
	notificacionService := services.NewNotificacionService(db)
	profesorXComisionService := services.NewProfesorXComisionService(db) // Nuevo service

	// Setup de rutas
	routes.SetupAuthRoutes(router, authService)
	routes.SetupProfesoresRoutes(router, profesorService)
	routes.SetupAdminsRoutes(router, adminService)
	routes.SetupAlumnosRoutes(router, alumnoService)
	routes.SetupMateriasRoutes(router, materiaService)
	routes.SetupComisionRoutes(router, comisionService)
	routes.SetupCursadasRoutes(router, cursadaService)
	routes.SetupNotificacionRoutes(router, notificacionService)
	routes.SetupProfesorXComisionRoutes(router, profesorXComisionService) // Nuevas rutas

	// Run
	router.Run()

}
