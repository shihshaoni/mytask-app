package config

import (
	"fmt"
	"mytask-app/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDatabase() *gorm.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "mytask"
		password = "mytaskpassword"
		dbname   = "mytask_app"
	)

	dbUri := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic("failed to connect database")
	}

	// create a table named my_tasks
	db.AutoMigrate(&model.MyTask{})
	return db
}
