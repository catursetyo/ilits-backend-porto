package main

import (
	"log"
	"os"

	"ilits-porto-backend/config"
	"ilits-porto-backend/models"
	"ilits-porto-backend/routes"
)

func main() {
	config.ConnectDatabase()

	err := config.DB.AutoMigrate(&models.Portfolio{})
	if err != nil {
		log.Fatalf("Gagal melakukan migrasi database: %v", err)
	}
	log.Println("Migrasi database berhasil dijalankan!")

	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server berjalan di port :%s ...\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
