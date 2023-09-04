package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	controllError(err)

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create a new product
	db.Create(&Product{Code: "123", Price: 100})

	// Read the product
	var product Product
	db.First(&product, "code =?", "123")
	fmt.Println(product)

	// Read all products
	products := []Product{}
	db.Find(&products)

	for _, product := range products {
		println(product.Code)
	}

	// Update the product
	db.Model(&product).Update("price", 200)

	// Delete the product
	db.Delete(&product)
}

func controllError(err error) {
	if err != nil {
		panic(err)
	}
}
