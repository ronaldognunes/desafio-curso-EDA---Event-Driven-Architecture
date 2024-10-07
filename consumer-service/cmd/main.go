package main

import (
	"consumer-service/internal/database"
	"consumer-service/internal/service"
	"consumer-service/internal/usecase/find_by_id"
	"consumer-service/internal/web"
	"consumer-service/internal/web/webserver"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysqlapi", "3310", "serviceapi"))
	//db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3310", "serviceapi"))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("iniciando..6")

	accountDb := database.NewAccountDb(db)
	consumer := service.NewConsumeService(accountDb)
	go consumer.ExecuteConsumerKafka()
	findByidUseCase := find_by_id.NewFindByIdAccountUseCase(accountDb)
	accounthandler := web.NewWebAccountHandler(*findByidUseCase)
	webserver := webserver.NewWebServer(":3003")
	webserver.AddHandler("/balance/{account_id}", accounthandler.FindAccount)
	fmt.Println("start Api")
	webserver.Start()
}
