package main

import (
	"fmt"
	"namazu/Config"
	"namazu/Routes"

	// "namazu/Models"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	// Config.DB.AutoMigrate(&Models.User{})

	r := Routes.SetupRouter()
	//running
	r.Run(":4000")
}
