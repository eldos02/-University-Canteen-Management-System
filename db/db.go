package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// делаем базу данных глобальной переменной чтобы все функций имели доступ
var database *gorm.DB
var err error

func ConnectDatabase() {
	//Данные для входа в базу данных
	dbs := "host=db user=sdudent dbname=sdu password=kek port=5556"
	database, err = gorm.Open(postgres.Open(dbs), &gorm.Config{})

	//Проверяем валидность данных для подключение DataBase
	if err != nil {
		log.Fatal("Не получилось подключиться, Данные фигня, исправь", err)
	} else {
		log.Println("Красавчик, база данных подключена!")
	}
}
