package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/vikram305/ecom/cmd/api"
	"github.com/vikram305/ecom/config"
	"github.com/vikram305/ecom/db"
)



func main(){

	database, error := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})
	if error!=nil {
		log.Fatal(error)
	}
	
	db.InitStorage(database)

	defer database.Close()

	server := api.NewAPIServer(":8080", database)
	err := server.Run()
	if err!=nil{
		log.Fatal(err)
	}
}