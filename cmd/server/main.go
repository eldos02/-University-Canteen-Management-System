package main

import (
	"UniversityCanteenManagementSystem/db"
)

func main() {
	//Подключаемся к postgreSQL и иницализируем модели
	db.ConnectDatabase()

}
