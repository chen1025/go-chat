package main

import (
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:letter123@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(interface{}("failed to connect database"))
	}

	// Migrate the schema
	db.AutoMigrate(&models.UserBasic{})
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.Contact{})
	db.AutoMigrate(&models.GroupBasic{})

	/*// Create
	user := &models.UserBasic{}
	user.Name = "chen"
	user.Password = "chen"
	db.Create(user)

	// Read
	var u models.UserBasic
	db.First(&u, 1) // find product with integer primary key
	fmt.Println(u)
	// Update - update product's price to 200
	db.Model(&u).Update("Password", "123456")
	// Update - update multiple fields

	// Delete - delete product
	//db.Delete(&u, 1)*/
}
