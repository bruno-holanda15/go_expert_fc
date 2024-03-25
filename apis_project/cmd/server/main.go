package main

import (
	"net/http"

	"github.com/bruno-holanda15/go_expert_fc/APIs_project/configs"
	"github.com/bruno-holanda15/go_expert_fc/APIs_project/internal/entity"
	"github.com/bruno-holanda15/go_expert_fc/APIs_project/internal/infra/database"
	"github.com/bruno-holanda15/go_expert_fc/APIs_project/internal/infra/webserver"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProductDB(db)
	productHandler := webserver.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8001", nil)
}
