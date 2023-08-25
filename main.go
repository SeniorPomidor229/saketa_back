package main

import (
	"log"
	"saketa/configs"
	"saketa/controllers"
	"saketa/services"
)



func main() {
	db := configs.NewDB()
	defer db.Close()

	employeeService := services.NewEmployeeService(db)
	app := controllers.NewApp(employeeService)

	router := app.SetupRoutes()

	log.Fatal(router.Listen(":3000"))
}