package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/luizhen1/go-clean-gin-crud/internal/database"
	"github.com/luizhen1/go-clean-gin-crud/internal/repositories"
	"github.com/luizhen1/go-clean-gin-crud/internal/routes"
	"github.com/luizhen1/go-clean-gin-crud/internal/services"
)

// @title           Go Clean Arch CRUD API
// @version         1.0
// @description     API de exemplo de CRUD utilizando Clean Architecture e Gin.
// @host            localhost:8080
// @BasePath        /

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}
	defer db.Close()

	repo := repositories.NewMySQLProductRepository(db)

	service := services.NewProductService(repo)

	r := gin.Default()

	routes.InitProductRoutes(r, service)

	log.Println("Servidor rodando na porta 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}
