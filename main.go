package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/handlers"
	"kasir-api/models"
	"kasir-api/repositories"
	"kasir-api/services"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

// In-memory storage (sementara, nanti ganti database)
var products = []models.Product{
	{ID: 1, Name: "Indomie Godog", Price: 3500, Stock: 10},
	{ID: 2, Name: "Vit 1000ml", Price: 3000, Stock: 40},
	{ID: 3, Name: "kecap", Price: 12000, Stock: 20},
}

var categories = []models.Category{
	{ID: 1, Name: "Mie Instant", Description: "Mie"},
	{ID: 2, Name: "Minuman", Description: "Minum Botol"},
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port: viper.GetString("PORT"),
	}

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	productRepository := repositories.NewProductRepository(products)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	http.HandleFunc("GET /api/products", productHandler.GetAll)
	http.HandleFunc("POST /api/products", productHandler.Create)
	http.HandleFunc("GET /api/products/{id}", productHandler.GetByID)
	http.HandleFunc("PUT /api/products/{id}", productHandler.Update)
	http.HandleFunc("DELETE /api/products/{id}", productHandler.Delete)

	categoryRepository := repositories.NewCategoryRepository(categories)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	http.HandleFunc("GET /api/categories", categoryHandler.GetAll)
	http.HandleFunc("POST /api/categories", categoryHandler.Create)
	http.HandleFunc("GET /api/categories/{id}", categoryHandler.GetByID)
	http.HandleFunc("PUT /api/categories/{id}", categoryHandler.Update)
	http.HandleFunc("DELETE /api/categories/{id}", categoryHandler.Delete)

	fmt.Println("Server Running di Localhost:" + config.Port)

	err := http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		fmt.Println("Gagal Running Server")
	}
}
