package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/vikram305/ecom/cmd/api"
	"github.com/vikram305/ecom/db"
)



func main(){

	server := api.NewAPIServer(":8080", nil)
	err := server.Run()
	if err!=nil{
		log.Fatal(err)
	}
}