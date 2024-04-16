package main

import (
	"net/http"

	"github.com/bruno-holanda15/go_expert_fc/APIs_project/configs"
	"github.com/bruno-holanda15/go_expert_fc/APIs_project/internal/entity"
	"github.com/bruno-holanda15/go_expert_fc/APIs_project/internal/infra/database"
	"github.com/bruno-holanda15/go_expert_fc/APIs_project/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	userDB := database.NewUserDB(db)
	productHandler := handlers.NewProductHandler(productDB)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/product", productHandler.CreateProduct)
	r.Get("/product/{id}", productHandler.GetProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Put("/product/{id}", productHandler.UpdateProduct)
	r.Delete("/product/{id}", productHandler.DeleteProduct)

	r.Post("/user", userHandler.CreateUser)

	http.ListenAndServe(":8001", r)
}
