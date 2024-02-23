package main

import (
	"mytask-app/config"
	"mytask-app/delivery"
	"mytask-app/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDatabase()
	// this is a layer that process the database
	repo := repository.NewMyTaskRepository(db)

	r := gin.Default()
	// this is a layer that process the parameter of API
	delivery.NewMyTaskHandler(r, repo)

	r.Run()
}
