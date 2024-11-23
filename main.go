package main

import (
	"fmt"
	"go-ducpa/Config"
	"go-ducpa/Models/Schema"
	"go-ducpa/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(
		&Schema.OAuthAccessToken{},
		&Schema.OAuthClient{},
		&Schema.OAuthRefreshToken{},
		&Schema.User{},
	)

	r := Routes.SetupRouter()
	_ = r.Run()
}
