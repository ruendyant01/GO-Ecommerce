package main

import (
	"customer_service/controller"
	"customer_service/repo/impl"
	router2 "customer_service/router"
	"customer_service/service"
	"net/http"

	"github.com/ruendyant01/GO-Ecommerce/go_config/config"
	"github.com/ruendyant01/GO-Ecommerce/go_config/database"
)

func main() {
	db, err := OpenDb()
	if err != nil {
		panic(err)
	}

	repo := impl.NewCustomerRepositoryImpl(db.OrmInstance)
	svc := service.NewCustomerServiceImpl(repo)
	customerContr := controller.NewCustomerController(svc)
	router := router2.CustomerRouter(customerContr)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func OpenDb() (db *database.OrmDb, err error) {
	host := config.Default().GetString("db.postgres.host")
	port := config.Default().GetInt("db.postgres.port")
	username := config.Default().GetString("db.postgres.username")
	password := config.Default().GetString("db.postgres.password")
	databaseName := config.Default().GetString("db.postgres.database")

	db, err = database.OpenOrm(host, port, username, password, databaseName)
	if err != nil {
		return nil, err
	}
	return db, nil
}
