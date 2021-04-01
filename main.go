package main

import (
	"fmt"

	"github.com/michaelpege/arc/Config"
	"github.com/michaelpege/arc/Models"
	"github.com/michaelpege/arc/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main(){
	//connect to SQL
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	//check for an error
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	//migrate models
	Config.DB.AutoMigrate(&Models.Post{})
	Config.DB.AutoMigrate(&Models.Comment{})

	//create router
	r := Routes.SetupRouter()
	//running
	r.Run()
}