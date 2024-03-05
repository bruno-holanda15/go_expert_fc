package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&Product{})

	// //create in batch
	// productsCreate := []Product{
	// 	{Name: "Nintendo", Price: 2000.00},
	// 	{Name: "Keycron", Price: 1000.00},
	// 	{Name: "Cama", Price: 3000.00},
	// }

	// db.Create(&productsCreate)

	// select one
	var productOne Product
	db.First(&productOne)
	fmt.Println(productOne)

	// db.First(&productOne, "name= ?", "Teclado")
	// fmt.Println(productOne)

	// select all
	// var products []Product
	// db.Find(&products)
	
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// db.Limit(2).Offset(2).Find(&products)
	
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// where
	// db.Where("price < ?", 1100).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// db.Where("name LIKE ?", "%PS%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// update and delete
	var p Product
	db.First(&p)
	p.Name = "Xbox One"
	db.Save(&p)

	// var p2 Product
	// db.First(&p2)
	// fmt.Println(p2)
	// db.Delete(&p2)


}
