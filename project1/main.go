package main

import (
	"fmt"

	"example.go/Config"
	"example.go/Models"
	"example.go/Routes"
	"github.com/gin-contrib/cors"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("status: ", err)
	}

	r := Routes.SetUpRouter()

	// Update CORS middleware configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"}, // Update with your frontend URL
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"}, // Include necessary headers
	}))

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Todo{})

	r.Run()
}
