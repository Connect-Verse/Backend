package utils

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)



func DatabaseConnection() *gorm.DB {

	err:= godotenv.Load()

	if err!=nil {
		fmt.Println("error occured while loading the env")
	}

	port , err := strconv.Atoi(os.Getenv("DATABASE_PORT"))

	var (
		host= os.Getenv("DATABASE_HOST")
		user= os.Getenv("DATABASE_USER")
		password= os.Getenv("DATABASE_PASSWORD")
		dbName= os.Getenv("DATABASE_NAME")
	)

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db,err:= gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
    
	if err!=nil{
		panic(err)
	}

	return db
}