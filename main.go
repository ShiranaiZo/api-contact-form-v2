package main

import (
	"log"

	"github.com/ShiranaiZo/api-contact-form-v2/config"
	"github.com/ShiranaiZo/api-contact-form-v2/handlers"
	"github.com/ShiranaiZo/api-contact-form-v2/helpers"
	"github.com/ShiranaiZo/api-contact-form-v2/repositories"
	"github.com/ShiranaiZo/api-contact-form-v2/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config.InitDB()

	mainHandler := handlers.NewMainHandler()
	healthHandler := handlers.NewHealthHandler()
	contactRepository := repositories.NewContactRepository(config.DB)
	contactService := services.NewContactService(contactRepository)
	contactHandler := handlers.NewContactHandler(contactService)

	router := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     helpers.ParseEnvList("CORS_ALLOWED_ORIGINS"),
		AllowMethods:     helpers.ParseEnvList("CORS_ALLOWED_METHODS"),
		AllowHeaders:     helpers.ParseEnvList("CORS_ALLOWED_HEADERS"),
		AllowCredentials: helpers.GetEnvBool("CORS_ALLOW_CREDENTIALS", false),
		ExposeHeaders:    helpers.ParseEnvList("CORS_EXPOSE_HEADERS"),
		MaxAge:           12 * 60 * 60,
	}

	router.Use(cors.New(corsConfig))

	router.GET("/", mainHandler.MainHandler)
	router.GET("/health", healthHandler.HealthCheck)
	router.GET("/contacts", contactHandler.GetContacts)
	router.GET("/contacts/:id", contactHandler.GetContact)
	router.POST("/contacts", contactHandler.CreateContact)
	router.PUT("/contacts/:id", contactHandler.UpdateContact)
	router.DELETE("/contacts/:id", contactHandler.DeleteContact)

	appPort := config.GetEnv("APP_PORT", "8080")
	err = router.Run(":" + appPort)
	if err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
